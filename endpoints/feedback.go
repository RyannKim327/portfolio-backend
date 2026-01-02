package endpoints

import (
	"portfolio-backend/utils"

	"github.com/gin-gonic/gin"
)

var GetFeedback = utils.Route{
	Path:   "/feedback",
	Method: utils.METHOD_GET,
	Handler: func(ctx *gin.Context) {
		data := utils.GistHandlerList("feedback.json")
		utils.Reverse(data)
		ctx.JSON(200, gin.H{
			"count": len(data),
			"data":  data,
		})
	},
}
