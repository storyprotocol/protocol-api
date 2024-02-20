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

// GetRoyaltyPolicy Example godoc
// @Summary Get a RoyaltyPolicy
// @Schemes
// @Description Retrieve a RoyaltyPolicy
// @Security ApiKeyAuth
// @Host https://edge.stg.storyprotocol.net
// @param X-API-Key header string true "API Key"
// @Tags RoyaltyPolicies
// @Accept json
// @Produce json
// @Param        royaltyPolicyId   path      string  true  "Royalty Policy ID"
// @Success 200 {object} RoyaltyPolicyResponse
// @Router /api/v1/royaltypolicies/{royaltyPolicyId} [get]
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
// @Tags RoyaltyPolicies
// @Accept json
// @Produce json
// @Success 200 {object} RoyaltyPoliciesResponse
// @Router /api/v1/royaltypolicies [post]
func NewListRoyaltyPolicies(graphService thegraph.TheGraphServiceBeta, httpClient xhttp.Client) func(c *gin.Context) {
	return func(c *gin.Context) {
		var requestBody beta_v0.RoyaltyPolicyRequestBody
		if err := c.BindJSON(&requestBody); err != nil {
			logger.Errorf("Failed to read request body: %v", err)
			requestBody = beta_v0.RoyaltyPolicyRequestBody{}
		}

		roys, err := graphService.ListRoyaltyPolicies(fromRoyaltyPolicyRequestQueryOptions(requestBody.Options))
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

func fromRoyaltyPolicyRequestQueryOptions(options *beta_v0.RoyaltyPolicyQueryOptions) *thegraph.TheGraphQueryOptions {
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

	return queryOptions
}
