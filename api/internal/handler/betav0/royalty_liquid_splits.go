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

// GetRoyaltySplit Example godoc
// @Summary Get a Royalty Split
// @Schemes
// @Description Retrieve a Royalty Split
// @Security ApiKeyAuth
// @Host https://edge.stg.storyprotocol.net
// @param X-API-Key header string true "API Key"
// @Tags Royalties
// @Accept json
// @Produce json
// @Param        royaltySplitId   path      string  true  "Royalty Split ID"
// @Success 200 {object} RoyaltySplitResponse
// @Router /api/v1/royalties/splits/{royaltySplitId} [get]
func NewGetRoyaltyLiquidSplit(graphService thegraph.TheGraphServiceBeta, httpClient xhttp.Client) func(c *gin.Context) {
	return func(c *gin.Context) {
		royaltySplitId := c.Param("royaltySplitId")

		splits, renSplits, err := graphService.GetRoyaltyLiquidSplit(royaltySplitId)
		if err != nil {
			logger.Errorf("Failed to get royalty splits: %v", err)
			c.JSON(http.StatusInternalServerError, messages.ErrorMessage("Internal server error"))
			return
		}
		if splits != nil {
			c.JSON(http.StatusOK, beta_v0.RoyaltySplitResponse{
				Data: splits,
			})
		} else {
			c.JSON(http.StatusOK, beta_v0.RenRoyaltySplitResponse{
				Data: renSplits,
			})
		}

	}
}
