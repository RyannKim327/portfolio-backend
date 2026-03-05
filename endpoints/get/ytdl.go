package get

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"regexp"

	"portfolio-backend/utils"

	"github.com/gin-gonic/gin"
)

const URL = "https://v2.ytmp3.wtf/%s"

type Response struct {
	// adjust fields based on actual API response
	Link string `json:"link"`
}

var YoutubeDL = utils.Route{
	Path:    "yt",
	Method:  utils.METHOD_GET,
	Handler: youtube,
}

func matcher(ytUrl string) string {
	// TODO: To filter the link and to extract only the video ID
	shareVideoFormat := regexp.MustCompile(`youtu\.be/([A-Za-z0-9_-]{11})`)
	linkVideoFormat := regexp.MustCompile(`youtube\.com/watch\?v=([A-Za-z0-9_-]{11})`)

	matchShare := shareVideoFormat.FindStringSubmatch(ytUrl)
	matchLink := linkVideoFormat.FindStringSubmatch(ytUrl)

	if len(matchShare) >= 1 {
		return matchShare[1]
	} else if len(matchLink) >= 1 {
		return matchLink[1]
	} else {
		return ytUrl
	}
}

func youtube(ctx *gin.Context) {
	// TODO: To get the parameters
	videoId := ctx.Query("videoID")

	if videoId == "" {
		ctx.JSON(400, gin.H{
			"error": "videoID is required as parameter",
		})
		return
	}

	videoID := matcher(videoId)

	url_ := fmt.Sprintf(
		"https://youtube-mp36.p.rapidapi.com/dl?id=%s",
		videoID,
	)

	req, err := http.NewRequest("GET", url_, nil)
	if err != nil {
		ctx.JSON(200, gin.H{
			"error": err,
		})
		return
	}

	req.Header.Add("x-rapidapi-key", os.Getenv("RAPIDKEY"))
	req.Header.Add("x-rapidapi-host", os.Getenv("RAPIDHOST"))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		ctx.JSON(200, gin.H{
			"error": err,
		})
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		ctx.JSON(200, gin.H{
			"error": err,
		})
		return
	}

	var result Response
	err = json.Unmarshal(body, &result)
	if err != nil {
		ctx.JSON(200, gin.H{
			"error": err,
		})
		return
	}

	re := regexp.MustCompile(`n=([^&]+)`)
	matches := re.FindStringSubmatch(result.Link)

	if len(matches) < 2 {
		ctx.JSON(200, gin.H{
			"error": "No title found",
		})
		return
	}

	title, err := url.QueryUnescape(matches[1])
	if err != nil {
		ctx.JSON(200, gin.H{
			"error": err,
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "Done",
		"title":   title,
		"url":     result.Link,
	})
}
