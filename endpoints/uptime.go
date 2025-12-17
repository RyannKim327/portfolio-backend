package endpoints

import (
	utils "portfolio-backend/utils"

	"github.com/gin-gonic/gin"
)

var Uptime = utils.Route{
	Path: "/uptime",
	Handler: func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "A successful connection",
		})
	},
}
