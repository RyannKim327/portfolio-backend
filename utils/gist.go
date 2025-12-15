package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

var URL = fmt.Sprintf("https://api.github.com/gists/%s", os.Getenv("API_KEY"))

func access(params AccessApi) GistResponseHandler {
	/*
	 * TODO: To create a door to access gist to the backend and
	 * to distribute the data through the other platform connected
	 * to my domain
	 */
	if params.Method == "" {
		params.Method = "GET"
	}

	req, err := http.NewRequest(params.Method, params.URL, params.Body)
	if err != nil {
		return GistResponseHandler{Error: err}
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("API_KEY")))
	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("X-Github-Api-Version", "2022-11-28")

	client := &http.Client{}
	resp, err := client.Do(req)

	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return GistResponseHandler{Error: err}
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return GistResponseHandler{Error: err}
	}

	return GistResponseHandler{Response: body}
}

func Get(url string) GistResponseHandler {
	resp := access(AccessApi{
		URL: URL,
	})
	return resp
}

func Post(url string, data map[string]string) GistResponseHandler {
	jsonBody, _ := json.Marshal(data)

	resp := access(AccessApi{
		Method: "POST",
		URL:    URL,
		Body:   bytes.NewBuffer(jsonBody),
	})
	return resp
}
