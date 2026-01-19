package get

import (
	"portfolio-backend/utils"

	"github.com/gin-gonic/gin"
)

var Experiences = utils.Route{
	Path:   "/experiences",
	Method: utils.METHOD_GET,
	Handler: func(ctx *gin.Context) {
		data := utils.GistHandlerList("experiences.json")
		ctx.JSON(200, data)
	},
}
