package post

import (
	"portfolio-backend/utils"

	"github.com/gin-gonic/gin"
)

func poetry(ctx *gin.Context) {
	var body gin.H

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	gist := utils.GistHandlerList("poetry.json")
	gist = append(gist, body)
	response := utils.GistPostHandler("poetry.json", gist)

	ctx.JSON(200, gin.H{
		"from": response,
	})
}

var Poetry = utils.Route{
	Path:       "/poetry",
	Method:     utils.METHOD_POST,
	Permission: utils.PERMISSION_ADMIN,
	Handler:    poetry,
}
