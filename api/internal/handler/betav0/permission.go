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

// GetPermission Example godoc
// @Summary Get a Permission
// @Schemes
// @Description Retrieve a Permission
// @Tags Permissions
// @Accept json
// @Produce json
// @Param        permissionId   path      string  true  "Permission ID"
// @Success 200 {object} PermissionResponse
// @Router /api/v1/permissions/{permissionId} [get]
func NewGetPermission(graphService thegraph.TheGraphServiceBeta, httpClient xhttp.Client) func(c *gin.Context) {
	return func(c *gin.Context) {
		permissionId := c.Param("permissionId")

		perms, err := graphService.GetPermission(permissionId)
		if err != nil {
			logger.Errorf("Failed to get permission: %v", err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, messages.ErrorMessage("Internal server error"))
			return
		}

		c.JSON(http.StatusOK, beta_v0.PermissionResponse{
			Data: perms,
		})
	}
}

// @BasePath /

// ListPermissions Example godoc
// @Summary List Permissions
// @Schemes
// @Description Retrieve a paginated, filtered list of Permissions
// @Security ApiKeyAuth
// @param X-API-Key header string true "API Key"
// @Param data body betav0.PermissionRequestBody true "Query Parameters ("where" values are optional. Remove if not using)"
// @Tags Permissions
// @Accept json
// @Produce json
// @Success 200 {object} PermissionsResponse
// @Router /api/v1/permissions [post]
func NewListPermissions(graphService thegraph.TheGraphServiceBeta, httpClient xhttp.Client) func(c *gin.Context) {
	return func(c *gin.Context) {
		var requestBody beta_v0.PermissionRequestBody
		if err := c.BindJSON(&requestBody); err != nil {
			logger.Errorf("Failed to read request body: %v", err)
			requestBody = beta_v0.PermissionRequestBody{}
		}

		perms, err := graphService.ListPermissions(fromPermissionRequestQueryOptions(requestBody.Options))
		if err != nil {
			logger.Errorf("Failed to list permissions: %v", err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, messages.ErrorMessage("Internal server error"))
			return
		}

		c.JSON(http.StatusOK, beta_v0.PermissionsResponse{
			Data: perms,
		})
	}
}

func fromPermissionRequestQueryOptions(options *beta_v0.PermissionQueryOptions) *thegraph.TheGraphQueryOptions {
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

	queryOptions.Where.Signer = options.Where.Signer
	queryOptions.Where.To = options.Where.To

	return queryOptions
}
