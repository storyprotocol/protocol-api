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

// GetLicenseFramework Example godoc
// @Summary Get a LicenseFramework
// @Schemes
// @Description Retrieve a LicenseFramework
// @Tags LicenseFrameworks
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @param X-API-Key header string true "API Key"
// @Param        frameworkId   path      string  true  "Framework ID"
// @Success 200 {object} LicenseFrameworkResponse
// @Router /api/v1/licenseframeworks/{frameworkId} [get]
func NewGetLicenseFramework(graphService thegraph.TheGraphServiceBeta, httpClient xhttp.Client) func(c *gin.Context) {
	return func(c *gin.Context) {
		licenseId := c.Param("frameworkId")

		licenses, err := graphService.GetLicenseFramework(licenseId)
		if err != nil {
			logger.Errorf("Failed to get license: %v", err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, messages.ErrorMessage("Internal server error"))
			return
		}

		c.JSON(http.StatusOK, beta_v0.LicenseFrameworkResponse{
			Data: licenses,
		})
	}
}

// @BasePath /

// ListLicenseFrameworks Example godoc
// @Summary List LicenseFrameworks
// @Schemes
// @Description Retrieve a paginated, filtered list of LicenseFrameworks
// @Tags LicenseFrameworks
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @param X-API-Key header string true "API Key"
// @Param data body betav0.LicenseFrameworkRequestBody true "Query Parameters ("where" values are optional. Remove if not using)"
// @Success 200 {object} LicenseFrameworksResponse
// @Router /api/v1/licenseframeworks [post]
func NewListLicenseFrameworks(graphService thegraph.TheGraphServiceBeta, httpClient xhttp.Client) func(c *gin.Context) {
	return func(c *gin.Context) {
		var requestBody beta_v0.LicenseFrameworkRequestBody
		if err := c.BindJSON(&requestBody); err != nil {
			logger.Errorf("Failed to read request body: %v", err)
			requestBody = beta_v0.LicenseFrameworkRequestBody{}
		}

		licenses, err := graphService.ListLicenseFrameworks(fromLicenseFWRequestQueryOptions(requestBody.Options))
		if err != nil {
			logger.Errorf("Failed to get license frameworks: %v", err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, messages.ErrorMessage("Internal server error"))
			return
		}

		c.JSON(http.StatusOK, beta_v0.LicenseFrameworksResponse{
			Data: licenses,
		})
	}
}

func fromLicenseFWRequestQueryOptions(options *beta_v0.LFWQueryOptions) *thegraph.TheGraphQueryOptions {
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

	queryOptions.Where.Creator = options.Where.Creator

	return queryOptions
}
