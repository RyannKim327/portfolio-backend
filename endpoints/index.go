package endpoints

import (
	utils "portfolio-backend/utils"

	"github.com/gin-gonic/gin"
)

/*
 * TODO: This is just a template for the other endpoint
 */

var Index = utils.Route{
	Path: "/",
	Handler: func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "The server is now active",
		})
	},
}
