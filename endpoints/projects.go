package endpoints

import (
	"portfolio-backend/utils"

	"github.com/gin-gonic/gin"
)

var GetProjects = utils.Route{
	Path:   "/projects",
	Method: utils.METHOD_GET,
	Handler: func(ctx *gin.Context) {
		data := utils.GistHandler("projects.json")
		ctx.JSON(200, data)
	},
}
