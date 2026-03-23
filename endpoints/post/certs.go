package post

import (
	"portfolio-backend/utils"

	"github.com/gin-gonic/gin"
)

var Certificates = utils.Route{
	Path:       "/certs",
	Method:     utils.METHOD_POST,
	Permission: utils.PERMISSION_ADMIN,
	Handler:    certificate,
}

func certificate(ctx *gin.Context) {
	var body gin.H
	file := "certificates.json"

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	gist := utils.GistHandlerList(file)
	gist = append(gist, body)
	response := utils.GistPostHandler(file, gist)

	ctx.JSON(200, gin.H{
		"from": response,
	})
}
