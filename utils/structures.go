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

type GistFile struct {
	Content string `json:"content"`
}

type Gist struct {
	Files map[string]GistFile `json:"files"`
}

type GistResponseHandler struct {
	Error    error
	Response Gist
}

type AccessApi struct {
	Method string
	URL    string
	Body   io.Reader
}

type GPTMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type BodyAIStructure struct {
	Messages []GPTMessage `json:"messages"`
}
