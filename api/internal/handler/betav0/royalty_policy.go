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

// GetRoyaltyPolicy Example godoc
// @Summary Get a RoyaltyPolicy
// @Schemes
// @Description Retrieve a RoyaltyPolicy
// @Security ApiKeyAuth
// @Host https://edge.stg.storyprotocol.net
// @param X-API-Key header string true "API Key"
// @Tags Royalties
// @Accept json
// @Produce json
// @Param        royaltyPolicyId   path      string  true  "Royalty Policy ID"
// @Success 200 {object} RoyaltyPolicyResponse
// @Router /api/v1/royalties/policies/{royaltyPolicyId} [get]
func NewGetRoyaltyPolicy(graphService thegraph.TheGraphServiceBeta, httpClient xhttp.Client) func(c *gin.Context) {
	return func(c *gin.Context) {
		royaltyPolicyId := c.Param("royaltyPolicyId")

		roys, err := graphService.GetRoyaltyPolicy(royaltyPolicyId)
		if err != nil {
			logger.Errorf("Failed to get royalty policy: %v", err)
			c.JSON(http.StatusInternalServerError, messages.ErrorMessage("Internal server error"))
			return
		}

		c.JSON(http.StatusOK, beta_v0.RoyaltyPolicyResponse{
			Data: roys,
		})
	}
}

// @BasePath /

// ListRoyaltyPolicies Example godoc
// @Summary List RoyaltyPolicies
// @Schemes
// @Description Retrieve a paginated, filtered list of RoyaltyPolicies
// @Host https://edge.stg.storyprotocol.net
// @Security ApiKeyAuth
// @param X-API-Key header string true "API Key"
// @Param data body betav0.RoyaltyRequestBody true "Query Parameters ("where" values are optional. Remove if not using)"
// @Tags Royalties
// @Accept json
// @Produce json
// @Success 200 {object} RoyaltyPoliciesResponse
// @Router /api/v1/royalties/policies [post]
func NewListRoyaltyPolicies(graphService thegraph.TheGraphServiceBeta, httpClient xhttp.Client) func(c *gin.Context) {
	return func(c *gin.Context) {
		var requestBody *beta_v0.RoyaltyPolicyRequestBody
		if err := c.ShouldBindJSON(&requestBody); err != nil {
			logger.Errorf("Failed to read request body: %v", err)
			requestBody = &beta_v0.RoyaltyPolicyRequestBody{}
		}

		roys, err := graphService.ListRoyaltyPolicies(fromRoyaltyPolicyRequestQueryOptions(requestBody))
		if err != nil {
			logger.Errorf("Failed to list royalties: %v", err)
			c.JSON(http.StatusInternalServerError, messages.ErrorMessage("Internal server error"))
			return
		}

		c.JSON(http.StatusOK, beta_v0.RoyaltyPoliciesResponse{
			Data: roys,
		})
	}
}

func fromRoyaltyPolicyRequestQueryOptions(body *beta_v0.RoyaltyPolicyRequestBody) *thegraph.TheGraphQueryOptions {
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
	return queryOptions
}
