package get

import (
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
	id := ctx.Query("id")

	data := utils.GistHandler("downloads.json")

	result, exists := data[id]
	if !exists {
		ctx.JSON(404, gin.H{
			"error": "Download not found",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"data": result,
	})
}
