package endpoints

import (
	"crypto/rand"
	"encoding/base64"
	"net/http"

	utils "portfolio-backend/utils"

	"github.com/gin-gonic/gin"
)

/*
 * TODO: This is just a template for the other endpoint
 */

var GetCookie = utils.Route{
	Path:    "/set-cookie",
	Handler: setCookie,
}

func tokenGenerator(length int) (string, error) {
	b := make([]byte, length)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return base64.RawURLEncoding.EncodeToString(b), nil
}

func setCookie(ctx *gin.Context) {
	cookie, _ := ctx.Cookie("temporary")

	if cookie != "" {
		ctx.JSON(200, gin.H{
			"message": "You still have an active token",
		})
		return
	}

	value, err := tokenGenerator(30)
	if err != nil {
		ctx.JSON(403, gin.H{
			"error": "There's an internal problem. Please try again",
		})
		return
	}

	http.SetCookie(ctx.Writer, &http.Cookie{
		Name:     "temporary",
		Value:    value,
		Path:     "/",
		MaxAge:   1800,
		Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteNoneMode,
	})

	ctx.JSON(200, gin.H{
		"message": "Token is now generated, you may now continue",
	})
}
