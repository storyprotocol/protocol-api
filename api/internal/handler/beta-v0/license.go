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

// GetLicense Example godoc
// @Summary Get an License
// @Schemes
// @Description Retrieve a License
// @Tags Licenses
// @Accept json
// @Produce json
// @Param        licenseId   path      string  true  "License ID"
// @Success 200 {object} LicenseResponse
// @Router /licenses/{licenseId} [get]
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

// @BasePath /api/v1

// ListLicenses Example godoc
// @Summary List Licenses
// @Schemes
// @Description Retrieve a paginated, filtered list of Licenses
// @Tags Licenses
// @Accept json
// @Produce json
// @Success 200 {object} LicensesResponse
// @Router /licenses [post]
func NewListLicenses(graphService thegraph.TheGraphServiceBeta, httpClient xhttp.Client) func(c *gin.Context) {
	return func(c *gin.Context) {
		var requestBody options2.RequestBody
		if err := c.BindJSON(&requestBody); err != nil {
			logger.Errorf("Failed to read request body: %v", err)
			requestBody = options2.RequestBody{}
		}

		licenses, err := graphService.ListLicenses(thegraph.FromRequestQueryOptions(requestBody.Options))
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
