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

// GetModule Example godoc
// @Summary Get a GetModule
// @Schemes
// @Description Retrieve a Module
// @Security ApiKeyAuth
// @Host https://edge.stg.storyprotocol.net
// @param X-API-Key header string true "API Key"
// @Tags Modules
// @Accept json
// @Produce json
// @Param        moduleId   path      string  true  "Module ID"
// @Success 200 {object} ModuleResponse
// @Router /api/v1/modules/{moduleId} [get]
func NewGetModule(graphService thegraph.TheGraphServiceBeta, httpClient xhttp.Client) func(c *gin.Context) {
	return func(c *gin.Context) {
		moduleId := c.Param("moduleId")

		mods, err := graphService.GetModule(moduleId)
		if err != nil {
			logger.Errorf("Failed to get module: %v", err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, messages.ErrorMessage("Internal server error"))
			return
		}

		c.JSON(http.StatusOK, beta_v0.ModuleResponse{
			Data: mods,
		})
	}
}

// @BasePath /

// ListModules Example godoc
// @Summary List Modules
// @Schemes
// @Description Retrieve a paginated, filtered list of Modules
// @Host https://edge.stg.storyprotocol.net
// @Security ApiKeyAuth
// @param X-API-Key header string true "API Key"
// @Param data body betav0.ModuleRequestBody true "Query Parameters ("where" values are optional. Remove if not using)"
// @Tags Modules
// @Accept json
// @Produce json
// @Success 200 {object} ModulesResponse
// @Router /api/v1/modules [post]
func NewListModules(graphService thegraph.TheGraphServiceBeta, httpClient xhttp.Client) func(c *gin.Context) {
	return func(c *gin.Context) {
		var requestBody beta_v0.ModuleRequestBody
		if err := c.BindJSON(&requestBody); err != nil {
			logger.Errorf("Failed to read request body: %v", err)
			requestBody = beta_v0.ModuleRequestBody{}
		}

		mods, err := graphService.ListModules(fromModuleRequestQueryOptions(requestBody.Options))
		if err != nil {
			logger.Errorf("Failed to get added modules: %v", err)
			c.JSON(http.StatusInternalServerError, messages.ErrorMessage("Internal server error"))
			return
		}

		c.JSON(http.StatusOK, beta_v0.ModulesResponse{
			Data: mods,
		})
	}
}

func fromModuleRequestQueryOptions(options *beta_v0.ModuleQueryOptions) *thegraph.TheGraphQueryOptions {
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

	queryOptions.Where.Name = options.Where.Name

	return queryOptions
}
