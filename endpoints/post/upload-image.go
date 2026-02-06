package post

import (
	"portfolio-backend/utils"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var _ = godotenv.Load()

var UploadImage = utils.Route{
	Path:       "/upload",
	Method:     utils.METHOD_POST,
	Permission: utils.PERMISSION_ADMIN,
	Handler:    uploadImage,
}

func uploadImage(ctx *gin.Context) {
	// expectedKey := os.Getenv("POST_API")
	// file_url := fmt.Sprintf("https://api.telegram.org/file/bot%s/sendPhoto", expectedKey)

	// resp := access(utils.AccessAPI{
	// 	Method: "POST",
	// 	URL:    file_url,
	// 	Body:   bytes.NewBuffer(jsonBody),
	// })
	ctx.JSON(200, gin.H{
		"from": "Hi",
	})
}
