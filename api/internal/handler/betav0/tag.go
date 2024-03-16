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

		tags, renTags, err := graphService.GetTag(ipId)
		if err != nil {
			logger.Errorf("Failed to get tag: %v", err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, messages.ErrorMessage("Internal server error"))
			return
		}

		if tags != nil {
			c.JSON(http.StatusOK, beta_v0.TagResponse{
				Data: tags,
			})
		} else {
			c.JSON(http.StatusOK, beta_v0.RenTagResponse{
				Data: renTags,
			})
		}
	}
}

func NewListTags(graphService thegraph.TheGraphServiceBeta, httpClient xhttp.Client) func(c *gin.Context) {
	return func(c *gin.Context) {
		var requestBody *beta_v0.TagRequestBody
		if err := c.ShouldBindJSON(&requestBody); err != nil {
			logger.Errorf("Failed to read request body: %v", err)
			requestBody = &beta_v0.TagRequestBody{}
		}

		tags, renTags, err := graphService.ListTag(fromTagRequestQueryOptions(requestBody))
		if err != nil {
			logger.Errorf("Failed to list tags: %v", err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, messages.ErrorMessage("Internal server error"))
			return
		}

		if tags != nil {
			c.JSON(http.StatusOK, beta_v0.TagsResponse{
				Data: tags,
			})
		} else {
			c.JSON(http.StatusOK, beta_v0.RenTagsResponse{
				Data: renTags,
			})
		}
	}
}

func fromTagRequestQueryOptions(body *beta_v0.TagRequestBody) *thegraph.TheGraphQueryOptions {
	if body == nil {
		return &thegraph.TheGraphQueryOptions{
			First: 100,
			Limit: 100,
			Skip:  0,
		}
	}
	if body.Options == nil {
		return &thegraph.TheGraphQueryOptions{
			First: 100,
			Limit: 100,
			Skip:  0,
		}
	}

	var queryOptions = &thegraph.TheGraphQueryOptions{}

	if body.Options.Pagination.Limit == 0 {
		body.Options.Pagination.Limit = 100
	}

	queryOptions.First = body.Options.Pagination.Limit
	queryOptions.Limit = body.Options.Pagination.Limit
	queryOptions.Skip = body.Options.Pagination.Offset
	queryOptions.OrderDirection = body.Options.OrderDirection
	queryOptions.OrderBy = body.Options.OrderBy
	queryOptions.Where.IPID = body.Options.Where.IPID
	queryOptions.Where.Tag = body.Options.Where.Tag

	return queryOptions
}
