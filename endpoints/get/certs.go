package get

import (
	"strconv"

	utils "portfolio-backend/utils"

	"github.com/gin-gonic/gin"
)

/*
 * TODO: This is just a template for the other endpoint
 */

var Certificates = utils.Route{
	Path:    "certs",
	Method:  utils.METHOD_GET,
	Handler: certi_handler,
}

func certi_handler(ctx *gin.Context) {
	page := 1
	limit := 8

	if ctx.Query("limit") != "" {
		if l, err := strconv.Atoi(ctx.Query("limit")); err == nil && l > 0 {
			limit = l
		}
	}

	if ctx.Query("page") != "" {
		if p, err := strconv.Atoi(ctx.Query("page")); err == nil && p > 0 {
			page = p
		}
	}

	data := utils.GistHandlerList("certificates.json")
	utils.Reverse(data)

	// TODO: This is just to add how many pages
	pages := (len(data) / limit) + 1

	if pages < page {
		// TODO: To prevent out bound exception error
		page = 1
	}

	// TODO: Start of pagination
	total := len(data)
	start := limit * (page - 1)
	end := start + limit

	// TODO: To create a version for last page
	if end > total {
		end = (total - (start - end)) - limit
	}

	// TODO: This is for furture debugging
	// fmt.Printf("Start %d\n", start)
	// fmt.Printf("End %d\n", end)
	// fmt.Printf("Limit %d\n", limit)

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

	ctx.JSON(200, gin.H{
		"total":   total,
		"pages":   pages,
		"current": page,
		"count":   len(response),
		"data":    response,
	})
}
