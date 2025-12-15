package main

import (
	endpoints "portfolio-backend/endpoints"
	mw "portfolio-backend/middleware"
)

func main() {
	for _, routes := range endpoints.Routes {
		mw.Register(routes)
	}

	mw.StartServer(8000)
}
