package main

import (
	endpoints "portfolio-backend/endpoints"
	mw "portfolio-backend/middleware"
	"portfolio-backend/utils"
)

func main() {
	/*
	 * TODO: There's nothing to do with this file
	 * Everything is set
	 */

	utils.GistHandler()

	for _, routes := range endpoints.Routes {
		mw.Register(routes)
	}

	mw.StartServer(8000)
}
