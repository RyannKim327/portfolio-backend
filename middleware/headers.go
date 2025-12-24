package middleware

import "github.com/gin-gonic/gin"

func ResponseHeaderMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Header("Developer", "MPOP Reverse II")
		ctx.Next()
	}
}
