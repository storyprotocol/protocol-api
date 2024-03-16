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

// GetPermission Example godoc
// @Summary Get a Permission
// @Schemes
// @Description Retrieve a Permission
// @Host https://edge.stg.storyprotocol.net
// @Tags Permissions
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @param X-API-Key header string true "API Key"
// @Param        permissionId   path      string  true  "Permission ID"
// @Success 200 {object} PermissionResponse
// @Router /api/v1/permissions/{permissionId} [get]
func NewGetPermission(graphService thegraph.TheGraphServiceBeta, httpClient xhttp.Client) func(c *gin.Context) {
	return func(c *gin.Context) {
		permissionId := c.Param("permissionId")

		perms, renPerms, err := graphService.GetPermission(permissionId)
		if err != nil {
			logger.Errorf("Failed to get permission: %v", err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, messages.ErrorMessage("Internal server error"))
			return
		}
		if perms != nil {
			c.JSON(http.StatusOK, beta_v0.PermissionResponse{
				Data: perms,
			})
		} else {
			c.JSON(http.StatusOK, beta_v0.RenPermissionResponse{
				Data: renPerms,
			})
		}

	}
}

// @BasePath /

// ListPermissions Example godoc
// @Summary List Permissions
// @Schemes
// @Description Retrieve a paginated, filtered list of Permissions
// @Host https://edge.stg.storyprotocol.net
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
		var requestBody *beta_v0.PermissionRequestBody
		if err := c.ShouldBindJSON(&requestBody); err != nil {
		}

		perms, renPerms, err := graphService.ListPermissions(fromPermissionRequestQueryOptions(requestBody))
		if err != nil {
			logger.Errorf("Failed to list permissions: %v", err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, messages.ErrorMessage("Internal server error"))
			return
		}
		if perms != nil {
			c.JSON(http.StatusOK, beta_v0.PermissionsResponse{
				Data: perms,
			})
		} else {
			c.JSON(http.StatusOK, beta_v0.RenPermissionsResponse{
				Data: renPerms,
			})
		}

	}
}

func fromPermissionRequestQueryOptions(requestBody *beta_v0.PermissionRequestBody) *thegraph.TheGraphQueryOptions {
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

	queryOptions.Where.Signer = requestBody.Options.Where.Signer
	queryOptions.Where.To = requestBody.Options.Where.To

	return queryOptions
}
