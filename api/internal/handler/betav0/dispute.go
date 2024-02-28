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

// GetDispute Example godoc
// @Summary Get a Dispute
// @Schemes
// @Description Retrieve a Dispute
// @Tags Disputes
// @Host https://edge.stg.storyprotocol.net
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
// @Host https://edge.stg.storyprotocol.net
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
		var requestBody *beta_v0.DisputeRequestBody
		if err := c.ShouldBindJSON(&requestBody); err != nil {
		}

		disputes, err := graphService.ListDisputes(fromDisputeRequestQueryOptions(requestBody))
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

func fromDisputeRequestQueryOptions(requestBody *beta_v0.DisputeRequestBody) *thegraph.TheGraphQueryOptions {
	if requestBody.Options == nil {
		return &thegraph.TheGraphQueryOptions{
			First: 100,
			Skip:  0,
		}
	}

	var queryOptions = &thegraph.TheGraphQueryOptions{}

	if requestBody.Options.Pagination.Limit == 0 {
		requestBody.Options.Pagination.Limit = 100
	}

	queryOptions.First = requestBody.Options.Pagination.Limit
	queryOptions.Skip = requestBody.Options.Pagination.Offset
	queryOptions.OrderDirection = strings.ToLower(requestBody.Options.OrderDirection)
	queryOptions.OrderBy = requestBody.Options.OrderBy

	queryOptions.Where.TargetTag = requestBody.Options.Where.TargetTag
	queryOptions.Where.TargetIpId = requestBody.Options.Where.TargetIpId
	queryOptions.Where.CurrentTag = requestBody.Options.Where.CurrentTag
	queryOptions.Where.Initiator = requestBody.Options.Where.Initiator

	return queryOptions
}
