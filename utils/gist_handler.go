package utils

import "fmt"

func GistHandler() {
	gist := Get()
	response := gist.Response
	fmt.Println(response["files"])
	// fmt.Println(gist)
}
