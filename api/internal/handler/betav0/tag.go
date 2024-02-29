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

func NewListTags(graphService thegraph.TheGraphServiceBeta, httpClient xhttp.Client) func(c *gin.Context) {
	return func(c *gin.Context) {
		var requestBody *beta_v0.TagRequestBody
		if err := c.ShouldBindJSON(&requestBody); err != nil {
			logger.Errorf("Failed to read request body: %v", err)
			requestBody = &beta_v0.TagRequestBody{}
		}

		tags, err := graphService.ListTag(fromTagRequestQueryOptions(requestBody))
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

func fromTagRequestQueryOptions(body *beta_v0.TagRequestBody) *thegraph.TheGraphQueryOptions {
	if body == nil {
		return &thegraph.TheGraphQueryOptions{
			First: 100,
			Skip:  0,
		}
	}
	if body.Options == nil {
		return &thegraph.TheGraphQueryOptions{
			First: 100,
			Skip:  0,
		}
	}

	var queryOptions = &thegraph.TheGraphQueryOptions{}

	if body.Options.Pagination.Limit == 0 {
		body.Options.Pagination.Limit = 100
	}

	queryOptions.First = body.Options.Pagination.Limit
	queryOptions.Skip = body.Options.Pagination.Offset
	queryOptions.OrderDirection = body.Options.OrderDirection
	queryOptions.OrderBy = body.Options.OrderBy
	queryOptions.Where.IPID = body.Options.Where.IPID
	queryOptions.Where.Tag = body.Options.Where.Tag

	return queryOptions
}
