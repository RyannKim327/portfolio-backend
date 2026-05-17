package post

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"

	"portfolio-backend/utils"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var _ = godotenv.Load()

var AIAgent = utils.Route{
	Path:       "/ai/chat",
	Method:     utils.METHOD_POST,
	Permission: utils.PERMISSION_ALL,
	Handler:    aiagent,
}

func aiagent(ctx *gin.Context) {
	var body utils.BodyAIStructure

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	msgs := []utils.GPTMessage{}

	// TODO: To add default formatting for AI Responses
	if len(msgs) <= 0 {
		msgs = append(msgs, utils.GPTMessage{
			Role:    "user",
			Content: "You are capable to use markdown. Just response in very detailed but readable.",
		})
		msgs = append(msgs, utils.GPTMessage{
			Role:    "assistant",
			Content: "Okay. I will, now what's on your mind?",
		})
	}

	// TODO: Basically to append two lists
	msgs = append(msgs, body.Messages...)

	reqBody, _ := json.Marshal(map[string]any{
		"model":    "openai/gpt-4o-mini",
		"messages": msgs,
		// "temperature": 1,
		// "max_tokens":  1000,
		"stream": false,
	})

	// TODO: To Fetch data thru API
	req, _ := http.NewRequest("POST", "https://openrouter.ai/api/v1/chat/completions", bytes.NewBuffer(reqBody))

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("AI_API")))
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

	choices, ok := apiResponse["choices"].([]interface{})

	if !ok || len(choices) <= 0 {
		ctx.JSON(500, gin.H{
			"error": apiResponse, // "Invalid API response due to choices formating",
		})
		return
	}

	choice, ok := choices[0].(map[string]interface{})

	if !ok || len(choices) <= 0 {
		ctx.JSON(500, gin.H{
			"error": "Invalid API response due to choices length",
		})
		return
	}

	message, ok := choice["message"].(map[string]interface{})
	if !ok || len(choices) <= 0 {
		ctx.JSON(500, gin.H{
			"error": "Invalid API response due to invalidity of message",
		})
		return
	}

	content := message["content"]

	pattern := `---\n\*Support Pollinations\.AI:\*\n---\n🌸 \*Ad\* 🌸\nPowered by Pollinations\.AI free text APIs\. \[Support our mission\]\(https:\/\/pollinations\.ai\/redirect\/kofi\) to keep AI accessible for everyone\.`

	re := regexp.MustCompile(pattern)

	str, _ := content.(string)

	clean := re.ReplaceAllString(str, "")

	ctx.JSON(200, gin.H{
		"role":    message["role"],
		"content": clean,
	})
}
