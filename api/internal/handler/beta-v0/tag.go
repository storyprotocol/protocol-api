package beta_v0

import (
	"github.com/gin-gonic/gin"
	"github.com/storyprotocol/protocol-api/api/internal/handler/beta-v0/messages"
	beta_v0 "github.com/storyprotocol/protocol-api/api/internal/models/beta-v0"
	options2 "github.com/storyprotocol/protocol-api/api/internal/models/beta-v0/options"
	"github.com/storyprotocol/protocol-api/api/internal/service/thegraph"
	xhttp "github.com/storyprotocol/protocol-api/pkg/http"
	"github.com/storyprotocol/protocol-api/pkg/logger"
	"net/http"
)

// @BasePath /api/v1

// GetTag Example godoc
// @Summary Get a Tag
// @Schemes
// @Description Retrieve a Tag
// @Tags Tags
// @Accept json
// @Produce json
// @Param        tagId   path      string  true  "Tag ID"
// @Success 200 {object} TagResponse
// @Router /tags/{tagId} [get]
func NewGetTag(graphService thegraph.TheGraphServiceBeta, httpClient xhttp.Client) func(c *gin.Context) {
	return func(c *gin.Context) {
		ipId := c.Param("tagId")

		tags, err := graphService.GetTag(ipId)
		if err != nil {
			logger.Errorf("Failed to get tag: %v", err)
			c.JSON(http.StatusInternalServerError, messages.ErrorMessage("Internal server error"))
			return
		}

		c.JSON(http.StatusOK, beta_v0.TagResponse{
			Data: tags,
		})
	}
}

// @BasePath /api/v1

// ListTags Example godoc
// @Summary List Tags
// @Schemes
// @Description Retrieve a paginated, filtered list of Tags
// @Tags Tags
// @Accept json
// @Produce json
// @Success 200 {object} TagsResponse
// @Router /tags [post]
func NewListTags(graphService thegraph.TheGraphServiceBeta, httpClient xhttp.Client) func(c *gin.Context) {
	return func(c *gin.Context) {
		var requestBody options2.RequestBody
		if err := c.BindJSON(&requestBody); err != nil {
			logger.Errorf("Failed to read request body: %v", err)
			requestBody = options2.RequestBody{}
		}

		tags, err := graphService.ListTag(thegraph.FromRequestQueryOptions(requestBody.Options))
		if err != nil {
			logger.Errorf("Failed to list tags: %v", err)
			c.JSON(http.StatusInternalServerError, messages.ErrorMessage("Internal server error"))
			return
		}

		c.JSON(http.StatusOK, beta_v0.TagsResponse{
			Data: tags,
		})
	}
}
