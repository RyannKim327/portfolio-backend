package get

import (
	"fmt"
	"strings"
	"sync"

	utils "portfolio-backend/utils"

	"github.com/gin-gonic/gin"
	"github.com/gocolly/colly/v2"
)

/*
 * TODO: This is just a template for the other endpoint
 */

const URL = "https://www.mangaread.org/%s"

type MangaSearchResponse struct {
	Title    string `json:"title`
	CoverUrl string `json:cover_url`
	Author   string `json:author`
	Link     string `json:link`
}

var Manga = utils.Route{
	Path:    "manga",
	Handler: manga,
}

func manga(ctx *gin.Context) {
	search := ctx.Query("s")
	read := ctx.Query("r")
	chapter := ctx.Query("c")

	if search != "" {
		// TODO: Search Manga
		ctx.JSON(200, manga_search(search))
		return
	} else if read != "" {
		// TODO: Read Manga
		ctx.JSON(200, manga_read(read, chapter))
		return
	}
	ctx.JSON(404, gin.H{
		"error": "Not found",
	})
}

func manga_search(search string) gin.H {
	s := fmt.Sprintf("?s=%s&post_type=wp-manga", search)
	c := colly.NewCollector()
	var response []MangaSearchResponse
	var mu sync.Mutex

	c.OnHTML("div.row.c-tabs-item__content", func(e *colly.HTMLElement) {
		r := MangaSearchResponse{
			Title:    e.ChildAttr("a", "title"),
			Author:   e.ChildText("div.post-content_item.mg_author .summary-content"),
			CoverUrl: e.ChildAttr("img.img-responsive", "src"),
			Link:     e.ChildAttr("a", "href"),
		}

		mu.Lock()
		response = append(response, r)
		mu.Unlock()
		// fmt.Println(element)
	})

	c.OnError(func(_ *colly.Response, err error) {
		fmt.Println("Scrape Error", err)
	})

	c.Visit(fmt.Sprintf(URL, s))
	c.Wait()

	return gin.H{
		"response": response,
	}
}

func manga_read(read string, chapter string) gin.H {
	c := colly.NewCollector()

	url := read
	if strings.HasSuffix(read, "/") {
		url = read[:len(read)-1]
	}
	split := strings.Split(url, "/")
	last := ""

	if len(split) > 0 {
		last = split[len(split)-1]
	}

	s := fmt.Sprintf("manga/%s/chapter-%s/", last, chapter)

	fmt.Println(s)
	var response []string
	var mu sync.Mutex

	c.OnHTML("div.reading-content", func(e *colly.HTMLElement) {
		// r := MangaSearchResponse{
		// 	Title:    e.ChildAttr("a", "title"),
		// 	Author:   e.ChildText("div.post-content_item.mg_author .summary-content"),
		// 	CoverUrl: e.ChildAttr("img.img-responsive", "src"),
		// 	Link:     e.ChildAttr("a", "href"),
		// }

		r := e.ChildAttr("img.wp-manga-chapter-img", "src")
		fmt.Println(r)
		mu.Lock()
		response = e.ChildAttrs("img.wp-manga-chapter-img", "src") // append(response, r)
		mu.Unlock()
		// fmt.Println(element)
	})

	c.OnError(func(_ *colly.Response, err error) {
		fmt.Println("Scrape Error", err)
	})

	c.Visit(fmt.Sprintf(URL, s))
	c.Wait()

	return gin.H{
		"response": response,
	}
}
