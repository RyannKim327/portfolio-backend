package post

import (
	"time"

	"portfolio-backend/utils"

	"github.com/gin-gonic/gin"
)

var Blog = utils.Route{
	Path:       "/blog",
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

	maxID := 0
	for _, g := range gist {
		id := int(g["id"].(float64))
		if id > maxID {
			maxID = id
		}
	}

	now := time.Now().Format("01-02-2006")

	newBody := gin.H{
		"id":      maxID + 1,
		"title":   body["title"].(string),
		"content": body["content"].(string),
		"tags":    body["tags"],
		"time":    now,
	}

	gist = append(gist, newBody)
	response := utils.GistPostHandler("blog.json", gist)

	ctx.JSON(200, gin.H{
		"from": response,
	})
}
