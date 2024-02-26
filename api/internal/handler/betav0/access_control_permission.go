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

func NewListAccessControlPermissions(graphService thegraph.TheGraphServiceBeta, httpClient xhttp.Client) func(c *gin.Context) {
	return func(c *gin.Context) {
		var requestBody *beta_v0.AccessControlPermissionsRequestBody
		if err := c.BindJSON(&requestBody); err != nil {
			logger.Errorf("Failed to read request body: %v", err)
			requestBody = &beta_v0.AccessControlPermissionsRequestBody{}
		}

		acps, err := graphService.ListAccessControlPermissions(fromACPRequestQueryOptions(requestBody))
		if err != nil {
			logger.Errorf("Failed to get access control permissions: %v", err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, messages.ErrorMessage("Internal server error"))
			return
		}

		c.JSON(http.StatusOK, beta_v0.AccessControlPermissionResponse{
			Data: acps,
		})
	}
}

func fromACPRequestQueryOptions(requestBody *beta_v0.AccessControlPermissionsRequestBody) *thegraph.TheGraphQueryOptions {
	if requestBody == nil {
		return &thegraph.TheGraphQueryOptions{
			First: 100,
			Skip:  0,
		}
	}

	var queryOptions = &thegraph.TheGraphQueryOptions{}

	if requestBody.Options.Pagination.Limit == 0 {
		requestBody.Options.Pagination.Limit = 100
	}

	queryOptions.First = requestBody.Options.Pagination.Limit
	queryOptions.Skip = requestBody.Options.Pagination.Offset
	queryOptions.OrderDirection = requestBody.Options.OrderDirection
	queryOptions.OrderBy = requestBody.Options.OrderBy

	queryOptions.Where.Module = requestBody.Options.Where.Module
	queryOptions.Where.Name = requestBody.Options.Where.Name

	return queryOptions
}
