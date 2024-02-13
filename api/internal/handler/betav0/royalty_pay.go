package betav0

import (
	"github.com/gin-gonic/gin"
	"github.com/storyprotocol/protocol-api/api/internal/handler/betav0/messages"
	beta_v0 "github.com/storyprotocol/protocol-api/api/internal/models/betav0"
	options2 "github.com/storyprotocol/protocol-api/api/internal/models/betav0/options"
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
// @Security ApiKeyAuth
// @param X-API-Key header string true "API Key"
// @Tags RoyaltyPays
// @Accept json
// @Produce json
// @Param        royaltyPayId   path      string  true  "RoyaltyPay ID"
// @Success 200 {object} RoyaltyPayResponse
// @Router /api/v1/royaltypays/{royaltyPayId} [get]
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
// @param X-API-Key header string true "API Key"
// @Param data body options.RequestBody true "Query Parameters ("where" values are optional. Remove if not using)"
// @Tags RoyaltyPays
// @Accept json
// @Produce json
// @Success 200 {object} RoyaltyPaysResponse
// @Router /api/v1/royaltypays [post]
func NewListRoyaltyPays(graphService thegraph.TheGraphServiceBeta, httpClient xhttp.Client) func(c *gin.Context) {
	return func(c *gin.Context) {
		var requestBody options2.RequestBody
		if err := c.BindJSON(&requestBody); err != nil {
			logger.Errorf("Failed to read request body: %v", err)
			requestBody = options2.RequestBody{}
		}

		roys, err := graphService.ListRoyaltyPays(thegraph.FromRequestQueryOptions(requestBody.Options))
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
