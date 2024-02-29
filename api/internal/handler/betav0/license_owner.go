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
		var requestBody *beta_v0.LicenseOwnersRequestBody
		if err := c.ShouldBindJSON(&requestBody); err != nil {

		}

		licenses, err := graphService.ListLicenseOwners(fromLicenseOwnerRequestQueryOptions(requestBody))
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

func fromLicenseOwnerRequestQueryOptions(requestBody *beta_v0.LicenseOwnersRequestBody) *thegraph.TheGraphQueryOptions {
	if requestBody == nil {
		return &thegraph.TheGraphQueryOptions{
			First: 100,
			Skip:  0,
		}
	}
	if requestBody.Options == nil {
		return &thegraph.TheGraphQueryOptions{
			First: 100,
			Skip:  0,
		}
	}

	var queryOptions = &thegraph.TheGraphQueryOptions{}

	if requestBody.Options.Pagination.Limit == 0 {
		requestBody.Options.Pagination.Limit = 100
	}

	queryOptions.First = requestBody.Options.Pagination.Limit
	queryOptions.Skip = requestBody.Options.Pagination.Offset
	queryOptions.OrderDirection = strings.ToLower(requestBody.Options.OrderDirection)
	queryOptions.OrderBy = requestBody.Options.OrderBy

	queryOptions.Where.PolicyId = requestBody.Options.Where.PolicyId
	queryOptions.Where.Owner = requestBody.Options.Where.Owner

	return queryOptions
}
