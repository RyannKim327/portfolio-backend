package get

import (
	"strconv"

	"portfolio-backend/utils"

	"github.com/gin-gonic/gin"
)

var Blog = utils.Route{
	Path:    "/blog",
	Method:  utils.METHOD_GET,
	Handler: blog_handler,
}

func blog_handler(ctx *gin.Context) {
	page := 1
	limit := 9

	if ctx.Query("page") != "" {
		if p, err := strconv.Atoi(ctx.Query("page")); err == nil && p > 0 {
			page = p
		}
	}
	data := utils.GistHandlerList("blog.json")
	utils.Reverse(data)

	// TODO: This is just to add how many pages
	pages := len(data) / limit

	if pages < page {
		// TODO: To prevent out bound exception error
		page = 1
	}

	// TODO: Start of pagination
	start := limit * (page - 1)
	end := start + limit

	// TODO: Condition of paginator
	if start >= len(data) && data != nil {
		ctx.JSON(200, gin.H{
			"pages":   1,
			"current": page,
			"count":   len(data),
			"data":    data,
		})
		return
	}

	response := []gin.H{}

	if len(data) > 0 {
		for i := start; i < end; i++ {
			response = append(response, data[i])
		}
	} else {
		response = data
	}
	if ctx.Query("id") != "" {
		if id, err := strconv.Atoi(ctx.Query("id")); err == nil {
			if id > 0 {
				ctx.JSON(200, gin.H{
					"data": data[len(data)-id],
				})
				return
			}
		}
	}

	ctx.JSON(200, gin.H{
		"pages":   pages,
		"current": page,
		"count":   len(response),
		"data":    response,
	})
}
