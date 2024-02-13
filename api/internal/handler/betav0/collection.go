package betav0

import (
	"github.com/gin-gonic/gin"
	"github.com/storyprotocol/protocol-api/api/internal/handler/betav0/messages"
	beta_v0 "github.com/storyprotocol/protocol-api/api/internal/models/betav0"
	options2 "github.com/storyprotocol/protocol-api/api/internal/models/betav0/options"
	"github.com/storyprotocol/protocol-api/api/internal/service/thegraph"
	xhttp "github.com/storyprotocol/protocol-api/pkg/http"
	"github.com/storyprotocol/protocol-api/pkg/logger"
	"net/http"
)

//var CollectionResponse =
// @BasePath /

// GetCollection Example godoc
// @Summary Get an Collection
// @Schemes
// @Description Retrieve a Collection
// @Tags Collections
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @param X-API-Key header string true "API Key"
// @Param        collectionId   path      string  true  "Collection ID"
// @Success 200 {object} CollectionResponse
// @Router /api/v1/collections/{collectionId} [get]
func NewGetCollection(graphService thegraph.TheGraphServiceBeta, httpClient xhttp.Client) func(c *gin.Context) {
	return func(c *gin.Context) {
		collectionId := c.Param("collectionId")

		cols, err := graphService.GetCollection(collectionId)
		if err != nil {
			logger.Errorf("Failed to get collection: %v", err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, messages.ErrorMessage("Internal server error"))
			return
		}

		c.JSON(http.StatusOK, beta_v0.CollectionResponse{
			Data: cols,
		})
	}
}

// @BasePath /

// ListCollections Example godoc
// @Summary List Collections
// @Schemes
// @Description Retrieve a paginated, filtered list of Collections
// @Tags Collections
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @param X-API-Key header string true "API Key"
// @Param data body options.RequestBody true "Query Parameters ("where" values are optional. Remove if not using)"
// @Success 200 {object} CollectionsResponse
// @Router /api/v1/collections [post]
func NewListCollections(graphService thegraph.TheGraphServiceBeta, httpClient xhttp.Client) func(c *gin.Context) {
	return func(c *gin.Context) {
		var requestBody options2.RequestBody
		if err := c.BindJSON(&requestBody); err != nil {
			logger.Errorf("Failed to read request body: %v", err)
			requestBody = options2.RequestBody{}
		}

		cols, err := graphService.ListCollections(thegraph.FromRequestQueryOptions(requestBody.Options))
		if err != nil {
			logger.Errorf("Failed to list licenses: %v", err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, messages.ErrorMessage("Internal server error"))
			return
		}

		c.JSON(http.StatusOK, beta_v0.CollectionsResponse{
			Data: cols,
		})
	}
}
