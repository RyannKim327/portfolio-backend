package middleware

import (
	"github.com/gin-gonic/gin"
)

/*
 * TODO: This file is to handle the post request to prevent spam
 * for each feedback and other related stuffs to developer
 */

func RequestHandlerCookie() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		key, err := ctx.Cookie("temporary")

		if len(key) < 30 || err != nil {
			ctx.AbortWithStatusJSON(403, gin.H{
				"error": "Unauthorized",
			})
			return
		}

		ctx.Next()
	}
}
