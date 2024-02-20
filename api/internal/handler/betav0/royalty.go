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
// @Summary Get a Royalty
// @Schemes
// @Description Retrieve a Royalty
// @Security ApiKeyAuth
// @Host https://edge.stg.storyprotocol.net
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
// @Host https://edge.stg.storyprotocol.net
// @Security ApiKeyAuth
// @param X-API-Key header string true "API Key"
// @Param data body betav0.RoyaltyRequestBody true "Query Parameters ("where" values are optional. Remove if not using)"
// @Tags Royalties
// @Accept json
// @Produce json
// @Success 200 {object} RoyaltiesResponse
// @Router /api/v1/royalties [post]
func NewListRoyalties(graphService thegraph.TheGraphServiceBeta, httpClient xhttp.Client) func(c *gin.Context) {
	return func(c *gin.Context) {
		var requestBody beta_v0.RoyaltyRequestBody
		if err := c.BindJSON(&requestBody); err != nil {
			logger.Errorf("Failed to read request body: %v", err)
			requestBody = beta_v0.RoyaltyRequestBody{}
		}

		roys, err := graphService.ListRoyalties(fromRoyaltyRequestQueryOptions(requestBody.Options))
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

func fromRoyaltyRequestQueryOptions(options *beta_v0.RoyaltyQueryOptions) *thegraph.TheGraphQueryOptions {
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

	queryOptions.Where.IPID = options.Where.IPID
	queryOptions.Where.RoyaltyPolicy = options.Where.RoyaltyPolicy

	return queryOptions
}
