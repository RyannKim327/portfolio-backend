package get

import (
	utils "portfolio-backend/utils"

	"github.com/gin-gonic/gin"
)

/*
 * TODO: This is just a template for the other endpoint
 */

var Dev = utils.Route{
	Path:    "dev",
	Method:  utils.METHOD_GET,
	Handler: dev,
}

func dev(ctx *gin.Context) {
	data := utils.GistHandler("resume.json")
	ctx.JSON(200, gin.H{
		"data": data,
	})
	return
}
