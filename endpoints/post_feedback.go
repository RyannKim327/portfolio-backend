package endpoints

import (
	"portfolio-backend/utils"

	"github.com/gin-gonic/gin"
)

var PostFeedback = utils.Route{
	Path:   "/feedback",
	Method: "POST",
	Handler: func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"response": "Response",
		})
	},
}
