package post

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"

	"portfolio-backend/utils"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var _ = godotenv.Load()

var UploadImage = utils.Route{
	Path:       "/upload-image",
	Method:     utils.METHOD_POST,
	Permission: utils.PERMISSION_ADMIN,
	Handler:    uploadImage,
}

func uploadImage(ctx *gin.Context) {
	// TODO: Initiation of url for sending photo
	api_key := os.Getenv("TG_API")
	chat_id := os.Getenv("TG_CHATID")
	file_url := fmt.Sprintf("https://api.telegram.org/bot%s/sendPhoto", api_key)

	// TODO: To fetch file from uploads
	file, err := ctx.FormFile("image")
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

	var buff bytes.Buffer
	writer := multipart.NewWriter(&buff)
	_ = writer.WriteField("chat_id", chat_id)

	part, err := writer.CreateFormFile("photo", file.Filename)
	if err != nil {
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
	}

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
