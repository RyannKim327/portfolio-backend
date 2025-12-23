package middleware

import (
	"os"

	"github.com/gin-gonic/gin"
)

func PostRequestHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		key := ctx.GetHeader("X-API-Key")
		expectedKey := os.Getenv("POST_API")

		if key == "" || key != expectedKey {
			ctx.AbortWithStatusJSON(403, gin.H{
				"error": "Unauthorized",
			})
			return
		}
		ctx.Next()
	}
}
