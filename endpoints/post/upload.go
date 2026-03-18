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
	// TODO: Initiation of url for sending photo
	api_key := os.Getenv("TG_API")
	chat_id := os.Getenv("TG_CHATID")

	// TODO: To fetch file from uploads
	file, err := ctx.FormFile("media")
	if err != nil {
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	open, err := file.Open()
	if err != nil {
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	defer open.Close()

	// TODO: Create a copy for mimetype

	copy_, err := file.Open()
	if err != nil {
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	defer copy_.Close()

	buf := make([]byte, 512)
	_, err = copy_.Read(buf)
	if err != nil {
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	// TODO: Initialized file_url
	file_url := fmt.Sprintf("https://api.telegram.org/bot%s/sendPhoto", api_key)
	mimeType := http.DetectContentType(buf)

	// TODO: Creating buffers for Telegram Upload
	var buff bytes.Buffer
	filetype := "photo"
	writer := multipart.NewWriter(&buff)
	_ = writer.WriteField("chat_id", chat_id)

	// TODO: To identify what file type it is for different upload
	if strings.HasPrefix(mimeType, "image") {
		file_url = fmt.Sprintf("https://api.telegram.org/bot%s/sendPhoto", api_key)
		filetype = "photo"
	} else if strings.HasPrefix(mimeType, "video") {
		file_url = fmt.Sprintf("https://api.telegram.org/bot%s/sendVideo", api_key)
		filetype = "video"
	} else {
		file_url = fmt.Sprintf("https://api.telegram.org/bot%s/sendDocument", api_key)
		filetype = "document"
	}

	part, err := writer.CreateFormFile(filetype, file.Filename)
	if err != nil {
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
	}

	// TODO: Copying files
	_, err = io.Copy(part, open)
	if err != nil {
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
	}

	writer.Close()

	// TODO: To Fetch data thru API
	req, err := http.NewRequest("POST", file_url, &buff)
	if err != nil {
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.ContentLength = int64(buff.Len())

	client := &http.Client{}
	resp, err := client.Do(req)
	// TODO: Return error
	if err != nil {
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	defer resp.Body.Close()

	respBody, _ := io.ReadAll(resp.Body)

	ctx.JSON(200, gin.H{
		"from": json.RawMessage(respBody),
	})
}
