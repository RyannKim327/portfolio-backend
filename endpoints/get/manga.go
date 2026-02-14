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

const MANGA_URL = "https://www.mangaread.org/%s"

type MangaSearchResponse struct {
	Title                string `json:"title"`
	CoverUrl             string `json:"cover_url"`
	Author               string `json:"author"`
	Link                 string `json:"link"`
	CurrentLatestChapter string `json:"current_last_chapter"`
	Status               string `json:"status"`
}

type MangaChapters struct {
	ChapterName string `json:"chapter_name"`
	ChapterLink string `json:"chapter_link"`
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
	} else if read != "" && chapter == "" {
		ctx.JSON(200, check_chapters(read))
		return
	} else if read != "" && chapter != "" {
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
		chap := strings.Split(e.ChildText("div.meta-item.latest-chap span.font-meta.chapter"), " ")
		link := strings.Split(e.ChildAttr("a", "href"), "/")

		r := MangaSearchResponse{
			Title:                e.ChildAttr("a", "title"),
			Author:               e.ChildText("div.post-content_item.mg_author .summary-content"),
			CoverUrl:             e.ChildAttr("img.img-responsive", "src"),
			Link:                 link[len(link)-2],
			CurrentLatestChapter: chap[1],
			Status:               e.ChildText("div.post-content_item.mg_status div.summary-content"),
		}

		mu.Lock()
		response = append(response, r)
		mu.Unlock()
	})

	c.OnError(func(_ *colly.Response, err error) {
		fmt.Println("Scrape Error", err)
	})

	c.Visit(fmt.Sprintf(MANGA_URL, s))
	c.Wait()

	if response != nil {
		return gin.H{
			"response": response,
		}
	}
	return gin.H{
		"error": fmt.Sprintf("Nothing found for %s", search),
	}
}

func check_chapters(manga string) gin.H {
	c := colly.NewCollector()
	s := fmt.Sprintf("manga/%s/", manga)

	var mu sync.Mutex
	var response []MangaChapters

	c.OnHTML("li.wp-manga-chapter", func(e *colly.HTMLElement) {
		link := strings.Split(e.ChildAttr("a", "href"), "/")
		r := MangaChapters{
			ChapterName: e.ChildText("li.wp-manga-chapter a"),
			ChapterLink: link[len(link)-2],
		}

		mu.Lock()
		response = append(response, r)
		mu.Unlock()
	})

	c.Visit(fmt.Sprintf(MANGA_URL, s))
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

	s := fmt.Sprintf("manga/%s/%s/", last, chapter)

	var response []string
	var mu sync.Mutex

	c.OnHTML("div.reading-content", func(e *colly.HTMLElement) {
		r := e.ChildAttr("img.wp-manga-chapter-img", "src")
		fmt.Println(r)
		mu.Lock()
		response = e.ChildAttrs("img.wp-manga-chapter-img", "src")
		mu.Unlock()
	})

	c.OnError(func(_ *colly.Response, err error) {
		fmt.Println("Scrape Error", err)
	})

	c.Visit(fmt.Sprintf(MANGA_URL, s))
	c.Wait()

	return gin.H{
		"response": response,
	}
}
