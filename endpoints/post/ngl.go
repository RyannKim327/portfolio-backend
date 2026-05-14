package post

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"portfolio-backend/utils"

	"github.com/gin-gonic/gin"
)

var NGL = utils.Route{
	Path:       "/ngl",
	Method:     utils.METHOD_POST,
	Permission: utils.PERMISSION_ALL,
	Handler:    ngl,
}

func ngl(ctx *gin.Context) {
	var body utils.NGLParams

	// Parse request body
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Prepare payload
	payload := map[string]string{
		"username": body.Username,
		"question": body.Question,
		"deviceId": "",
		"referrer": "https://snapchat.com",
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	// Create request
	req, err := http.NewRequest("POST", "https://ngl.link/api/submit", bytes.NewBuffer(jsonData))
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64; rv:109.0) Gecko/20100101 Firefox/113.0")

	// Send request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	defer resp.Body.Close()

	// Read response
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	// Return response to client
	ctx.Data(resp.StatusCode, "application/json", respBody)
}
