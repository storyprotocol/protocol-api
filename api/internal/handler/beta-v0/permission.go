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

// GetPermission Example godoc
// @Summary Get a Permission
// @Schemes
// @Description Retrieve a Permission
// @Tags Permissions
// @Accept json
// @Produce json
// @Param        permissionId   path      string  true  "Permission ID"
// @Success 200 {object} PermissionResponse
// @Router /permissions/{permissionId} [get]
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

// @BasePath /api/v1

// ListPermissions Example godoc
// @Summary List Permissions
// @Schemes
// @Description Retrieve a paginated, filtered list of Permissions
// @Security ApiKeyAuth
// @param X-API-Key header string true "API Key"
// @Param data body options.RequestBody true "Query Parameters ("where" values are optional. Remove if not using)"
// @Tags Permissions
// @Accept json
// @Produce json
// @Success 200 {object} PermissionsResponse
// @Router /permissions [post]
func NewListPermissions(graphService thegraph.TheGraphServiceBeta, httpClient xhttp.Client) func(c *gin.Context) {
	return func(c *gin.Context) {
		var requestBody options2.RequestBody
		if err := c.BindJSON(&requestBody); err != nil {
			logger.Errorf("Failed to read request body: %v", err)
			requestBody = options2.RequestBody{}
		}

		perms, err := graphService.ListPermissions(thegraph.FromRequestQueryOptions(requestBody.Options))
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
