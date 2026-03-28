package post

import (
	"bytes"
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

var Upload = utils.Route{
	Path:       "/upload",
	Method:     utils.METHOD_POST,
	Permission: utils.PERMISSION_ADMIN,
	Handler:    uploadFile,
}

func uploadFile(ctx *gin.Context) {
	apiKey := os.Getenv("TG_API")
	chatID := os.Getenv("TG_CHATID")

	// Fetch file
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

	// Read first 512 bytes for MIME detection
	buf := make([]byte, 512)
	n, err := f.Read(buf)
	if err != nil && err != io.EOF {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	mimeType := http.DetectContentType(buf[:n])

	// Prepare multipart buffer
	var buff bytes.Buffer
	writer := multipart.NewWriter(&buff)
	_ = writer.WriteField("chat_id", chatID)

	fileType := "document"
	fileURL := fmt.Sprintf("https://api.telegram.org/bot%s/sendDocument", apiKey)
	if strings.HasPrefix(mimeType, "image") {
		fileType = "photo"
		fileURL = fmt.Sprintf("https://api.telegram.org/bot%s/sendPhoto", apiKey)
	} else if strings.HasPrefix(mimeType, "video") {
		fileType = "document"
		fileURL = fmt.Sprintf("https://api.telegram.org/bot%s/sendDocument", apiKey)
	}

	part, err := writer.CreateFormFile(fileType, file.Filename)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	// Reset file pointer to start for copying
	if _, err := f.Seek(0, 0); err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	if _, err := io.Copy(part, f); err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	writer.Close()

	// Send to Telegram
	req, err := http.NewRequest("POST", fileURL, &buff)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	defer resp.Body.Close()

	respBody, _ := io.ReadAll(resp.Body)
	ctx.JSON(200, gin.H{
		"from": json.RawMessage(respBody),
	})
}
