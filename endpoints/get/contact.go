package get

import (
	"portfolio-backend/utils"

	"github.com/gin-gonic/gin"
)

var Contact = utils.Route{
	Path:       "/contact",
	Method:     utils.METHOD_GET,
	Permission: utils.PERMISSION_ADMIN,
	Handler: func(ctx *gin.Context) {
		data := utils.GistHandlerList("contact.json")
		ctx.JSON(200, gin.H{
			"count": len(data),
			"data":  data,
		})
	},
}
