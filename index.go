package main

import (
	endpoints "portfolio-backend/endpoints"
	mw "portfolio-backend/middleware"
)

func main() {
	/*
	 * TODO: There's nothing to do with this file
	 * Everything is set
	 */

	mw.CorsSetuo()

	for _, routes := range endpoints.Routes {
		mw.Register(routes)
	}

	mw.StartServer()
}
