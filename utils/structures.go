package utils

import (
	"io"

	"github.com/gin-gonic/gin"
)

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
