package endpoints

import (
	utils "portfolio-backend/utils"

	"github.com/gin-gonic/gin"
)

var Index = utils.Route{
	Path: "/",
	Handler: func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"error": "Hello World",
		})
	},
}
