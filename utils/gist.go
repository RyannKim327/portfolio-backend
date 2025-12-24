package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

var _ = godotenv.Load()

var URL = fmt.Sprintf("https://api.github.com/gists/%s", os.Getenv("GIST_ID"))

func access(params AccessAPI) GistResponseHandler {
	/*
	 * TODO: To create a door to access gist to the backend and
	 * to distribute the data through the other platform connected
	 * to my domain
	 */

	// TODO: To automatically use GET request if params.Method not exists
	if params.Method == "" {
		params.Method = "GET"
	}

	// TODO: To initiate request
	req, err := http.NewRequest(params.Method, params.URL, params.Body)
	if err != nil {
		return GistResponseHandler{Error: err}
	}

	// TODO: To setup the headers for github gist access
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("API_KEY")))
	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("X-Github-Api-Version", "2022-11-28")

	// TODO: Request Executor
	client := &http.Client{}
	resp, err := client.Do(req)
	// TODO: To prevent errors
	if err != nil {
		return GistResponseHandler{
			Error: err,
		}
	}

	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return GistResponseHandler{Error: err}
	}

	// TODO: To send the response data
	var data Gist
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return GistResponseHandler{Error: err}
	}

	return GistResponseHandler{Response: data}
}

func Get() GistResponseHandler {
	resp := access(AccessAPI{
		URL: URL,
	})
	return resp
}

func Post(data interface{}) GistResponseHandler {
	jsonBody, err := json.Marshal(data)
	if err != nil {
		return GistResponseHandler{
			Error: err,
		}
	}

	resp := access(AccessAPI{
		Method: "POST",
		URL:    URL,
		Body:   bytes.NewBuffer(jsonBody),
	})

	return resp
}
