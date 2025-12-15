package endpoints

import (
	utils "portfolio-backend/utils"

	"github.com/gin-gonic/gin"
)

var Fetch = utils.Route{
	Path: "/fetch",
	Handler: func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "A successful connection",
		})
	},
}
