package main
import mw "portfolio-backend/middleware"
import endpoints "portfolio-backend/endpoints"
// import utils "portfolio-backend/utils"

func main(){
	for _, routes := range endpoints.Routes {
		mw.Register(routes)
	}
	
	mw.StartServer(8000)

}
