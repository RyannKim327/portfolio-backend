package endpoints
import "github.com/gin-gonic/gin"
import utils "portfolio-backend/utils"

var Index = utils.Route {
	Path: "/",
	Handler: func(ctx *gin.Context){
		ctx.JSON(200, gin.H{
			"error": "Hello World",
		})	
	},
}
