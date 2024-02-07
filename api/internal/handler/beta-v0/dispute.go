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

// GetDispute Example godoc
// @Summary Get a Dispute
// @Schemes
// @Description Retrieve a Dispute
// @Tags Disputes
// @Accept json
// @Produce json
// @Param        disputeId   path      string  true  "Dispute ID"
// @Security ApiKeyAuth
// @param X-API-Key header string true "API Key"
// @Success 200 {object} DisputeResponse
// @Router /disputes/{disputeId} [get]
func NewGetDispute(graphService thegraph.TheGraphServiceBeta, httpClient xhttp.Client) func(c *gin.Context) {
	return func(c *gin.Context) {
		disputeId := c.Param("disputeId")

		disputes, err := graphService.GetDispute(disputeId)
		if err != nil {
			logger.Errorf("Failed to get dispute: %v", err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, messages.ErrorMessage("Internal server error"))
			return
		}

		c.JSON(http.StatusOK, beta_v0.DisputeResponse{
			Data: disputes,
		})
	}
}

// @BasePath /api/v1

// ListDisputes Example godoc
// @Summary List Disputes
// @Schemes
// @Description Retrieve a paginated, filtered list of Disputes
// @Tags Disputes
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @param X-API-Key header string true "API Key"
// @Param data body options.RequestBody true "Query Parameters ("where" values are optional. Remove if not using)"
// @Success 200 {object} DisputesResponse
// @Router /disputes [post]
func NewListDisputes(graphService thegraph.TheGraphServiceBeta, httpClient xhttp.Client) func(c *gin.Context) {
	return func(c *gin.Context) {
		var requestBody options2.RequestBody
		if err := c.BindJSON(&requestBody); err != nil {
			logger.Errorf("Failed to read request body: %v", err)
			requestBody = options2.RequestBody{}
		}

		disputes, err := graphService.ListDisputes(thegraph.FromRequestQueryOptions(requestBody.Options))
		if err != nil {
			logger.Errorf("Failed to get added disputes: %v", err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, messages.ErrorMessage("Internal server error"))
			return
		}

		c.JSON(http.StatusOK, beta_v0.DisputesResponse{
			Data: disputes,
		})
	}
}
