package post

import (
	"time"

	"portfolio-backend/utils"

	"github.com/gin-gonic/gin"
)

var Feedback = utils.Route{
	Path:       "/feedback",
	Method:     utils.METHOD_POST,
	Permission: utils.PERMISSION_COOKIE,
	Handler:    feedback,
}

func feedback(ctx *gin.Context) {
	var body gin.H
	file := "feedback.json"

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	gist := utils.GistHandlerList(file)

	newBody := gin.H{
		"application": body["application"].(string),
		"message":     body["message"].(string),
		"userId":      body["userId"],
		"date":        time.Now(),
	}

	gist = append(gist, newBody)
	response := utils.GistPostHandler(file, gist)

	ctx.JSON(200, gin.H{
		"from": response,
	})
}
