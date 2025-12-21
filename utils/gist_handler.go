package utils

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
)

func GistHandlerList(file string) []gin.H {
	if !strings.Contains(file, ".") || strings.HasSuffix(file, ".") {
		// TODO: To auto add .json suffix to the file (string)
		file = file + ".json"
	}

	gist := Get()
	response, ok := gist.Response.Files[file]

	// TODO: To check if the response is still on
	if !ok {
		return []gin.H{
			{
				"error": "Have some bugs",
			},
		}
	}

	// TODO: Interpretation to JSON format
	var parse []gin.H

	err := json.Unmarshal([]byte(response.Content), &parse)
	if err != nil {
		return []gin.H{
			{
				"error": err,
			},
		}
	}

	return parse
}

func GistHandler(file string) gin.H {
	if !strings.Contains(file, ".") || strings.HasSuffix(file, ".") {
		// TODO: To auto add .json suffix to the file (string)
		file = file + ".json"
	}

	gist := Get()
	response, ok := gist.Response.Files[file]

	// TODO: To check if the response is still on
	if !ok {
		return gin.H{
			"error": "Have Error",
		}
	}

	fmt.Println(response.Content)
	// TODO: Interpretation to JSON format
	var parse map[string]interface{}

	err := json.Unmarshal([]byte(response.Content), &parse)
	if err != nil {
		return gin.H{
			"error": err.Error(),
		}
	}

	return parse
}

func GistPostHandler(file string, data interface{}) GistResponseHandler {
	if !strings.Contains(file, ".") || strings.HasSuffix(file, ".") {
		// TODO: To auto add .json suffix to the file (string)
		file = file + ".json"
	}

	contentBytes, _ := json.MarshalIndent(data, "", " ")

	data_ := gin.H{
		"files": gin.H{
			file: gin.H{
				"content": string(contentBytes),
			},
		},
	}

	gist := Post(data_)
	return gist
}
