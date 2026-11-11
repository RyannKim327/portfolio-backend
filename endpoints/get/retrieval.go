package get

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"

	utils "portfolio-backend/utils"

	"github.com/gin-gonic/gin"
)

type getFileResult struct {
	Ok     bool `json:"ok"`
	Result struct {
		FilePath string `json:"file_path"`
	} `json:"result"`
}

var Retrieve = utils.Route{
	Path:    "retrieve",
	Method:  utils.METHOD_GET,
	Handler: fetchFile,
}

func fileDownload(apiKey, file string) string {
	return fmt.Sprintf("https://api.telegram.org/file/bot%s/%s", apiKey, file)
}

func resolveFilePath(apiKey, fileID string) (string, error) {
	fileURL := fmt.Sprintf("https://api.telegram.org/bot%s/getFile?file_id=%s", apiKey, url.QueryEscape(fileID))

	resp, err := http.Get(fileURL)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return "", fmt.Errorf("telegram responded with %s", resp.Status)
	}

	var data getFileResult
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return "", err
	}

	if !data.Ok || data.Result.FilePath == "" {
		return "", errors.New("invalid file id")
	}

	return data.Result.FilePath, nil
}

func fetchFile(ctx *gin.Context) {
	apiKey := os.Getenv("TG_API")

	var ids []string
	for _, value := range ctx.QueryArray("file") {
		for _, id := range strings.Split(value, ",") {
			if id = strings.TrimSpace(id); id != "" {
				ids = append(ids, id)
			}
		}
	}

	if len(ids) == 0 {
		ctx.JSON(400, gin.H{"error": "missing file id"})
		return
	}

	paths := make([]string, 0, len(ids))
	for _, id := range ids {
		path, err := resolveFilePath(apiKey, id)
		if err != nil {
			ctx.JSON(502, gin.H{"error": err.Error()})
			return
		}
		paths = append(paths, path)
	}

	first, err := http.Get(fileDownload(apiKey, paths[0]))
	if err != nil {
		ctx.JSON(502, gin.H{"error": err.Error()})
		return
	}
	defer first.Body.Close()

	if first.StatusCode >= 400 {
		ctx.JSON(502, gin.H{"error": first.Status})
		return
	}

	reader := bufio.NewReader(first.Body)
	peek, err := reader.Peek(512)
	if err != nil && err != io.EOF && err != bufio.ErrBufferFull {
		ctx.JSON(502, gin.H{"error": err.Error()})
		return
	}

	name := ctx.Query("name")
	if name == "" {
		name = "download"
	}

	ctx.Header("Content-Type", http.DetectContentType(peek))
	ctx.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%q", name))
	ctx.Status(200)

	if _, err := io.Copy(ctx.Writer, reader); err != nil {
		return
	}

	for _, path := range paths[1:] {
		resp, err := http.Get(fileDownload(apiKey, path))
		if err != nil {
			return
		}
		if resp.StatusCode >= 400 {
			resp.Body.Close()
			return
		}
		_, err = io.Copy(ctx.Writer, resp.Body)
		resp.Body.Close()
		if err != nil {
			return
		}
	}
}
