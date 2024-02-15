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

// GetPolicy Example godoc
// @Summary Get a Policy
// @Schemes
// @Description Retrieve a Policy
// @Security ApiKeyAuth
// @Host https://edge.stg.storyprotocol.net
// @param X-API-Key header string true "API Key"
// @Tags Policies
// @Accept json
// @Produce json
// @Param        policyId   path      string  true  "Policy ID"
// @Success 200 {object} PolicyResponse
// @Router /api/v1/policies/{policyId} [get]
func NewGetPolicy(graphService thegraph.TheGraphServiceBeta, httpClient xhttp.Client) func(c *gin.Context) {
	return func(c *gin.Context) {
		policyId := c.Param("policyId")

		policies, err := graphService.GetPolicy(policyId)
		if err != nil {
			logger.Errorf("Failed to get policy: %v", err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, messages.ErrorMessage("Internal server error"))
			return
		}

		c.JSON(http.StatusOK, beta_v0.PolicyResponse{
			Data: policies,
		})
	}
}

// @BasePath /

// ListPolicies Example godoc
// @Summary List Policies
// @Schemes
// @Description Retrieve a paginated, filtered list of Policies
// @Security ApiKeyAuth
// @Host https://edge.stg.storyprotocol.net
// @param X-API-Key header string true "API Key"
// @Param data body betav0.PolicyRequestBody true "Query Parameters ("where" values are optional. Remove if not using)"
// @Tags Policies
// @Accept json
// @Produce json
// @Success 200 {object} PoliciesResponse
// @Router /api/v1/policies [post]
func NewListPolicies(graphService thegraph.TheGraphServiceBeta, httpClient xhttp.Client) func(c *gin.Context) {
	return func(c *gin.Context) {
		var requestBody beta_v0.PolicyRequestBody
		if err := c.BindJSON(&requestBody); err != nil {
			logger.Errorf("Failed to read request body: %v", err)
			requestBody = beta_v0.PolicyRequestBody{}
		}

		pols, err := graphService.ListPolicies(fromPolicuyRequestQueryOptions(requestBody.Options))
		if err != nil {
			logger.Errorf("Failed to list policies: %v", err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, messages.ErrorMessage("Internal server error"))
			return
		}

		c.JSON(http.StatusOK, beta_v0.PoliciesResponse{
			Data: pols,
		})
	}
}

func fromPolicuyRequestQueryOptions(options *beta_v0.PolicyQueryOptions) *thegraph.TheGraphQueryOptions {
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

	queryOptions.Where.PolicyFrameworkManager = options.Where.PolicyFrameworkManager

	return queryOptions
}
