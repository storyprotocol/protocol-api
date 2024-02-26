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

// GetRoyalty Example godoc
// @Summary Get a RoyaltyPay
// @Schemes
// @Description Retrieve a RoyaltyPay
// @Host https://edge.stg.storyprotocol.net
// @Security ApiKeyAuth
// @param X-API-Key header string true "API Key"
// @Tags Royalties
// @Accept json
// @Produce json
// @Param        royaltyPayId   path      string  true  "RoyaltyPay ID"
// @Success 200 {object} RoyaltyPayResponse
// @Router /api/v1/royalties/payments/{royaltyPayId} [get]
func NewGetRoyaltyPay(graphService thegraph.TheGraphServiceBeta, httpClient xhttp.Client) func(c *gin.Context) {
	return func(c *gin.Context) {
		royaltyPayId := c.Param("royaltyPayId")

		roys, err := graphService.GetRoyaltyPay(royaltyPayId)
		if err != nil {
			logger.Errorf("Failed to get royalty pay: %v", err)
			c.JSON(http.StatusInternalServerError, messages.ErrorMessage("Internal server error"))
			return
		}

		c.JSON(http.StatusOK, beta_v0.RoyaltyPayResponse{
			Data: roys,
		})
	}
}

// @BasePath /

// ListRoyaltyPays Example godoc
// @Summary List RoyaltyPays
// @Schemes
// @Description Retrieve a paginated, filtered list of RoyaltyPays
// @Security ApiKeyAuth
// @Host https://edge.stg.storyprotocol.net
// @param X-API-Key header string true "API Key"
// @Param data body betav0.RoyaltyPayRequestBody true "Query Parameters ("where" values are optional. Remove if not using)"
// @Tags Royalties
// @Accept json
// @Produce json
// @Success 200 {object} RoyaltyPaysResponse
// @Router /api/v1/royalties/payments [post]
func NewListRoyaltyPays(graphService thegraph.TheGraphServiceBeta, httpClient xhttp.Client) func(c *gin.Context) {
	return func(c *gin.Context) {
		var requestBody *beta_v0.RoyaltyPayRequestBody
		if err := c.BindJSON(&requestBody); err != nil {
			logger.Errorf("Failed to read request body: %v", err)
			requestBody = &beta_v0.RoyaltyPayRequestBody{}
		}

		roys, err := graphService.ListRoyaltyPays(fromRoyaltyPayRequestQueryOptions(requestBody))
		if err != nil {
			logger.Errorf("Failed to list royalty pays: %v", err)
			c.JSON(http.StatusInternalServerError, messages.ErrorMessage("Internal server error"))
			return
		}

		c.JSON(http.StatusOK, beta_v0.RoyaltyPaysResponse{
			Data: roys,
		})
	}
}

func fromRoyaltyPayRequestQueryOptions(body *beta_v0.RoyaltyPayRequestBody) *thegraph.TheGraphQueryOptions {
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

	queryOptions.Where.ReceiverIpId = body.Options.Where.ReceiverIpId
	queryOptions.Where.Sender = body.Options.Where.Sender
	queryOptions.Where.Token = body.Options.Where.Token
	queryOptions.Where.PayerIpId = body.Options.Where.PayerIpId

	return queryOptions
}
