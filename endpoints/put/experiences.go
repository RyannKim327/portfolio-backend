package put

import (
	"fmt"

	"portfolio-backend/utils"

	"github.com/gin-gonic/gin"
)

var Experience = utils.Route{
	Path:       "/experiences",
	Method:     utils.METHOD_PUT,
	Permission: utils.PERMISSION_ADMIN,
	Handler:    exp,
}

func exp(ctx *gin.Context) {
	var body []gin.H

	if err := ctx.ShouldBindJSON(&body); err != nil {
		fmt.Println(err)
		ctx.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	// gist := utils.GistHandlerList("experiences.json")
	response := utils.GistPostHandler("experiences.json", body)

	ctx.JSON(200, gin.H{
		"from": response,
	})
}
