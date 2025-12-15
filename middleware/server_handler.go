package middleware

import "github.com/gin-gonic/gin"
import "fmt"
import utils "portfolio-backend/utils"

var app = gin.Default()

func Get(endpoint string, function gin.HandlerFunc){
	app.GET(endpoint, function)
}

func Post(endpoint string, function gin.HandlerFunc){
	app.POST(endpoint, function)
}

func Register(routes utils.Route){
	if routes.Method == "" {
		routes.Method = "GET"
	}
	
	switch(routes.Method) {
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

func StartServer(port ...int){

	p := 8000
	if len(port) > 0 {
		p = port[0]
	}

	app.NoRoute(func (ctx *gin.Context){

		ctx.JSON(404, gin.H{
			"error": "Miss ko na sya",
		})
	})
	
	app.Run(fmt.Sprintf(":%d", p))
}
