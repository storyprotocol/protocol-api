package beta_v0

import (
	"github.com/gin-gonic/gin"
	"github.com/storyprotocol/protocol-api/api/internal/handler/beta-v0/messages"
	beta_v0 "github.com/storyprotocol/protocol-api/api/internal/models/beta-v0"
	options2 "github.com/storyprotocol/protocol-api/api/internal/models/beta-v0/options"
	"github.com/storyprotocol/protocol-api/api/internal/service/thegraph"
	xhttp "github.com/storyprotocol/protocol-api/pkg/http"
	"github.com/storyprotocol/protocol-api/pkg/logger"
	"net/http"
)

// @BasePath /api/v1

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
// @Router /ipapolicies/{ipaPolicyId} [get]
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

// @BasePath /api/v1

// ListIPAPolicies Example godoc
// @Summary List IPAPolicies
// @Schemes
// @Description Retrieve a paginated, filtered list of Policies
// @Security ApiKeyAuth
// @param X-API-Key header string true "API Key"
// @Param data body options.RequestBody true "Query Parameters ("where" values are optional. Remove if not using)"
// @Tags IPAPolicies
// @Accept json
// @Produce json
// @Success 200 {object} IPAPoliciesResponse
// @Router /ipapolicies [post]
func NewListIPAPolicies(graphService thegraph.TheGraphServiceBeta, httpClient xhttp.Client) func(c *gin.Context) {
	return func(c *gin.Context) {
		var requestBody options2.RequestBody
		if err := c.BindJSON(&requestBody); err != nil {
			logger.Errorf("Failed to read request body: %v", err)
			requestBody = options2.RequestBody{}
		}

		pols, err := graphService.ListIPAPolicies(thegraph.FromRequestQueryOptions(requestBody.Options))
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
