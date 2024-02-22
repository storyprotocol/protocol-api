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

// GetLicenseOwner Example godoc
// @Summary Get a LicenseOwner
// @Schemes
// @Description Retrieve a LicenseOwner
// @Tags Licenses
// @Host https://edge.stg.storyprotocol.net
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @param X-API-Key header string true "API Key"
// @Param        licenseOwnerId   path      string  true  "LicenseOwner ID"
// @Success 200 {object} LicenseOwnerResponse
// @Router /api/v1/licenses/owners/{licenseOwnerId} [get]
func NewGetLicenseOwner(graphService thegraph.TheGraphServiceBeta, httpClient xhttp.Client) func(c *gin.Context) {
	return func(c *gin.Context) {
		licenseOwnerId := c.Param("licenseOwnerId")

		licenses, err := graphService.GetLicenseOwner(licenseOwnerId)
		if err != nil {
			logger.Errorf("Failed to get license owner: %v", err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, messages.ErrorMessage("Internal server error"))
			return
		}

		c.JSON(http.StatusOK, beta_v0.LicenseOwnerResponse{
			Data: licenses,
		})
	}
}

// @BasePath /

// ListLicenseOwners Example godoc
// @Summary List LicenseOwners
// @Schemes
// @Description Retrieve a paginated, filtered list of LicenseOwners
// @Host https://edge.stg.storyprotocol.net
// @Tags Licenses
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @param X-API-Key header string true "API Key"
// @Param data body betav0.LicenseOwnersRequestBody true "Query Parameters ("where" values are optional. Remove if not using)"
// @Success 200 {object} LicenseOwnersResponse
// @Router /api/v1/licenses/owners [post]
func NewListLicenseOwners(graphService thegraph.TheGraphServiceBeta, httpClient xhttp.Client) func(c *gin.Context) {
	return func(c *gin.Context) {
		var requestBody beta_v0.LicenseOwnersRequestBody
		if err := c.BindJSON(&requestBody); err != nil {
			logger.Errorf("Failed to read request body: %v", err)
			requestBody = beta_v0.LicenseOwnersRequestBody{}
		}

		licenses, err := graphService.ListLicenseOwners(fromLicenseOwnerRequestQueryOptions(requestBody.Options))
		if err != nil {
			logger.Errorf("Failed to list license owners: %v", err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, messages.ErrorMessage("Internal server error"))
			return
		}

		c.JSON(http.StatusOK, beta_v0.LicenseOwnersResponse{
			Data: licenses,
		})
	}
}

func fromLicenseOwnerRequestQueryOptions(options *beta_v0.LicenseOwnerQueryOptions) *thegraph.TheGraphQueryOptions {
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
	queryOptions.Where.Owner = options.Where.Owner

	return queryOptions
}
