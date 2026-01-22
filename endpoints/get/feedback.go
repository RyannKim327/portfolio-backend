package get

import (
	"reflect"
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

	ctx.JSON(200, gin.H{
		"count": len(cached),
		"data":  cached,
	})
}
