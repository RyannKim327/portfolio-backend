package put

import (
	"portfolio-backend/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

var Projects = utils.Route{
	Path:       "/projects",
	Method:     utils.METHOD_PUT,
	Permission: utils.PERMISSION_ADMIN,
	Handler:    updateProject,
}

func updateProject(ctx *gin.Context) {
	var body gin.H
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Get the ID from body and convert to integer index
	idVal, ok := body["id"]
	if !ok {
		ctx.JSON(400, gin.H{
			"error": "missing project id",
		})
		return
	}

	var index int
	switch v := idVal.(type) {
	case string:
		i, err := strconv.Atoi(v)
		if err != nil {
			ctx.JSON(400, gin.H{
				"error": "invalid project id",
			})
			return
		}
		index = i
	case float64: // JSON numbers become float64 in Go
		index = int(v)
	default:
		ctx.JSON(400, gin.H{
			"error": "invalid project id type",
		})
		return
	}

	// Remove the ID field from body as it's only used for indexing
	delete(body, "id")

	// Read the existing gist (map with categories and projects)
	gist := utils.GistHandler("projects.json")
	if gist == nil {
		gist = make(gin.H)
	}

	// Extract the projects slice
	projectsInterface, ok := gist["projects"]
	if !ok {
		ctx.JSON(400, gin.H{
			"error": "projects not found in gist",
		})
		return
	}
	projects, ok := projectsInterface.([]interface{})
	if !ok {
		ctx.JSON(400, gin.H{
			"error": "projects is not a slice",
		})
		return
	}

	// Check if index is valid
	if index < 0 || index >= len(projects) {
		ctx.JSON(400, gin.H{
			"error": "project index out of range",
		})
		return
	}

	// Replace the project at the specified index with the body (without ID field)
	projects[index] = body

	// Update the gist with the modified projects slice
	gist["projects"] = projects

	// Write the updated gist back
	response := utils.GistPostHandler("projects.json", gist)

	ctx.JSON(200, gin.H{
		"from": response,
	})
}