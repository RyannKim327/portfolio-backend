package put

import (
	"strconv"

	"portfolio-backend/utils"

	"github.com/gin-gonic/gin"
)

var Blog = utils.Route{
	Path:       "blog",
	Method:     utils.METHOD_PUT,
	Permission: utils.PERMISSION_ADMIN,
	Handler:    put_blog,
}

func put_blog(ctx *gin.Context) {
	var body gin.H

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	id, ok := body["id"].(string)
	if !ok {
		ctx.JSON(200, gin.H{"error": "invalid id"})
		return
	}

	num, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(200, gin.H{"error": "invalid id"})
		return
	}

	gist := utils.GistHandlerList("blog.json")
	body["id"] = num
	gist[num-1] = body

	response := utils.GistPostHandler("blog.json", gist)

	ctx.JSON(200, gin.H{
		"from": response,
	})
}
