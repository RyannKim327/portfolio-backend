package get

import (
	"strings"

	utils "portfolio-backend/utils"

	"github.com/gin-gonic/gin"
)

/*
 * TODO: This is just a template for the other endpoint
 */

// type DownloadStruct struct {
// 	Name         string
// 	Description  string
// 	DownloadLink string
// }

var Download = utils.Route{
	Path:    "download",
	Handler: download,
}

func download(ctx *gin.Context) {
	id := strings.ToLower(ctx.Query("id"))

	data := utils.GistHandler("downloads.json")
	gistcnt := utils.GistHandler("dlcounts.json")

	result, exists := data[id].(map[string]interface{})
	if !exists {
		ctx.JSON(404, gin.H{
			"error": "Download not found",
		})
		return
	}

	count, exists := gistcnt[id]
	if !exists {
		count = 0
	}

	result["downloads"] = count

	ctx.JSON(200, gin.H{
		"data": result,
	})
}
