package get

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strings"

	"portfolio-backend/utils"

	"github.com/gin-gonic/gin"
	"github.com/gocolly/colly/v2"
)

const URL = "https://v2.ytmp3.wtf/%s"

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

	tokenizer := ytToken(videoID)

	fmt.Println(tokenizer)
	body := strings.NewReader(fmt.Sprintf(`
		{
			"url": "https://youtube.com/watch?v=%s",
			"convert": "gogogo",
			"token": "t_58bdf31e8ed7a1fcf2dedcfb7d7c881a424e1051130b21cb2ddf6b44a8caec04",
			"token_validto": "99999999999"
		}
	`, &videoID))

	// TODO: To initiate request
	req, err := http.NewRequest("POST", fmt.Sprintf(URL, "convert"), body)
	if err != nil {
		ctx.JSON(400, gin.H{
			"error": err,
		})
		return
	}

	// TODO: Request Executor
	client := &http.Client{}
	resp, err := client.Do(req)
	// TODO: To prevent errors
	if err != nil {
		ctx.JSON(400, gin.H{
			"error": err,
		})
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		ctx.JSON(400, gin.H{
			"error": "Something went wrong",
			"code":  resp.StatusCode,
		})
		return
	}

	// TODO: To send the response data
	var data gin.H
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		ctx.JSON(500, gin.H{
			"error": err,
		})
		return
	}
}

