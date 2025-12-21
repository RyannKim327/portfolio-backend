package endpoints

import (
	"portfolio-backend/utils"

	"github.com/gin-gonic/gin"
)

var Poetry = utils.Route{
	Path:   "/poetry",
	Method: "GET",
	Handler: func(ctx *gin.Context) {
		data := utils.GistHandlerList("poetry.json")
		utils.Reverse(data)
		ctx.JSON(200, data)
	},
}

