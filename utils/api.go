package utils

import (
	"encoding/json"
	"net/http"
)

func TelegramAccess(params AccessAPI) TelegramAPI {
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
		return TelegramAPI{Error: err}
	}

	// TODO: Request Executor
	client := &http.Client{}
	resp, err := client.Do(req)
	// TODO: To prevent errors
	if err != nil {
		return TelegramAPI{
			Error: err,
		}
	}

	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return TelegramAPI{Error: err}
	}

	// TODO: To send the response data
	var data TelegramResponse
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return TelegramAPI{Error: err}
	}

	return TelegramAPI{Response: data}
}
