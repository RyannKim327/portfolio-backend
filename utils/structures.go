package utils

import (
	"io"

	"github.com/gin-gonic/gin"
)

// TODO: Just create a structure inside and call it outside

type Route struct {
	Method     string
	Path       string
	Permission string
	Handler    gin.HandlerFunc
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

type AccessAPI struct {
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

type AdminSideStructure struct {
	Key string `json:"key"`
}

type NGLParams struct {
	Question string `json:"question"`
	Username string `json:"username"`
}

type TelegramResponse struct {
	ChatID   string `json:"chat_id"`
	Document string `json:"document"`
}

type TelegramAPI struct {
	Response TelegramResponse
	Error    error
}

type AccessTelegramAPI struct {
	Method string
	URL    string
	Body   io.Reader
	File   string
}

type BaybayinCharacters = map[string]int
