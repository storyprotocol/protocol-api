package betav0

import (
	"github.com/gin-gonic/gin"
	"github.com/storyprotocol/protocol-api/api/internal/handler/betav0/messages"
	beta_v0 "github.com/storyprotocol/protocol-api/api/internal/models/betav0"
	"github.com/storyprotocol/protocol-api/api/internal/service/thegraph"
	xhttp "github.com/storyprotocol/protocol-api/pkg/http"
	"github.com/storyprotocol/protocol-api/pkg/logger"
	"net/http"
)

// @BasePath /

// GetTag Example godoc
// @Summary Get a Tag
// @Schemes
// @Description Retrieve a Tag
// @Security ApiKeyAuth
// @param X-API-Key header string true "API Key"
// @Tags Tags
// @Accept json
// @Produce json
// @Param        tagId   path      string  true  "Tag ID"
// @Success 200 {object} TagResponse
// @Router /api/v1/tags/{tagId} [get]
func NewGetTag(graphService thegraph.TheGraphServiceBeta, httpClient xhttp.Client) func(c *gin.Context) {
	return func(c *gin.Context) {
		ipId := c.Param("tagId")

		tags, err := graphService.GetTag(ipId)
		if err != nil {
			logger.Errorf("Failed to get tag: %v", err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, messages.ErrorMessage("Internal server error"))
			return
		}

		c.JSON(http.StatusOK, beta_v0.TagResponse{
			Data: tags,
		})
	}
}

// @BasePath /

// ListTags Example godoc
// @Summary List Tags
// @Schemes
// @Description Retrieve a paginated, filtered list of Tags
// @Security ApiKeyAuth
// @param X-API-Key header string true "API Key"
// @Param data body betav0.TagRequestBody true "Query Parameters ("where" values are optional. Remove if not using)"
// @Tags Tags
// @Accept json
// @Produce json
// @Success 200 {object} TagsResponse
// @Router /api/v1/tags [post]
func NewListTags(graphService thegraph.TheGraphServiceBeta, httpClient xhttp.Client) func(c *gin.Context) {
	return func(c *gin.Context) {
		var requestBody beta_v0.TagRequestBody
		if err := c.BindJSON(&requestBody); err != nil {
			logger.Errorf("Failed to read request body: %v", err)
			requestBody = beta_v0.TagRequestBody{}
		}

		tags, err := graphService.ListTag(fromTagRequestQueryOptions(requestBody.Options))
		if err != nil {
			logger.Errorf("Failed to list tags: %v", err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, messages.ErrorMessage("Internal server error"))
			return
		}

		c.JSON(http.StatusOK, beta_v0.TagsResponse{
			Data: tags,
		})
	}
}

func fromTagRequestQueryOptions(options *beta_v0.TagQueryOptions) *thegraph.TheGraphQueryOptions {
	if options == nil {
		return &thegraph.TheGraphQueryOptions{
			First: 100,
			Skip:  0,
		}
	}

	var queryOptions = &thegraph.TheGraphQueryOptions{}

	if options.Pagination.Limit == 0 {
		options.Pagination.Limit = 100
	}

	queryOptions.First = options.Pagination.Limit
	queryOptions.Skip = options.Pagination.Offset

	queryOptions.Where.IPID = options.Where.IPID
	queryOptions.Where.Tag = options.Where.Tag

	return queryOptions
}
