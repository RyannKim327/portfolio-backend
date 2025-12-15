package main
import mw "portfolio-backend/middleware"
import endpoints "portfolio-backend/endpoints"


func main(){
	for _, routes := range endpoints.Routes {
		mw.Register(mw.Routes{routes.Method, routes.Path, routes.Handler})
	}
	mw.StartServer(8000)
}
