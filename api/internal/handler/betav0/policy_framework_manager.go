package betav0

import (
	"github.com/gin-gonic/gin"
	"github.com/storyprotocol/protocol-api/api/internal/handler/betav0/messages"
	beta_v0 "github.com/storyprotocol/protocol-api/api/internal/models/betav0"
	"github.com/storyprotocol/protocol-api/api/internal/service/thegraph"
	xhttp "github.com/storyprotocol/protocol-api/pkg/http"
	"github.com/storyprotocol/protocol-api/pkg/logger"
	"net/http"
	"strings"
)

// @BasePath /

// GetPolicyFrameworkManager Example godoc
// @Summary Get a PolicyFrameworkManager
// @Schemes
// @Description Retrieve a PolicyFrameworkManager
// @Security ApiKeyAuth
// @Host https://edge.stg.storyprotocol.net
// @param X-API-Key header string true "API Key"
// @Tags Policies
// @Accept json
// @Produce json
// @Param        pfwmId   path      string  true  "PolicyFrameworkManager ID"
// @Success 200 {object} PolicyFrameworkManagerResponse
// @Router /api/v1/policies/frameworks/{pfwmId} [get]
func NewGetPolicyFrameworkManager(graphService thegraph.TheGraphServiceBeta, httpClient xhttp.Client) func(c *gin.Context) {
	return func(c *gin.Context) {
		pfwmId := c.Param("pfwmId")

		pfwms, err := graphService.GetPolicyFrameworkManager(pfwmId)
		if err != nil {
			logger.Errorf("Failed to get policy framework manager: %v", err)
			c.JSON(http.StatusInternalServerError, messages.ErrorMessage("Internal server error"))
			return
		}

		c.JSON(http.StatusOK, beta_v0.PolicyFrameworkManagerResponse{
			Data: pfwms,
		})
	}
}

// @BasePath /

// ListPolicyFrameworkManagers Example godoc
// @Summary List PolicyFrameworkManagers
// @Schemes
// @Description Retrieve a paginated, filtered list of PolicyFrameworkManagers
// @Host https://edge.stg.storyprotocol.net
// @Security ApiKeyAuth
// @param X-API-Key header string true "API Key"
// @Param data body betav0.PolicyFrameworkManagerRequestBody true "Query Parameters ("where" values are optional. Remove if not using)"
// @Tags Policies
// @Accept json
// @Produce json
// @Success 200 {object} PolicyFrameworkManagersResponse
// @Router /api/v1/policies/frameworks [post]
func NewListPolicyFrameworkManagers(graphService thegraph.TheGraphServiceBeta, httpClient xhttp.Client) func(c *gin.Context) {
	return func(c *gin.Context) {
		var requestBody *beta_v0.PolicyFrameworkManagerRequestBody
		if err := c.ShouldBindJSON(&requestBody); err != nil {
			logger.Errorf("Failed to read request body: %v", err)
			requestBody = &beta_v0.PolicyFrameworkManagerRequestBody{}
		}

		pfwms, err := graphService.ListPolicyFrameworkManagers(fromPolicyFWMRequestQueryOptions(requestBody))
		if err != nil {
			logger.Errorf("Failed to list policy framework managers : %v", err)
			c.JSON(http.StatusInternalServerError, messages.ErrorMessage("Internal server error"))
			return
		}

		c.JSON(http.StatusOK, beta_v0.PolicyFrameworkManagersResponse{
			Data: pfwms,
		})
	}
}

func fromPolicyFWMRequestQueryOptions(body *beta_v0.PolicyFrameworkManagerRequestBody) *thegraph.TheGraphQueryOptions {
	if body == nil {
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
	queryOptions.OrderDirection = strings.ToLower(body.Options.OrderDirection)
	queryOptions.OrderBy = body.Options.OrderBy
	queryOptions.Where.Address = body.Options.Where.Address
	queryOptions.Where.Name = body.Options.Where.Name

	return queryOptions
}
