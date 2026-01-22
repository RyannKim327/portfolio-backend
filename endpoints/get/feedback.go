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
	CachedArrayContent []gin.H
	CacheTTL           time.Time
	CacheMU            sync.RWMutex
	CachedDuration     = (5 * time.Minute)
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

	CacheMU.RLock()
	cached := CachedArrayContent
	valid := time.Now().Before(CacheTTL) && cached != nil
	CacheMU.RUnlock()

	if valid {
		go func(old []gin.H) {
			data := utils.GistHandlerList("feedback.json")
			utils.Reverse(data)

			if !reflect.DeepEqual(old, data) {
				CacheMU.Lock()
				CachedArrayContent = data
				CacheTTL = time.Now().Add(CachedDuration)
				CacheMU.Unlock()
			}
		}(cached)

		ctx.JSON(200, gin.H{
			"count": len(cached),
			"data":  cached,
		})
		return
	}

	data := utils.GistHandlerList("feedback.json")
	utils.Reverse(data)

	CacheMU.Lock()
	CachedArrayContent = data
	CacheTTL = time.Now().Add(CachedDuration)
	CacheMU.Unlock()

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
	if start >= len(cached) {
		ctx.JSON(200, gin.H{
			"pages": 1,
			"count": len(cached),
			"data":  cached,
		})
		return
	}

	response := []interface{}{}

	for i := start; i < end; i++ {
		response = append(response, cached[i])
	}

	ctx.JSON(200, gin.H{
		"pages": pages,
		"count": len(cached),
		"data":  response,
	})
}
