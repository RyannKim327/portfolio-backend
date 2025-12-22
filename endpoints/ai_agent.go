package endpoints

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"portfolio-backend/utils"

	"github.com/gin-gonic/gin"
)

var PostAIAgent = utils.Route{
	Path:    "/ai/chat",
	Method:  "POST",
	Handler: aiagent,
}

type gptMessage struct {
	Role    string
	Content string
}

func aiagent(ctx *gin.Context) {
	var body gin.H

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	message_ := body["message"]

	message, _ := message_.(string)

	msgs := []gptMessage{}

	if len(msgs) <= 0 {
		msgs = append(msgs, gptMessage{
			Role:    "system",
			Content: "You are capable to use markdown.",
		})
	}

	msgs = append(msgs, gptMessage{
		Role:    "user",
		Content: message,
	})

	reqBody, _ := json.Marshal(map[string]interface{}{
		"model":       "openai-fast",
		"messages":    msgs,
		"temperature": 1,
		"max_tokens":  1000,
		"stream":      false,
	})

	// TODO: To Fetch data thru API
	req, _ := http.NewRequest("POST", "https://text.pollinations.ai/openai", bytes.NewBuffer(reqBody))

	client := &http.Client{}
	resp, _ := client.Do(req)

	defer resp.Body.Close()

	respBody, _ := io.ReadAll(resp.Body)

	// TODO: Decoder
	var apiResponse map[string]interface{}
	if err := json.Unmarshal(respBody, &apiResponse); err != nil {
		ctx.JSON(500, gin.H{
			"error": "Failed to parse response",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"response": apiResponse["choices"].([]interface{})[0].(map[string]interface{})["message"].(map[string]interface{})["content"],
	})
}
