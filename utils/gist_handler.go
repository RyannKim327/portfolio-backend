package utils

import (
	"encoding/json"
	"strings"

	"github.com/gin-gonic/gin"
)

func GistHandler(file string) gin.H {
	if !strings.Contains(file, ".") && strings.HasSuffix(file, ".") {
		// TODO: To auto add .json suffix to the file (string)
		file = file + ".json"
	}

	gist := Get()
	response, ok := gist.Response.Files[file]

	// TODO: To check if the response is still on
	if !ok {
		return gin.H{
			"error": "Have some bugs",
		}
	}

	// TODO: Interpretation to JSON format
	var parse map[string]interface{}

	err := json.Unmarshal([]byte(response.Content), &parse)
	if err != nil {
		return gin.H{
			"error": err,
		}
	}

	return parse
}
