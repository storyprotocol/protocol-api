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

// GetPolicyFrameworkManager Example godoc
// @Summary Get a PolicyFrameworkManager
// @Schemes
// @Description Retrieve a PolicyFrameworkManager
// @Security ApiKeyAuth
// @Host https://edge.stg.storyprotocol.net
// @param X-API-Key header string true "API Key"
// @Tags PolicyFrameworkManagers
// @Accept json
// @Produce json
// @Param        pfwmId   path      string  true  "PolicyFrameworkManager ID"
// @Success 200 {object} PolicyFrameworkManagerResponse
// @Router /api/v1/policyframeworks/{pfwmId} [get]
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
// @Tags PolicyFrameworkManagers
// @Accept json
// @Produce json
// @Success 200 {object} PolicyFrameworkManagersResponse
// @Router /api/v1/policyframeworks [post]
func NewListPolicyFrameworkManagers(graphService thegraph.TheGraphServiceBeta, httpClient xhttp.Client) func(c *gin.Context) {
	return func(c *gin.Context) {
		var requestBody beta_v0.PolicyFrameworkManagerRequestBody
		if err := c.BindJSON(&requestBody); err != nil {
			logger.Errorf("Failed to read request body: %v", err)
			requestBody = beta_v0.PolicyFrameworkManagerRequestBody{}
		}

		pfwms, err := graphService.ListPolicyFrameworkManagers(fromPolicyFWMRequestQueryOptions(requestBody.Options))
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

func fromPolicyFWMRequestQueryOptions(options *beta_v0.PFWMQueryOptions) *thegraph.TheGraphQueryOptions {
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

	queryOptions.Where.Address = options.Where.Address
	queryOptions.Where.Name = options.Where.Name

	return queryOptions
}
