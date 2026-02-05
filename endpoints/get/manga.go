package get

import (
	"fmt"
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

	if search != "" {
		// TODO: Search Manga
		ctx.JSON(200, manga_search(search))
	} else if read != "" {
		// TODO: Read Manga
		manga_read(read)
	}
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

func manga_read(read string) {
}
