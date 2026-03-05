package get

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"

	"portfolio-backend/utils"

	"github.com/gin-gonic/gin"
	"github.com/gocolly/colly/v2"
)

const URL = "https://v2.ytmp3.wtf/%s"

type Response struct {
	// adjust fields based on actual API response
	Link string `json:"link"`
}

var Youtube = utils.Route{
	Path:    "yt",
	Method:  utils.METHOD_GET,
	Handler: youtube,
}

func matcher(ytUrl string) string {
	shareVideoFormat := regexp.MustCompile(`youtu\.be/[A-Za-z0-9_-]{11}`)
	linkVideoFormat := regexp.MustCompile(`youtube\.com/watch\?v=[A-Za-z0-9_-]{11}`)

	matchShare := shareVideoFormat.FindStringSubmatch(ytUrl)
	matchLink := linkVideoFormat.FindStringSubmatch(ytUrl)
	if len(matchShare) > 1 {
		return matchShare[1]
	} else if len(matchLink) > 1 {
		return matchLink[1]
	} else {
		return ytUrl
	}
}

func ytToken(url string) string {
	c := colly.NewCollector()
	xURL := fmt.Sprintf("button/?url=https%3A%2F%2Fwww.youtube.com%2Fwatch%3Fv%3D%s", url)

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64; rv:144.0) Gecko/20100101 Firefox/144.0")
	})

	c.OnHTML("script", func(e *colly.HTMLElement) {
		fmt.Println(e.Text)
	})

	c.OnError(func(_ *colly.Response, err error) {
		fmt.Println("Scrape Error", err)
	})

	c.Visit(fmt.Sprintf(URL, xURL))
	c.Wait()

	return ""
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

	fmt.Printf(videoID)
	url := fmt.Sprintf(
		"https://youtube-mp36.p.rapidapi.com/dl?id=%s",
		videoID,
	)

	req, err := http.NewRequest("GET", url, nil)
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

	// ctx.Header("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s.mp3\"", "audio"))
	// ctx.Header("Content-Type", "audio/mpeg")

	// io.Copy(ctx.Writer, resp.Body)
	re := regexp.MustCompile(`n=([^&]+)`)
	matches := re.FindStringSubmatch(result.Link)

	if len(matches) < 2 {
		fmt.Println("No title found in link")
		return
	}

	title := matches[1]

	ctx.JSON(200, gin.H{
		"message": "Done",
		"title":   title,
		"url":     result.Link,
	})
}
