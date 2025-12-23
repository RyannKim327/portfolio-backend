package endpoints

import (
	"bytes"
	"encoding/json"
	"fmt"
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
	Role    string `json:"role"`
	Content string `json:"content"`
}

type bodyStruc struct {
	Messages []gptMessage `json:"messages"`
}

func aiagent(ctx *gin.Context) {
	var body bodyStruc

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	msgs := []gptMessage{}

	if len(msgs) <= 0 {
		msgs = append(msgs, gptMessage{
			Role:    "user",
			Content: "You are capable to use markdown. Just response in very detailed but readable.",
		})
		msgs = append(msgs, gptMessage{
			Role:    "assistant",
			Content: "Okay. I will, now what's on your mind?",
		})
	}

	msgs = append(msgs, body.Messages...)

	reqBody, _ := json.Marshal(map[string]any{
		"model":       "openai-fast",
		"messages":    msgs,
		"temperature": 1,
		"max_tokens":  1000,
		"stream":      false,
	})

	fmt.Println(bytes.NewBuffer(reqBody))

	// TODO: To Fetch data thru API
	req, _ := http.NewRequest("POST", "https://text.pollinations.ai/openai", bytes.NewBuffer(reqBody))

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

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
		"role":    apiResponse["choices"].([]interface{})[0].(map[string]interface{})["message"].(map[string]interface{})["role"],
		"content": apiResponse["choices"].([]interface{})[0].(map[string]interface{})["message"].(map[string]interface{})["content"],
	})
}
