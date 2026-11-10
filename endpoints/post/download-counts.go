package post

import (
	"strings"

	"portfolio-backend/utils"

	"github.com/gin-gonic/gin"
)

var DownloadCounts = utils.Route{
	Path:       "dlcounts",
	Method:     utils.METHOD_POST,
	Permission: utils.PERMISSION_ALL,
	Handler:    dlcounts,
}

func dlcounts(ctx *gin.Context) {
	id := ctx.Query("id")

	file := "dlcounts.json"
	gist := utils.GistHandler(file)
	name := strings.ToLower(id)

	// This is to validate that the data is existing
	data := utils.GistHandler("downloads.json")

	_, exists := data[name]
	if !exists {
		ctx.JSON(404, gin.H{
			"error": "Data not found",
		})
		return
	}

	count, exists := gist[name]
	if !exists {
		// This is to set a default value
		count = float64(0)
	}

	// Conversion from any to int
	count = count.(float64) + 1

	// To update the count
	gist[name] = count
	response := utils.GistPostHandler(file, gist)

	ctx.JSON(200, gin.H{
		"from":    response,
		"message": "Updated",
	})
}
