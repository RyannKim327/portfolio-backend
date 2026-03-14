package get

import (
	utils "portfolio-backend/utils"

	"github.com/gin-gonic/gin"
)

/*
 * TODO: This is just a template for the other endpoint
 */

var Certificates = utils.Route{
	Path:    "certs",
	Method:  utils.METHOD_GET,
	Handler: certi_handler,
}

func certi_handler(ctx *gin.Context) {
	data := utils.GistHandlerList("certificates.json")
	utils.Reverse(data)
	ctx.JSON(200, gin.H{
		"count": len(data),
		"data":  data,
	})
}
