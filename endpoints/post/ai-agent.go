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

var fallbackModels = []string{
	"google/gemma-4-26b-a4b-it:free",
	"google/gemma-4-26b-a4b-it",
}

type aiResponse struct {
	Role    string
	Content string
}

func aiagent(ctx *gin.Context) {
	var body utils.BodyAIStructure

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	msgs := []utils.GPTMessage{
		{
			Role:    "system",
			Content: "You are capable to use markdown. Just response in very detailed but readable.",
		},
	}
	msgs = append(msgs, body.Messages...)

	var lastErr error
	for _, model := range fallbackModels {
		res, err := requestAI(msgs, model)
		if err == nil {
			ctx.JSON(http.StatusOK, gin.H{
				"role":    res.Role,
				"content": res.Content,
			})
			return
		}
		lastErr = err
	}

	ctx.JSON(http.StatusInternalServerError, gin.H{
		"error": fmt.Sprintf("All AI models failed. Last error: %v", lastErr),
	})
}

func requestAI(msgs []utils.GPTMessage, model string) (*aiResponse, error) {
	reqBody, err := json.Marshal(map[string]any{
		"model":    model,
		"messages": msgs,
		"stream":   false,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to encode request: %w", err)
	}

	req, err := http.NewRequest("POST", "https://openrouter.ai/api/v1/chat/completions", bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("AI_API")))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request error for model %s: %w", model, err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response for model %s: %w", model, err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("model %s returned status %d: %s", model, resp.StatusCode, string(respBody))
	}

	var apiResponse map[string]interface{}
	if err := json.Unmarshal(respBody, &apiResponse); err != nil {
		return nil, fmt.Errorf("failed to parse JSON response for model %s: %w", model, err)
	}

	if errObj, exists := apiResponse["error"]; exists && errObj != nil {
		return nil, fmt.Errorf("API error for model %s: %v", model, errObj)
	}

	choices, ok := apiResponse["choices"].([]interface{})
	if !ok || len(choices) == 0 {
		return nil, fmt.Errorf("invalid or empty choices for model %s", model)
	}

	choice, ok := choices[0].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid choice format for model %s", model)
	}

	message, ok := choice["message"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid message format for model %s", model)
	}

	role, _ := message["role"].(string)
	if role == "" {
		role = "assistant"
	}

	contentStr, ok := message["content"].(string)
	if !ok {
		return nil, fmt.Errorf("invalid content string for model %s", model)
	}

	pattern := `---\n\*Support Pollinations\.AI:\*\n---\n🌸 \*Ad\* 🌸\nPowered by Pollinations\.AI free text APIs\. \[Support our mission\]\(https:\/\/pollinations\.ai\/redirect\/kofi\) to keep AI accessible for everyone\.`
	re := regexp.MustCompile(pattern)
	clean := re.ReplaceAllString(contentStr, "")

	return &aiResponse{
		Role:    role,
		Content: clean,
	}, nil
}
