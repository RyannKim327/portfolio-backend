package post

import (
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"strings"

	"portfolio-backend/utils"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var _ = godotenv.Load()

const (
	tgUploadEndpoint = "https://api.telegram.org/bot%s/%s"
	// Cloud Bot API uploads max out at 50 MB and bot downloads at 20 MB.
	// Chunks stay under 20 MB so the same bot can also fetch them back.
	chunkSize = 19 * 1024 * 1024
	maxUpload = 1 * 1024 * 1024 * 1024
)

var Upload = utils.Route{
	Path:       "/upload",
	Method:     utils.METHOD_POST,
	Permission: utils.PERMISSION_ADMIN,
	Handler:    uploadFile,
}

func uploadFile(ctx *gin.Context) {
	apiKey := os.Getenv("TG_API")
	chatID := os.Getenv("TG_CHATID")

	ctx.Request.Body = http.MaxBytesReader(ctx.Writer, ctx.Request.Body, maxUpload)

	file, err := ctx.FormFile("file")
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	f, err := file.Open()
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	defer f.Close()

	mimeType, err := detectMimeType(f)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	if _, err := f.Seek(0, io.SeekStart); err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	if file.Size <= chunkSize {
		endpoint, field := "sendDocument", "document"
		if strings.HasPrefix(mimeType, "image/") {
			endpoint, field = "sendPhoto", "photo"
		}

		resp, err := sendToTelegram(apiKey, chatID, endpoint, field, file.Filename, f)
		if err != nil {
			ctx.JSON(502, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(200, gin.H{
			"chunks": 1,
			"parts":  []json.RawMessage{resp},
		})
		return
	}

	total := (file.Size + chunkSize - 1) / chunkSize
	parts := make([]json.RawMessage, 0, total)

	for i := int64(0); i < total; i++ {
		offset := i * chunkSize
		length := min(int64(chunkSize), file.Size-offset)

		if _, err := f.Seek(offset, io.SeekStart); err != nil {
			ctx.JSON(500, gin.H{"error": err.Error()})
			return
		}

		name := fmt.Sprintf("%s.part%03dof%03d", file.Filename, i+1, total)
		resp, err := sendToTelegram(apiKey, chatID, "sendDocument", "document", name, io.LimitReader(f, length))
		if err != nil {
			ctx.JSON(502, gin.H{
				"error":  err.Error(),
				"chunk":  i + 1,
				"chunks": total,
				"parts":  parts,
			})
			return
		}

		parts = append(parts, resp)
	}

	ctx.JSON(200, gin.H{
		"chunks": total,
		"parts":  parts,
	})
}

func detectMimeType(f io.ReadSeeker) (string, error) {
	buf := make([]byte, 512)
	n, err := f.Read(buf)
	if err != nil && err != io.EOF {
		return "", err
	}
	return http.DetectContentType(buf[:n]), nil
}

func sendToTelegram(apiKey, chatID, endpoint, field, filename string, r io.Reader) (json.RawMessage, error) {
	pr, pw := io.Pipe()
	writer := multipart.NewWriter(pw)

	go func() {
		var werr error
		defer func() {
			pw.CloseWithError(werr)
		}()

		if werr = writer.WriteField("chat_id", chatID); werr != nil {
			return
		}

		part, err := writer.CreateFormFile(field, filename)
		if err != nil {
			werr = err
			return
		}

		if _, werr = io.Copy(part, r); werr != nil {
			return
		}

		werr = writer.Close()
	}()

	url := fmt.Sprintf(tgUploadEndpoint, apiKey, endpoint)
	req, err := http.NewRequest(http.MethodPost, url, pr)
	if err != nil {
		pr.CloseWithError(err)
		return nil, err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		pr.CloseWithError(err)
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode >= 400 {
		return json.RawMessage(body), fmt.Errorf("telegram responded with %s", resp.Status)
	}

	return json.RawMessage(body), nil
}
