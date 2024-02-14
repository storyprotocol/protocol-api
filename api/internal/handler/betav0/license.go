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

// GetLicense Example godoc
// @Summary Get an License
// @Schemes
// @Description Retrieve a License
// @Tags Licenses
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @param X-API-Key header string true "API Key"
// @Param        licenseId   path      string  true  "License ID"
// @Success 200 {object} LicenseResponse
// @Router /api/v1/licenses/{licenseId} [get]
func NewGetLicense(graphService thegraph.TheGraphServiceBeta, httpClient xhttp.Client) func(c *gin.Context) {
	return func(c *gin.Context) {
		licenseId := c.Param("licenseId")

		licenses, err := graphService.GetLicense(licenseId)
		if err != nil {
			logger.Errorf("Failed to get license: %v", err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, messages.ErrorMessage("Internal server error"))
			return
		}

		c.JSON(http.StatusOK, beta_v0.LicenseResponse{
			Data: licenses,
		})
	}
}

// @BasePath /

// ListLicenses Example godoc
// @Summary List Licenses
// @Schemes
// @Description Retrieve a paginated, filtered list of Licenses
// @Tags Licenses
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @param X-API-Key header string true "API Key"
// @Param data body betav0.LicenseRequestBody true "Query Parameters ("where" values are optional. Remove if not using)"
// @Success 200 {object} LicensesResponse
// @Router /api/v1/licenses [post]
func NewListLicenses(graphService thegraph.TheGraphServiceBeta, httpClient xhttp.Client) func(c *gin.Context) {
	return func(c *gin.Context) {
		var requestBody beta_v0.LicenseRequestBody
		if err := c.BindJSON(&requestBody); err != nil {
			logger.Errorf("Failed to read request body: %v", err)
			requestBody = beta_v0.LicenseRequestBody{}
		}

		licenses, err := graphService.ListLicenses(fromLicenseRequestQueryOptions(requestBody.Options))
		if err != nil {
			logger.Errorf("Failed to list licenses: %v", err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, messages.ErrorMessage("Internal server error"))
			return
		}

		c.JSON(http.StatusOK, beta_v0.LicensesResponse{
			Data: licenses,
		})
	}
}

func fromLicenseRequestQueryOptions(options *beta_v0.LicenseQueryOptions) *thegraph.TheGraphQueryOptions {
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

	queryOptions.Where.PolicyId = options.Where.PolicyId
	queryOptions.Where.LicensorIpdId = options.Where.LicensorIpdId

	return queryOptions
}
