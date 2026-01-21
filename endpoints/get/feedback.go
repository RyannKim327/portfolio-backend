package get

import (
	"sync"
	"time"

	"portfolio-backend/utils"

	"github.com/gin-gonic/gin"
)

var (
	CachedArrayContent []gin.H
	CacheTTL           time.Time
	CacheMU            sync.Mutex
	CachedDuration     = (5 * time.Minute)
	NextRef            time.Time
)

var Feedback = utils.Route{
	Path:    "/feedback",
	Method:  utils.METHOD_GET,
	Handler: feedback_handler,
}

func feedback_handler(ctx *gin.Context) {
	CacheMU.Lock()
	cached := CachedArrayContent
	now := time.Now()

	if cached != nil || now.After(NextRef) {
		data := utils.GistHandlerList("feedback.json")
		utils.Reverse(data)

		CachedArrayContent = data
		NextRef = now.Add(CachedDuration)
	}

	data := CachedArrayContent
	CacheMU.Unlock()

	ctx.JSON(200, gin.H{
		"count": len(data),
		"data":  data,
	})
}
