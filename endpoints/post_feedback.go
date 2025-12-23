package endpoints

import (
	"fmt"

	"portfolio-backend/utils"

	"github.com/gin-gonic/gin"
)

var PostFeedback = utils.Route{
	Path:    "/feedback",
	Method:  "POST",
	Handler: feedback,
}

func feedback(ctx *gin.Context) {
	var body gin.H

	fmt.Println("Feedback Posting")
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	gist := utils.GistHandlerList("feedback.json")
	gist = append(gist, body)
	response := utils.GistPostHandler("feedback.json", gist)

	ctx.JSON(200, gin.H{
		"from": response,
	})
}
