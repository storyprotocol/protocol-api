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

// GetLicenseMintingFeePay Example godoc
// @Summary Get a LicenseMintingFeePay
// @Schemes
// @Description Retrieve a LicenseMintingFeePay
// @Security ApiKeyAuth
// @Host https://edge.stg.storyprotocol.net
// @param X-API-Key header string true "API Key"
// @Tags Licenses
// @Accept json
// @Produce json
// @Param        licenseMintingFeePaidId   path      string  true  "LicenseMintingFeePay ID"
// @Success 200 {object} LicenseMintingFeePaidResponse
// @Router /api/v1/licenses/mintingfees/{licenseMintingFeePaidId} [get]
func NewGetLicenseMintingFeePay(graphService thegraph.TheGraphServiceBeta, httpClient xhttp.Client) func(c *gin.Context) {
	return func(c *gin.Context) {
		licenseMintingFeePaidId := c.Param("licenseMintingFeePaidId")

		lmfp, err := graphService.GetLicenseMintingFeePaid(licenseMintingFeePaidId)
		if err != nil {
			logger.Errorf("Failed to get license minting fee paid: %v", err)
			c.JSON(http.StatusInternalServerError, messages.ErrorMessage("Internal server error"))
			return
		}

		c.JSON(http.StatusOK, beta_v0.LicenseMintingFeePaidResponse{
			Data: lmfp,
		})
	}
}

// @BasePath /

// ListLicenseMintingFeePaids Example godoc
// @Summary List LicenseMintingFeePays
// @Schemes
// @Description Retrieve a paginated, filtered list of LicenseMintingFeePaids
// @Host https://edge.stg.storyprotocol.net
// @Security ApiKeyAuth
// @param X-API-Key header string true "API Key"
// @Param data body betav0.RoyaltyRequestBody true "Query Parameters ("where" values are optional. Remove if not using)"
// @Tags Licenses
// @Accept json
// @Produce json
// @Success 200 {object} LicenseMintingFeePaidsResponse
// @Router /api/v1/licenses/mintingfees [post]
func NewListLicenseMintingFeePaids(graphService thegraph.TheGraphServiceBeta, httpClient xhttp.Client) func(c *gin.Context) {
	return func(c *gin.Context) {
		var requestBody beta_v0.LicenseMintingFeePaidRequestBody
		if err := c.BindJSON(&requestBody); err != nil {
			logger.Errorf("Failed to read request body: %v", err)
			requestBody = beta_v0.LicenseMintingFeePaidRequestBody{}
		}

		lmfps, err := graphService.ListLicenseMintingFeePaids(fromLicenseMintingFeePaysRequestBodyRequestQueryOptions(requestBody.Options))
		if err != nil {
			logger.Errorf("Failed to list license minting fee paids: %v", err)
			c.JSON(http.StatusInternalServerError, messages.ErrorMessage("Internal server error"))
			return
		}

		c.JSON(http.StatusOK, beta_v0.LicenseMintingFeePaidsResponse{
			Data: lmfps,
		})
	}
}

func fromLicenseMintingFeePaysRequestBodyRequestQueryOptions(options *beta_v0.LicenseMintingFeePaidQueryOptions) *thegraph.TheGraphQueryOptions {
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

	return queryOptions
}
