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

// GetIPAsset Example godoc
// @Summary Get an IPAsset
// @Schemes
// @Description Retrieve an IPAsset
// @Host https://edge.stg.storyprotocol.net
// @Tags IPAssets
// @Accept json
// @Produce json
// @Param        assetId   path      string  true  "Asset ID"
// @Security ApiKeyAuth
// @param X-API-Key header string true "API Key"
// @Success 200 {object} IPAssetResponse
// @Router /api/v1/assets/{assetId} [get]
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

// @BasePath /

// ListIPAssets Example godoc
// @Summary List IPAssets
// @Schemes
// @Description Retrieve a paginated, filtered list of IPAssets
// @Host https://edge.stg.storyprotocol.net
// @Tags IPAssets
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @param X-API-Key header string true "API Key"
// @Param data body betav0.IpAssetRequestBody true "Query Parameters ("where" values are optional. Remove if not using)"
// @Success 200 {object} IPAssetsResponse
// @Router /api/v1/assets [post]
func NewListIPAssets(graphService thegraph.TheGraphServiceBeta, httpClient xhttp.Client) func(c *gin.Context) {
	return func(c *gin.Context) {
		var requestBody *beta_v0.IpAssetRequestBody
		err := c.ShouldBindJSON(&requestBody)
		if err != nil {
			logger.Info(err)
		}

		ipAssets, err := graphService.ListIPAssets(fromIPARequestQueryOptions(requestBody))
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

func fromIPARequestQueryOptions(body *beta_v0.IpAssetRequestBody) *thegraph.TheGraphQueryOptions {
	if body == nil {
		return &thegraph.TheGraphQueryOptions{
			First: 10,
			Skip:  0,
		}
	}

	var queryOptions = &thegraph.TheGraphQueryOptions{}

	if body.Options.Pagination.Limit == 0 {
		body.Options.Pagination.Limit = 10
	}

	queryOptions.First = body.Options.Pagination.Limit
	queryOptions.Skip = body.Options.Pagination.Offset
	queryOptions.OrderDirection = strings.ToLower(body.Options.OrderDirection)
	queryOptions.OrderBy = body.Options.OrderBy

	queryOptions.Where.MetadataResolverAddress = body.Options.Where.MetadataResolverAddress
	queryOptions.Where.TokenContract = body.Options.Where.TokenContract
	queryOptions.Where.TokenId = body.Options.Where.TokenId
	queryOptions.Where.ChainId = body.Options.Where.ChainId

	return queryOptions
}
