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
// @Summary Get a Royalty
// @Schemes
// @Description Retrieve a Royalty
// @Security ApiKeyAuth
// @param X-API-Key header string true "API Key"
// @Tags Royalties
// @Accept json
// @Produce json
// @Param        royaltyId   path      string  true  "Royalty ID"
// @Success 200 {object} RoyaltyResponse
// @Router /api/v1/royalties/{royaltyId} [get]
func NewGetRoyalty(graphService thegraph.TheGraphServiceBeta, httpClient xhttp.Client) func(c *gin.Context) {
	return func(c *gin.Context) {
		royaltyId := c.Param("royaltyId")

		roys, err := graphService.GetRoyalty(royaltyId)
		if err != nil {
			logger.Errorf("Failed to get royalty: %v", err)
			c.JSON(http.StatusInternalServerError, messages.ErrorMessage("Internal server error"))
			return
		}

		c.JSON(http.StatusOK, beta_v0.RoyaltyResponse{
			Data: roys,
		})
	}
}

// @BasePath /

// ListRoyalties Example godoc
// @Summary List Royalties
// @Schemes
// @Description Retrieve a paginated, filtered list of Royalties
// @Security ApiKeyAuth
// @param X-API-Key header string true "API Key"
// @Param data body options.RequestBody true "Query Parameters ("where" values are optional. Remove if not using)"
// @Tags Royalties
// @Accept json
// @Produce json
// @Success 200 {object} RoyaltiesResponse
// @Router /api/v1/royalties [post]
func NewListRoyalties(graphService thegraph.TheGraphServiceBeta, httpClient xhttp.Client) func(c *gin.Context) {
	return func(c *gin.Context) {
		var requestBody options2.RequestBody
		if err := c.BindJSON(&requestBody); err != nil {
			logger.Errorf("Failed to read request body: %v", err)
			requestBody = options2.RequestBody{}
		}

		roys, err := graphService.ListRoyalties(thegraph.FromRequestQueryOptions(requestBody.Options))
		if err != nil {
			logger.Errorf("Failed to list royalties: %v", err)
			c.JSON(http.StatusInternalServerError, messages.ErrorMessage("Internal server error"))
			return
		}

		c.JSON(http.StatusOK, beta_v0.RoyaltiesResponse{
			Data: roys,
		})
	}
}
