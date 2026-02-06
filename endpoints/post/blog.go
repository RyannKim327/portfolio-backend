package post

import (
	"portfolio-backend/utils"

	"github.com/gin-gonic/gin"
)

var Blog = utils.Route{
	Path:       "/feedback",
	Method:     utils.METHOD_POST,
	Permission: utils.PERMISSION_ADMIN,
	Handler:    blog,
}

func blog(ctx *gin.Context) {
	var body gin.H

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	gist := utils.GistHandlerList("blog.json")
	gist = append(gist, body)
	response := utils.GistPostHandler("blog.json", gist)

	ctx.JSON(200, gin.H{
		"from": response,
	})
}
