package get

import (
	"maps"
	"slices"
	"strconv"

	"portfolio-backend/utils"

	"github.com/gin-gonic/gin"
)

var DownloadLists = utils.Route{
	Path:       "dl-lists",
	Method:     utils.METHOD_GET,
	Permission: utils.PERMISSION_ALL,
	Handler:    downloadLists,
}

func downloadLists(ctx *gin.Context) {
	page := 1
	limit := 5

	if ctx.Query("page") != "" {
		if p, err := strconv.Atoi(ctx.Query("page")); err == nil && p > 0 {
			page = p
		}
	}

	if ctx.Query("limit") != "" {
		if l, err := strconv.Atoi(ctx.Query("limit")); err == nil && l > 0 {
			limit = l
		}
	}

	data := utils.GistHandler("downloads.json")
	keys := slices.Collect(maps.Keys(data))

	result := make([]gin.H, 0)

	// TODO: Start of pagination
	start := limit * (page - 1)
	end := start + limit

	if end > len(keys) {
		end = len(keys)
	}

	// TODO: Loop Process
	for i := len(keys) - 1; i >= start && i < end; i-- {
		key := keys[i]
		value := data[key]

		item := value.(map[string]any)

		result = append(result, gin.H{
			"id":   key,
			"name": item["name"],
		})
	}

	ctx.JSON(200, gin.H{
		"pages":   len(data) / limit,
		"current": page,
		"count":   len(data),
		"data":    result,
	})
}
