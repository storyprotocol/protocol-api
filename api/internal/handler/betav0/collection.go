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

//var CollectionResponse =
// @BasePath /

// GetCollection Example godoc
// @Summary Get an Collection
// @Schemes
// @Description Retrieve a Collection
// @Host https://edge.stg.storyprotocol.net
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

		cols, renCols, err := graphService.GetCollection(collectionId)
		if err != nil {
			logger.Errorf("Failed to get collection: %v", err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, messages.ErrorMessage("Internal server error"))
			return
		}

		if cols != nil {
			c.JSON(http.StatusOK, beta_v0.CollectionResponse{
				Data: cols,
			})
		} else {
			c.JSON(http.StatusOK, beta_v0.RenCollectionResponse{
				Data: renCols,
			})
		}
	}
}

// @BasePath /

// ListCollections Example godoc
// @Summary List Collections
// @Schemes
// @Description Retrieve a paginated, filtered list of Collections
// @host https://edge.stg.storyprotocol.net
// @Tags Collections
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @param X-API-Key header string true "API Key"
// @Param data body betav0.CollectionsRequestBody true "Query Parameters ("where" values are optional. Remove if not using)"
// @Success 200 {object} CollectionsResponse
// @Router /api/v1/collections [post]
func NewListCollections(graphService thegraph.TheGraphServiceBeta, httpClient xhttp.Client) func(c *gin.Context) {
	return func(c *gin.Context) {
		var requestBody *beta_v0.CollectionsRequestBody
		if err := c.ShouldBindJSON(&requestBody); err != nil {
		}

		cols, renCols, err := graphService.ListCollections(fromCollectionsRequestQueryOptions(requestBody))
		if err != nil {
			logger.Errorf("Failed to list collections: %v", err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, messages.ErrorMessage("Internal server error"))
			return
		}

		if cols != nil {
			c.JSON(http.StatusOK, beta_v0.CollectionsResponse{
				Data: cols,
			})
		} else {
			c.JSON(http.StatusOK, beta_v0.RenCollectionsResponse{
				Data: renCols,
			})
		}
	}
}

func fromCollectionsRequestQueryOptions(requestBody *beta_v0.CollectionsRequestBody) *thegraph.TheGraphQueryOptions {
	if requestBody == nil {
		return &thegraph.TheGraphQueryOptions{
			First: 100,
			Skip:  0,
			Limit: 100,
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

	return queryOptions
}
