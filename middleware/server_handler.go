package middleware

import (
	"fmt"

	utils "portfolio-backend/utils"

	"github.com/gin-gonic/gin"
)

var app = gin.Default()

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

	switch routes.Method {
	case "GET":
		app.GET(routes.Path, routes.Handler)
	case "POST":
		app.POST(routes.Path, routes.Handler)
	case "PUT":
		app.PUT(routes.Path, routes.Handler)
	case "DELETE":
		app.DELETE(routes.Path, routes.Handler)
	}
}

func StartServer(port ...int) {
	p := 8000
	if len(port) > 0 {
		p = port[0]
	}

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

	app.Run(fmt.Sprintf(":%d", p))
}
