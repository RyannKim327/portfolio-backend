package utils

import (
	"io"

	"github.com/gin-gonic/gin"
)

// TODO: Just create a structure inside and call it outside

type Route struct {
	Method  string
	Path    string
	Handler gin.HandlerFunc
}

type GistResponseHandler struct {
	Error    error
	Response []byte
}

type AccessApi struct {
	Method string
	URL    string
	Body   io.Reader
}
