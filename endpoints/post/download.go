package post

import (
	"strings"

	utils "portfolio-backend/utils"

	"github.com/gin-gonic/gin"
)

var Download = utils.Route{
	Path:       "download",
	Handler:    download,
	Method:     utils.METHOD_POST,
	Permission: utils.PERMISSION_ADMIN,
}

func download(ctx *gin.Context) {
	file := "downloads.json"

	var body struct {
		Name         string `json:"name" binding:"required"`
		Description  string `json:"description" binding:"required"`
		DownloadLink string `json:"downloadLink" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	name := strings.ToLower(body.Name)

	newBody := gin.H{
		"name":         body.Name,
		"description":  body.Description,
		"downloadLink": body.DownloadLink,
	}

	gist := utils.GistHandler(file)

	gist[name] = newBody
	response := utils.GistPostHandler(file, gist)

	ctx.JSON(200, gin.H{
		"from": response,
	})
}
