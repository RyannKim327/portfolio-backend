package middleware

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	utils "portfolio-backend/utils"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var app = gin.Default()

// TODO: CORS Setup
func CorsSetup() {
	app.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"https://ryannkim327.is-a.dev"}, // "http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// TODO: To setup Headers
	app.Use(ResponseHeaderMiddleware())
}

// TODO: Private Functions
func getPort(port []int) int {
	// TODO: Port from .env
	fmt.Println("Trying to fetch .env PORT")
	if port_ := os.Getenv("PORT"); port_ != "" {
		if v, err := strconv.Atoi(port_); err == nil {
			return v
		}
	}

	fmt.Println("Trying to fetch dev input PORT")
	if len(port) > 0 {
		return port[0]
	}

	return 8000
}

// TODO: Public Functions

func Register(routes utils.Route) {
	/*
	 * INFO: To register each of endpoint from endpoint package/folder
	 * The purpose is to make the files cleaner and better looking
	 */

	// TODO: Make GET Request as default
	if routes.Method == "" {
		routes.Method = utils.METHOD_GET
	}

	// TODO: To make the "Method" case insensitive
	method := strings.ToUpper(routes.Method)

	// TODO: To make sure that the endpoint is in lowercase
	routes.Path = strings.ToLower(routes.Path)

	// TODO: To automatically add "/submit" in every required post requests
	if method == strings.ToUpper(utils.METHOD_POST) && !strings.HasSuffix(routes.Path, "/submit") {

		// TODO: To prevent double slash in endpoints
		if !strings.HasSuffix(routes.Path, "/") {
			routes.Path += "/"
		}

		routes.Path += "submit"
	}

	// TODO: To setup default permission
	if routes.Permission == "" {
		if routes.Method == utils.METHOD_POST {
			routes.Permission = utils.PERMISSION_COOKIE
		} else {
			routes.Permission = utils.PERMISSION_ALL
		}
	}

	switch method {
	case utils.METHOD_GET:
		switch routes.Permission {
		case utils.PERMISSION_COOKIE:
			app.GET(routes.Path, RequestHandlerCookie(), routes.Handler)
		case utils.PERMISSION_ADMIN:
			app.GET(routes.Path, RequestHandlerAdmin(), routes.Handler)
		default:
			app.GET(routes.Path, routes.Handler)
		}
	case utils.METHOD_POST:
		switch routes.Permission {
		case utils.PERMISSION_COOKIE:
			app.POST(routes.Path, RequestHandlerCookie(), routes.Handler)
		case utils.PERMISSION_ADMIN:
			app.POST(routes.Path, RequestHandlerAdmin(), routes.Handler)
		default:
			app.POST(routes.Path, routes.Handler)
		}
	case utils.METHOD_PATCH:
		switch routes.Permission {
		case utils.PERMISSION_COOKIE:
			app.PATCH(routes.Path, RequestHandlerCookie(), routes.Handler)
		case utils.PERMISSION_ADMIN:
			app.PATCH(routes.Path, RequestHandlerAdmin(), routes.Handler)
		}
	case utils.METHOD_PUT:
		switch routes.Permission {
		case utils.PERMISSION_COOKIE:
			app.PUT(routes.Path, RequestHandlerCookie(), routes.Handler)
		case utils.PERMISSION_ADMIN:
			app.PUT(routes.Path, RequestHandlerAdmin(), routes.Handler)
		}
	case utils.METHOD_DELETE:
		switch routes.Permission {
		case utils.PERMISSION_COOKIE:
			app.DELETE(routes.Path, RequestHandlerCookie(), routes.Handler)
		case utils.PERMISSION_ADMIN:
			app.DELETE(routes.Path, RequestHandlerAdmin(), routes.Handler)
		}
	}
}

func StartServer(port ...int) {
	p := getPort(port)

	app.Use(func(ctx *gin.Context) {
		// TODO: To add manual LOGS to monitor by data or file

		path := ctx.Request.URL.Path
		method := ctx.Request.Method

		// TODO: Header setter

		ctx.Next()

		status := ctx.Writer.Status()
		// INFO: According to ChatGPT: Go’s time formatting uses a reference date instead of formatting codes (like YYYY or MM)
		fmt.Printf("[%s] [%s] %s -> %d\n", time.Now().Format("2006-01-02 15:04:05"), method, path, status)
	})

	app.NoRoute(func(ctx *gin.Context) {
		const html = `
			<html>
				<head>
					<title>Page Not Found</title>
					<style>
						body {
							display: flex;
							flex-direction: column;
							height: 100dvh;
							width: 100dvw;
							align-items: center;
							justify-content: center;
							background-color: #191919;
							color: white;
						}
						h1{
							font-size: 5em;
							margin: 0;
						}
						h3 {
							margin: 0;
							font-size: 3em;
						}
					</style>
				</head>
				<body>
					<h1>404</h1>
					<h3>Wag na ngani hanapin</h3>
					<p>Tapos kami sisisihin mo kapag di mo nakita</p>
				</body>
				<script>
					setTimeout(() => {
						location.href="https://ryannkim327.is-a.dev"
					}, 1000)
				</script>
			</html>
		`
		ctx.Data(404, "text/html; charset=utf-8", []byte(html))
	})

	fmt.Printf("Running in PORT %d\n\n", p)
	app.Run(fmt.Sprintf(":%d", p))
}
