package endpoints
import "github.com/gin-gonic/gin"

type Route struct {
	Method string
	Path string
	Handler gin.HandlerFunc
}

var Routes = []Route {
	{
		Method: "GET",
		Path: "/",
		Handler: Index,
	},
}
