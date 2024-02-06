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

// GetPolicy Example godoc
// @Summary Get a Policy
// @Schemes
// @Description Retrieve a Policy
// @Tags Policies
// @Accept json
// @Produce json
// @Param        policyId   path      string  true  "Policy ID"
// @Success 200 {object} PolicyResponse
// @Router /policies/{policyId} [get]
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

// @BasePath /api/v1

// ListPolicies Example godoc
// @Summary List Policies
// @Schemes
// @Description Retrieve a paginated, filtered list of Policies
// @Tags Policies
// @Accept json
// @Produce json
// @Success 200 {object} PoliciesResponse
// @Router /policies [post]
func NewListPolicies(graphService thegraph.TheGraphServiceBeta, httpClient xhttp.Client) func(c *gin.Context) {
	return func(c *gin.Context) {
		var requestBody options2.RequestBody
		if err := c.BindJSON(&requestBody); err != nil {
			logger.Errorf("Failed to read request body: %v", err)
			requestBody = options2.RequestBody{}
		}

		pols, err := graphService.ListPolicies(thegraph.FromRequestQueryOptions(requestBody.Options))
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
