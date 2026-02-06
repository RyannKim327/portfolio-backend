package get

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	utils "portfolio-backend/utils"

	"github.com/gin-gonic/gin"
)

/*
 * TODO: This is just a template for the other endpoint
 */

type result struct {
	Ok     bool `json:"ok"`
	Result struct {
		FilePath string `json:"file_path"`
	} `json:"result"`
}

var Image = utils.Route{
	Path:    "images",
	Method:  utils.METHOD_GET,
	Handler: fetchImage,
}

func fileDownload(file string) string {
	// TODO: Setup for credentials and URL
	api_key := os.Getenv("TG_API")
	file_url := fmt.Sprintf("https://api.telegram.org/file/bot%s/%s", api_key, file)

	return file_url
}

func fetchImage(ctx *gin.Context) {
	// TODO: Setup for credentials and URL
	api_key := os.Getenv("TG_API")
	file_id := ctx.Query("file")
	file_url := fmt.Sprintf("https://api.telegram.org/bot%s/getFile?file_id=%s", api_key, file_id)

	req, err := http.NewRequest("GET", file_url, nil)
	if err != nil {
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
	}

	// TODO: Request Executor
	client := &http.Client{}
	resp, err := client.Do(req)
	// TODO: To prevent errors
	if err != nil {
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	// TODO: To send the response data
	var data result
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		ctx.JSON(500, gin.H{
			"err": err.Error(),
		})
		return
	}

	image := fileDownload(data.Result.FilePath)

	fileResp, err := http.Get(image)
	if err != nil {
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	defer fileResp.Body.Close()
	fileBytes, err := io.ReadAll(fileResp.Body)
	if err != nil {
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
		return

	}
	ctx.Data(200, "image/jpeg", fileBytes)

	ctx.JSON(200, gin.H{
		"response": "",
	})
}
