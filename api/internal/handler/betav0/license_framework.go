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

func NewGetLicenseFramework(graphService thegraph.TheGraphServiceBeta, httpClient xhttp.Client) func(c *gin.Context) {
	return func(c *gin.Context) {
		licenseId := c.Param("frameworkId")

		licenses, renLfw, err := graphService.GetLicenseFramework(licenseId)
		if err != nil {
			logger.Errorf("Failed to get license: %v", err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, messages.ErrorMessage("Internal server error"))
			return
		}
		if licenses != nil {
			c.JSON(http.StatusOK, beta_v0.LicenseFrameworkResponse{
				Data: licenses,
			})
		} else {
			c.JSON(http.StatusOK, beta_v0.RenLicenseFrameworkResponse{
				Data: renLfw,
			})
		}

	}
}

func NewListLicenseFrameworks(graphService thegraph.TheGraphServiceBeta, httpClient xhttp.Client) func(c *gin.Context) {
	return func(c *gin.Context) {
		var requestBody *beta_v0.LicenseFrameworkRequestBody
		if err := c.ShouldBindJSON(&requestBody); err != nil {
		}

		licenses, renLfw, err := graphService.ListLicenseFrameworks(fromLicenseFWRequestQueryOptions(requestBody))
		if err != nil {
			logger.Errorf("Failed to get license frameworks: %v", err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, messages.ErrorMessage("Internal server error"))
			return
		}
		if licenses != nil {
			c.JSON(http.StatusOK, beta_v0.LicenseFrameworksResponse{
				Data: licenses,
			})
		} else {
			c.JSON(http.StatusOK, beta_v0.RenLicenseFrameworksResponse{
				Data: renLfw,
			})
		}

	}
}

func fromLicenseFWRequestQueryOptions(requestBody *beta_v0.LicenseFrameworkRequestBody) *thegraph.TheGraphQueryOptions {
	if requestBody == nil {
		return &thegraph.TheGraphQueryOptions{
			First: 100,
			Limit: 100,
			Skip:  0,
		}
	}
	if requestBody.Options == nil {
		return &thegraph.TheGraphQueryOptions{
			First: 100,
			Limit: 100,
			Skip:  0,
		}
	}

	var queryOptions = &thegraph.TheGraphQueryOptions{}

	if requestBody.Options.Pagination.Limit == 0 {
		requestBody.Options.Pagination.Limit = 100
	}

	queryOptions.First = requestBody.Options.Pagination.Limit
	queryOptions.Limit = requestBody.Options.Pagination.Limit
	queryOptions.Skip = requestBody.Options.Pagination.Offset
	queryOptions.OrderDirection = strings.ToLower(requestBody.Options.OrderDirection)
	queryOptions.OrderBy = requestBody.Options.OrderBy

	queryOptions.Where.Creator = requestBody.Options.Where.Creator

	return queryOptions
}
