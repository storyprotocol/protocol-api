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

// GetModule Example godoc
// @Summary Get a GetModule
// @Schemes
// @Description Retrieve a Module
// @Tags Modules
// @Accept json
// @Produce json
// @Param        moduleId   path      string  true  "Module ID"
// @Success 200 {object} ModuleResponse
// @Router /modules/{moduleId} [get]
func NewGetModule(graphService thegraph.TheGraphServiceBeta, httpClient xhttp.Client) func(c *gin.Context) {
	return func(c *gin.Context) {
		moduleId := c.Param("moduleId")

		mods, err := graphService.GetModule(moduleId)
		if err != nil {
			logger.Errorf("Failed to get module: %v", err)
			c.JSON(http.StatusInternalServerError, messages.ErrorMessage("Internal server error"))
			return
		}

		c.JSON(http.StatusOK, beta_v0.ModuleResponse{
			Data: mods,
		})
	}
}

// @BasePath /api/v1

// ListModules Example godoc
// @Summary List Modules
// @Schemes
// @Description Retrieve a paginated, filtered list of Modules
// @Tags Modules
// @Accept json
// @Produce json
// @Success 200 {object} ModulesResponse
// @Router /modules [post]
func NewListModules(graphService thegraph.TheGraphServiceBeta, httpClient xhttp.Client) func(c *gin.Context) {
	return func(c *gin.Context) {
		var requestBody options2.RequestBody
		if err := c.BindJSON(&requestBody); err != nil {
			logger.Errorf("Failed to read request body: %v", err)
			requestBody = options2.RequestBody{}
		}

		mods, err := graphService.ListModules(thegraph.FromRequestQueryOptions(requestBody.Options))
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
