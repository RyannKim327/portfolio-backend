package post

import (
	"portfolio-backend/utils"

	"github.com/gin-gonic/gin"
)

var Project = utils.Route{
	Path:       "/project",
	Method:     utils.METHOD_POST,
	Permission: utils.PERMISSION_ADMIN,
	Handler:    project,
}

func project(ctx *gin.Context) {
	var body gin.H
	file := "projects.json"

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	gist := utils.GistHandler(file)

	projects, ok := gist["projects"].([]interface{})
	if !ok {
		projects = []interface{}{}
	}

	projects = append(projects, body)

	gist["projects"] = projects
	response := utils.GistPostHandler(file, gist)

	ctx.JSON(200, gin.H{
		"from": response,
	})
}
