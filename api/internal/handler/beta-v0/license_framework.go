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

// @BasePath /api/v1

// ListLicenseFrameworks Example godoc
// @Summary List LicenseFrameworks
// @Schemes
// @Description Retrieve a paginated, filtered list of LicenseFrameworks
// @Tags LicenseFrameworks
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @param X-API-Key header string true "API Key"
// @Param data body options.RequestBody true "Query Parameters ("where" values are optional. Remove if not using)"
// @Success 200 {object} LicenseFrameworksResponse
// @Router /api/v1/licenseframeworks [post]
func NewListLicenseFrameworks(graphService thegraph.TheGraphServiceBeta, httpClient xhttp.Client) func(c *gin.Context) {
	return func(c *gin.Context) {
		var requestBody options2.RequestBody
		if err := c.BindJSON(&requestBody); err != nil {
			logger.Errorf("Failed to read request body: %v", err)
			requestBody = options2.RequestBody{}
		}

		licenses, err := graphService.ListLicenseFrameworks(thegraph.FromRequestQueryOptions(requestBody.Options))
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
