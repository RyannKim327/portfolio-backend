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

// TODO: Private Functions
func getPort(port []int) int {
	// TODO: Default

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

	// TODO: Randomized Port
	// fmt.Println("Trying for random Port")
	// if n, err := rand.Int(rand.Reader, big.NewInt(10000)); err == nil {
	// 	return int(n.Int64()) + 10000
	// }

	return 8000
}

// TODO: Public Functions

func Get(endpoint string, function gin.HandlerFunc) {
	app.GET(endpoint, function)
}

func Post(endpoint string, function gin.HandlerFunc) {
	app.POST(endpoint, function)
}

func Register(routes utils.Route) {
	/*
	 * INFO: To register each of endpoint from endpoint package/folder
	 * The purpose is to make the files cleaner and better looking
	 */
	if routes.Method == "" {
		routes.Method = "GET"
	}

	// TODO: To make the "Method" case insensitive
	routes.Method = strings.ToLower(routes.Method)

	switch routes.Method {
	case "get":
		app.GET(routes.Path, routes.Handler)
	case "post":
		app.POST(routes.Path, routes.Handler)
	case "put":
		app.PUT(routes.Path, routes.Handler)
	case "delete":
		app.DELETE(routes.Path, routes.Handler)
	}
}

func StartServer(port ...int) {
	p := getPort(port)

	app.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		// AllowOrigins:    []string{"https://ryannkim327.is-a.dev"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
		// AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		// ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
	}))

	// app.Use(cors.Default())

	app.Use(func(ctx *gin.Context) {
		// TODO: To add manual LOGS to monitor by data or file

		path := ctx.Request.URL.Path
		method := ctx.Request.Method

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
			</html>
		`
		ctx.Data(404, "text/html; charset=utf-8", []byte(html))
	})

	fmt.Printf("Running in PORT %d\n\n", p)
	app.Run(fmt.Sprintf(":%d", p))
}
