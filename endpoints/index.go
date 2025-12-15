package endpoints
import "github.com/gin-gonic/gin"

func Index(ctx *gin.Context){
	ctx.JSON(200, gin.H{
		"error": "Hello World",
	})	
}
