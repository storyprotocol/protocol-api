package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/storyprotocol/protocol-api/api/internal/entity"
	"github.com/storyprotocol/protocol-api/api/internal/service/thegraph"
	xhttp "github.com/storyprotocol/protocol-api/pkg/http"
	"github.com/storyprotocol/protocol-api/pkg/logger"
	"net/http"
)

func NewGetIPAccount(graphService thegraph.TheGraphServiceBeta, httpClient xhttp.Client) func(c *gin.Context) {
	return func(c *gin.Context) {
		accountId := c.Param("accountId")
		accounts, err := graphService.GetIPAccount(accountId)
		if err != nil {
			logger.Errorf("Failed to get account: %v", err)
			c.JSON(http.StatusInternalServerError, ErrorMessage("Internal server error"))
			return
		}

		c.JSON(http.StatusOK, entity.IPAccountsResponse{
			Data: accounts,
		})
	}
}

// GET /franchise
func NewListIPAccounts(graphService thegraph.TheGraphServiceBeta, httpClient xhttp.Client) func(c *gin.Context) {
	return func(c *gin.Context) {
		var requestBody entity.RequestBody
		if err := c.BindJSON(&requestBody); err != nil {
			logger.Errorf("Failed to read request body: %v", err)
			requestBody = entity.RequestBody{}
		}

		ipAccounts, err := graphService.ListIPAccounts(thegraph.FromRequestQueryOptions(requestBody.Options))
		if err != nil {
			logger.Errorf("Failed to get registered IP Accounts: %v", err)
			c.JSON(http.StatusInternalServerError, ErrorMessage("Internal server error"))
			return
		}

		c.JSON(http.StatusOK, entity.IPAccountsResponse{
			Data: ipAccounts,
		})
	}
}

func NewGetModule(graphService thegraph.TheGraphServiceBeta, httpClient xhttp.Client) func(c *gin.Context) {
	return func(c *gin.Context) {
		moduleName := c.Param("moduleName")

		mods, err := graphService.GetModule(moduleName)
		if err != nil {
			logger.Errorf("Failed to get module: %v", err)
			c.JSON(http.StatusInternalServerError, ErrorMessage("Internal server error"))
			return
		}

		c.JSON(http.StatusOK, entity.ModuleResponse{
			Data: mods,
		})
	}
}

func NewListModules(graphService thegraph.TheGraphServiceBeta, httpClient xhttp.Client) func(c *gin.Context) {
	return func(c *gin.Context) {
		var requestBody entity.RequestBody
		if err := c.BindJSON(&requestBody); err != nil {
			logger.Errorf("Failed to read request body: %v", err)
			requestBody = entity.RequestBody{}
		}

		mods, err := graphService.ListModules(thegraph.FromRequestQueryOptions(requestBody.Options))
		if err != nil {
			logger.Errorf("Failed to get added modules: %v", err)
			c.JSON(http.StatusInternalServerError, ErrorMessage("Internal server error"))
			return
		}

		c.JSON(http.StatusOK, entity.ModuleResponse{
			Data: mods,
		})
	}
}

//func NewGetIPsRegistered(graphService thegraph.TheGraphServiceBeta, httpClient xhttp.Client) func(c *gin.Context) {
//	return func(c *gin.Context) {
//		ips, err := graphService.GetIPsRegistered()
//		if err != nil {
//			logger.Errorf("Failed to get registered IPs: %v", err)
//			c.JSON(http.StatusInternalServerError, ErrorMessage("Internal server error"))
//			return
//		}
//
//		c.JSON(http.StatusOK, entity.IPRegisteredResponse{
//			Data: ips,
//		})
//	}
//}

//func NewGetSetIPAccounts(graphService thegraph.TheGraphServiceBeta, httpClient xhttp.Client) func(c *gin.Context) {
//	return func(c *gin.Context) {
//		accounts, err := graphService.GetSetIPAccounts()
//		if err != nil {
//			logger.Errorf("Failed to get set IP Accounts: %v", err)
//			c.JSON(http.StatusInternalServerError, ErrorMessage("Internal server error"))
//			return
//		}
//
//		c.JSON(http.StatusOK, entity.SetIPAccountResponse{
//			Data: accounts,
//		})
//	}
//}

//func NewGetSetIPResolvers(graphService thegraph.TheGraphServiceBeta, httpClient xhttp.Client) func(c *gin.Context) {
//	return func(c *gin.Context) {
//		accounts, err := graphService.GetSetIPResolvers()
//		if err != nil {
//			logger.Errorf("Failed to get set IP resolvers: %v", err)
//			c.JSON(http.StatusInternalServerError, ErrorMessage("Internal server error"))
//			return
//		}
//
//		c.JSON(http.StatusOK, entity.SetResolverResponse{
//			Data: accounts,
//		})
//	}
//}

//func NewGetModule(graphService thegraph.TheGraphServiceBeta, httpClient xhttp.Client) func(c *gin.Context) {
//
//	return func(c *gin.Context) {
//		mods, err := graphService.GetModule()
//		if err != nil {
//			logger.Errorf("Failed to get added modules: %v", err)
//			c.JSON(http.StatusInternalServerError, ErrorMessage("Internal server error"))
//			return
//		}
//
//		c.JSON(http.StatusOK, entity.ModuleAddedResponse{
//			Data: mods,
//		})
//	}
//}
