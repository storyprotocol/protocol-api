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

func NewGetPolicyFrameworkManager(graphService thegraph.TheGraphServiceBeta, httpClient xhttp.Client) func(c *gin.Context) {
	return func(c *gin.Context) {
		pfwmId := c.Param("pfwmId")

		pfwms, err := graphService.GetPolicyFrameworkManager(pfwmId)
		if err != nil {
			logger.Errorf("Failed to get policy framework manager: %v", err)
			c.JSON(http.StatusInternalServerError, messages.ErrorMessage("Internal server error"))
			return
		}

		c.JSON(http.StatusOK, beta_v0.PolicyFrameworkManagerResponse{
			Data: pfwms,
		})
	}
}

func NewListPolicyFrameworkManagers(graphService thegraph.TheGraphServiceBeta, httpClient xhttp.Client) func(c *gin.Context) {
	return func(c *gin.Context) {
		var requestBody options2.RequestBody
		if err := c.BindJSON(&requestBody); err != nil {
			logger.Errorf("Failed to read request body: %v", err)
			requestBody = options2.RequestBody{}
		}

		pfwms, err := graphService.ListPolicyFrameworkManagers(thegraph.FromRequestQueryOptions(requestBody.Options))
		if err != nil {
			logger.Errorf("Failed to list policy framework managers : %v", err)
			c.JSON(http.StatusInternalServerError, messages.ErrorMessage("Internal server error"))
			return
		}

		c.JSON(http.StatusOK, beta_v0.PolicyFrameworkManagersResponse{
			Data: pfwms,
		})
	}
}
