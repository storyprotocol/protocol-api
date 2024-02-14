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

// GetIPAPolicy Example godoc
// @Summary Get a IPAPolicy
// @Schemes
// @Description Retrieve a IPAPolicy
// @Security ApiKeyAuth
// @param X-API-Key header string true "API Key"
// @Tags IPAPolicies
// @Accept json
// @Produce json
// @Param        ipaPolicyId   path      string  true  "IPAPolicy ID"
// @Success 200 {object} IPAPolicyResponse
// @Router /api/v1/ipapolicies/{ipaPolicyId} [get]
func NewGetIPAPolicy(graphService thegraph.TheGraphServiceBeta, httpClient xhttp.Client) func(c *gin.Context) {
	return func(c *gin.Context) {
		ipaPolicyId := c.Param("ipaPolicyId")

		policies, err := graphService.GetIPAPolicy(ipaPolicyId)
		if err != nil {
			logger.Errorf("Failed to get ipa policy: %v", err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, messages.ErrorMessage("Internal server error"))
			return
		}

		c.JSON(http.StatusOK, beta_v0.IPAPolicyResponse{
			Data: policies,
		})
	}
}

// @BasePath /

// ListIPAPolicies Example godoc
// @Summary List IPAPolicies
// @Schemes
// @Description Retrieve a paginated, filtered list of Policies
// @Security ApiKeyAuth
// @param X-API-Key header string true "API Key"
// @Param data body betav0.IPAPolicyRequestBody true "Query Parameters ("where" values are optional. Remove if not using)"
// @Tags IPAPolicies
// @Accept json
// @Produce json
// @Success 200 {object} IPAPoliciesResponse
// @Router /api/v1/ipapolicies [post]
func NewListIPAPolicies(graphService thegraph.TheGraphServiceBeta, httpClient xhttp.Client) func(c *gin.Context) {
	return func(c *gin.Context) {
		var requestBody beta_v0.IPAPolicyRequestBody
		if err := c.BindJSON(&requestBody); err != nil {
			logger.Errorf("Failed to read request body: %v", err)
			requestBody = beta_v0.IPAPolicyRequestBody{}
		}

		pols, err := graphService.ListIPAPolicies(fromIPAPolicyRequestQueryOptions(requestBody.Options))
		if err != nil {
			logger.Errorf("Failed to list ipa policies: %v", err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, messages.ErrorMessage("Internal server error"))
			return
		}

		c.JSON(http.StatusOK, beta_v0.IPAPoliciesResponse{
			Data: pols,
		})
	}
}

func fromIPAPolicyRequestQueryOptions(options *beta_v0.IPAPQueryOptions) *thegraph.TheGraphQueryOptions {
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

	queryOptions.Where.PolicyId = options.Where.PolicyId
	queryOptions.Where.Active = options.Where.Active
	queryOptions.Where.Inherited = options.Where.Inherited

	return queryOptions
}
