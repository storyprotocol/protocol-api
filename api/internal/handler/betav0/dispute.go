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
// @Router /api/v1/disputes/{disputeId} [get]
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

// @BasePath /

// ListDisputes Example godoc
// @Summary List Disputes
// @Schemes
// @Description Retrieve a paginated, filtered list of Disputes
// @Tags Disputes
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @param X-API-Key header string true "API Key"
// @Param data body betav0.DisputeRequestBody true "Query Parameters ("where" values are optional. Remove if not using)"
// @Success 200 {object} DisputesResponse
// @Router /api/v1/disputes [post]
func NewListDisputes(graphService thegraph.TheGraphServiceBeta, httpClient xhttp.Client) func(c *gin.Context) {
	return func(c *gin.Context) {
		var requestBody beta_v0.DisputeRequestBody
		if err := c.BindJSON(&requestBody); err != nil {
			logger.Errorf("Failed to read request body: %v", err)
			requestBody = beta_v0.DisputeRequestBody{}
		}

		disputes, err := graphService.ListDisputes(fromDisputeRequestQueryOptions(requestBody.Options))
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

func fromDisputeRequestQueryOptions(options *beta_v0.DisputeQueryOptions) *thegraph.TheGraphQueryOptions {
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

	queryOptions.Where.TargetTag = options.Where.TargetTag
	queryOptions.Where.TargetIpId = options.Where.TargetIpId
	queryOptions.Where.CurrentTag = options.Where.CurrentTag
	queryOptions.Where.Initiator = options.Where.Initiator

	return queryOptions
}
