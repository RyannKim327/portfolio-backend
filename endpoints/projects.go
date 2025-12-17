package endpoints

import (
	"portfolio-backend/utils"

	"github.com/gin-gonic/gin"
)

var Projects = utils.Route{
	Path:   "/projects",
	Method: "GET",
	Handler: func(ctx *gin.Context) {
		data := utils.GistHandler("projects.json")
		ctx.JSON(200, data)
	},
}
