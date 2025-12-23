package middleware

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func PostRequestHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		key := ctx.GetHeader("X-API-Key")
		expectedKey := os.Getenv("POST_API")
		fmt.Println("Hello Lord")
		if key == "" || key != expectedKey {
			ctx.AbortWithStatusJSON(403, gin.H{
				"error": "Unauthorized",
			})
			return
		}
		ctx.Next()
	}
}
