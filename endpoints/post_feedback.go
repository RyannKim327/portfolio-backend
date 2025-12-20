package endpoints

import (
	"portfolio-backend/utils"

	"github.com/gin-gonic/gin"
)

var PostFeedback = utils.Route{
	Path:    "/feedback",
	Method:  "GET",
	Handler: main,
}

func main(ctx *gin.Context) {
	gist := utils.GistHandlerList("feedback.json")

	gist = append(gist, gin.H{
		"user":     "Test",
		"feedback": "This is a test",
	})

	response := utils.GistPostHandler("feedback.json", gist)

	ctx.JSON(200, gin.H{
		"response": response.Response,
		"error":    response.Error,
	})
}
