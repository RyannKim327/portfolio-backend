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

type BLOG_GIST struct {
	Id      int      `json:"id"`
	Title   string   `json:"title"`
	Content string   `json:"content"`
	Tags    []string `json:"tags"`
	Time    string   `json:"time"`
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

	newBody := gin.H{
		"id":      maxID + 1,
		"title":   body["title"].(string),
		"content": body["content"].(string),
		"tags":    body["tags"],
	}

	gist = append(gist, newBody)
	response := utils.GistPostHandler("blog.json", gist)

	ctx.JSON(200, gin.H{
		"from": response,
	})
}
