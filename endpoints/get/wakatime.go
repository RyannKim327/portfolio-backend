package get

import (
	"encoding/json"
	"io"
	"net/http"
	"os"

	utils "portfolio-backend/utils"

	"github.com/gin-gonic/gin"
)

/*
 * TODO: This is just a template for the other endpoint
 */

type Repo struct {
	HtmlUrl     string `json:"html_url"`
	Description string `json:"description"`
	Name        string `json:"name"`
}

type WakaData struct {
	Name         string `json:"urlencoded_name"`
	Started      string `json:"created_at"`
	LastModified string `json:"last_heartbeat_at"`
	Respository  Repo   `json:"repository"`
}

type WakaResponse struct {
	Data []WakaData `json:"data"`
}

type WakaHours struct {
	Name    string  `json:"name"`
	Hours   float32 `json:"hours"`
	Minutes float32 `json:"minutes"`
}

type WakaDataProfile struct {
	Categories []WakaHours `json:"categories"`
	Daily      float32     `json:"daily_average_including_other_language"`
	Languages  []WakaHours `json:"languages"`
}

type WakaProfile struct {
	Data WakaDataProfile `json:"data"`
}

var Wakatime = utils.Route{
	Path:    "wakatime",
	Handler: wakatime,
}

func wakaprojects(ctx *gin.Context) WakaResponse {
	url_ := "https://wakatime.com/api/v1/users/ryannkim327/projects"
	req, err := http.NewRequest("GET", url_, nil)
	if err != nil {
		ctx.JSON(200, gin.H{
			"error": err,
		})
	}
	token := os.Getenv("WAKATIME")
	if token == "" {
		panic("WAKATIME not set")
	}

	// req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

	req.SetBasicAuth(os.Getenv("WAKATIME"), "")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		ctx.JSON(200, gin.H{
			"error": err,
		})
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		ctx.JSON(200, gin.H{
			"error": err,
		})
	}

	var result WakaResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		ctx.JSON(200, gin.H{
			"error": err,
		})
	}
	return result
}

func wakaprofile(ctx *gin.Context) WakaProfile {
	url_ := "https://wakatime.com/api/v1/users/ryannkim327/stats"
	req, err := http.NewRequest("GET", url_, nil)
	if err != nil {
		ctx.JSON(200, gin.H{
			"error": err,
		})
	}
	token := os.Getenv("WAKATIME")
	if token == "" {
		panic("WAKATIME not set")
	}

	// req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

	req.SetBasicAuth(os.Getenv("WAKATIME"), "")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		ctx.JSON(200, gin.H{
			"error": err,
		})
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		ctx.JSON(200, gin.H{
			"error": err,
		})
	}

	var result WakaProfile
	err = json.Unmarshal(body, &result)
	if err != nil {
		ctx.JSON(200, gin.H{
			"error": err,
		})
	}
	return result
}

func wakatime(ctx *gin.Context) {
	projs := wakaprojects(ctx)
	prof := wakaprofile(ctx)
	ctx.JSON(200, gin.H{
		"data": gin.H{
			"projects":         projs.Data,
			"coding_in_a_week": prof.Data,
		},
	})
}
