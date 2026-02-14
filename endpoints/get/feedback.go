package get

import (
	"reflect"
	"strconv"
	"sync"
	"time"

	"portfolio-backend/utils"

	"github.com/gin-gonic/gin"
)

var (
	FeedbackCachedArrayContent []gin.H
	FeedbackCacheTTL           time.Time
	FeedbackCacheMU            sync.RWMutex
	FeedbackCachedDuration     = (5 * time.Minute)
)

var Feedback = utils.Route{
	Path:    "/feedback",
	Method:  utils.METHOD_GET,
	Handler: feedback_handler,
}

func feedback_handler(ctx *gin.Context) {
	page := 1
	limit := 10

	if ctx.Query("page") != "" {
		if p, err := strconv.Atoi(ctx.Query("page")); err == nil && p > 0 {
			page = p
		}
	}

	FeedbackCacheMU.RLock()
	cached := FeedbackCachedArrayContent
	valid := time.Now().Before(FeedbackCacheTTL) && cached != nil
	FeedbackCacheMU.RUnlock()

	if valid {
		go func(old []gin.H) {
			data := utils.GistHandlerList("feedback.json")
			utils.Reverse(data)

			if !reflect.DeepEqual(old, data) {
				FeedbackCacheMU.Lock()
				FeedbackCachedArrayContent = data
				FeedbackCacheTTL = time.Now().Add(FeedbackCachedDuration)
				FeedbackCacheMU.Unlock()
			}
		}(cached)

		// TODO: This is just to add how many pages
		pages := len(cached) / limit

		if pages < page || page <= 1 {
			// TODO: To prevent out bound exception error
			page = 1
		}

		// TODO: Start of pagination
		start := limit * (page - 1)
		end := start + limit

		if end > len(cached) {
			end = len(cached)
		}

		// TODO: Condition of paginator
		if start >= len(cached) {
			ctx.JSON(200, gin.H{
				"pages":   1,
				"current": page,
				"count":   len(cached),
				"data":    cached,
			})
			return
		}

		response := []interface{}{}

		for i := start; i < end; i++ {
			response = append(response, cached[i])
		}

		ctx.JSON(200, gin.H{
			"pages":   pages,
			"current": page,
			"count":   len(cached),
			"data":    response,
		})
		return
	}

	data := utils.GistHandlerList("feedback.json")
	utils.Reverse(data)

	FeedbackCacheMU.Lock()
	FeedbackCachedArrayContent = data
	FeedbackCacheTTL = time.Now().Add(FeedbackCachedDuration)
	FeedbackCacheMU.Unlock()

	// TODO: This is just to add how many pages
	pages := len(cached) / limit

	if pages < page {
		// TODO: To prevent out bound exception error
		page = 1
	}

	// TODO: Start of pagination
	start := limit * (page - 1)
	end := start + limit

	// TODO: Condition of paginator
	if start >= len(cached) && cached != nil {
		ctx.JSON(200, gin.H{
			"pages":   1,
			"current": page,
			"count":   len(cached),
			"data":    cached,
		})
		return
	}
	response := []gin.H{}

	if len(cached) > 0 {
		for i := start; i < end; i++ {
			response = append(response, cached[i])
		}
	} else {
		response = data
	}

	ctx.JSON(200, gin.H{
		"pages":   pages,
		"current": page,
		"count":   len(response),
		"data":    response,
	})
}
