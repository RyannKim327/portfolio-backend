package post

import (
	"portfolio-backend/utils"

	"github.com/gin-gonic/gin"
)

var Contact = utils.Route{
	Path:       "/contact",
	Method:     utils.METHOD_POST,
	Permission: utils.PERMISSION_COOKIE,
	Handler:    contact,
}

func contact(ctx *gin.Context) {
	var body gin.H
	file := "contact.json"

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	gist := utils.GistHandlerList(file)

	// maxID := 0
	// for _, g := range gist {
	// 	id := int(g["id"].(float64))
	// 	if id > maxID {
	// 		maxID = id
	// 	}
	// }
	//
	// now := time.Now().Format("01-02-2006")
	//
	// newBody := gin.H{
	// 	"id":      maxID + 1,
	// 	"title":   body["title"].(string),
	// 	"content": body["content"].(string),
	// 	"tags":    body["tags"],
	// 	"time":    now,
	// }

	gist = append(gist, body)
	response := utils.GistPostHandler(file, gist)

	ctx.JSON(200, gin.H{
		"from": response,
	})
}
