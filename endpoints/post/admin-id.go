package post

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
	"time"

	utils "portfolio-backend/utils"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var _ = godotenv.Load()

var AdminSide = utils.Route{
	Path:       "/bakitkosasabihinsayo",
	Method:     utils.METHOD_POST,
	Permission: utils.PERMISSION_ALL,
	Handler:    adminSide,
}

func adminSide(ctx *gin.Context) {
	var body utils.AdminSideStructure

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	env_admin := os.Getenv("ADMIN_SIDE")
	post_admin := os.Getenv("POST_API")

	fmt.Println(env_admin)
	fmt.Println(body.Key)

	if body.Key == env_admin {
		data := gin.H{
			"code": post_admin,
			"time": time.Now().UnixMilli() + 86400000,
		}

		jsonData, err := json.Marshal(data)
		if err != nil {
			ctx.JSON(500, gin.H{
				"error": "Marshal Law Error",
			})
			return
		}

		code := base64.StdEncoding.EncodeToString([]byte(jsonData))

		ctx.JSON(200, gin.H{
			"code": code,
		})
		return
	}

	ctx.JSON(403, gin.H{
		"error": "Unauthorized Access",
	})
}
