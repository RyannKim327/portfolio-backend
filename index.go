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

	// TODO: Call this first before the registration of each routes
	mw.CorsSetup()

	for _, routes := range endpoints.Routes {
		mw.Register(routes)
	}

	mw.StartServer()
}
