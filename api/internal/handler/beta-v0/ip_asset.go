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

// GetIPAsset Example godoc
// @Summary Get an IPAsset
// @Schemes
// @Description Retrieve an IPAsset
// @Tags IPAssets
// @Accept json
// @Produce json
// @Param        assetId   path      string  true  "Asset ID"
// @Security ApiKeyAuth
// @param X-API-Key header string true "API Key"
// @Success 200 {object} IPAssetResponse
// @Router /assets/{assetId} [get]
func NewGetIPAsset(graphService thegraph.TheGraphServiceBeta, httpClient xhttp.Client) func(c *gin.Context) {
	return func(c *gin.Context) {
		assetId := c.Param("assetId")
		assets, err := graphService.GetIPAsset(assetId)
		if err != nil {
			logger.Errorf("Failed to get asset: %v", err)
			c.JSON(http.StatusInternalServerError, messages.ErrorMessage("Internal server error"))
			return
		}

		c.JSON(http.StatusOK, beta_v0.IPAssetResponse{
			Data: assets,
		})
	}
}

// @BasePath /api/v1

// ListIPAssets Example godoc
// @Summary List IPAssets
// @Schemes
// @Description Retrieve a paginated, filtered list of IPAssets
// @Tags Assets
// @Accept json
// @Produce json
// @Success 200 {object} IPAssetsResponse
// @Security ApiKeyAuth
// @param X-API-Key header string true "API Key"
// @Param data body options.RequestBody true "Query Parameters ("where" values are optional. Remove if not using)"
// @Router /assets [post]
func NewListIPAssets(graphService thegraph.TheGraphServiceBeta, httpClient xhttp.Client) func(c *gin.Context) {
	return func(c *gin.Context) {
		var requestBody options2.RequestBody
		if err := c.BindJSON(&requestBody); err != nil {
			logger.Errorf("Failed to read request body: %v", err)
			requestBody = options2.RequestBody{}
		}

		ipAssets, err := graphService.ListIPAssets(thegraph.FromRequestQueryOptions(requestBody.Options))
		if err != nil {
			logger.Errorf("Failed to get registered IP Assets: %v", err)
			c.JSON(http.StatusInternalServerError, messages.ErrorMessage("Internal server error"))
			return
		}

		c.JSON(http.StatusOK, beta_v0.IPAssetsResponse{
			Data: ipAssets,
		})
	}
}
