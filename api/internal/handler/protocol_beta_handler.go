package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/storyprotocol/protocol-api/api/internal/entity"
	"github.com/storyprotocol/protocol-api/api/internal/service/thegraph"
	xhttp "github.com/storyprotocol/protocol-api/pkg/http"
	"github.com/storyprotocol/protocol-api/pkg/logger"
	"net/http"
)

// GET /franchise
func NewGetIPAccountsRegistered(graphService thegraph.TheGraphServiceBeta, httpClient xhttp.Client) func(c *gin.Context) {
	return func(c *gin.Context) {
		ipAccounts, err := graphService.GetIPAccountsRegistered()
		if err != nil {
			logger.Errorf("Failed to get registered IP Accounts: %v", err)
			c.JSON(http.StatusInternalServerError, ErrorMessage("Internal server error"))
			return
		}

		c.JSON(http.StatusOK, entity.IPAccountsRegisteredResponse{
			Data: ipAccounts,
		})
	}
}

func NewGetIPsRegistered(graphService thegraph.TheGraphServiceBeta, httpClient xhttp.Client) func(c *gin.Context) {
	return func(c *gin.Context) {
		ips, err := graphService.GetIPsRegistered()
		if err != nil {
			logger.Errorf("Failed to get registered IPs: %v", err)
			c.JSON(http.StatusInternalServerError, ErrorMessage("Internal server error"))
			return
		}

		c.JSON(http.StatusOK, entity.IPRegisteredResponse{
			Data: ips,
		})
	}
}

func NewGetSetIPAccounts(graphService thegraph.TheGraphServiceBeta, httpClient xhttp.Client) func(c *gin.Context) {
	return func(c *gin.Context) {
		accounts, err := graphService.GetSetIPAccounts()
		if err != nil {
			logger.Errorf("Failed to get set IP Accounts: %v", err)
			c.JSON(http.StatusInternalServerError, ErrorMessage("Internal server error"))
			return
		}

		c.JSON(http.StatusOK, entity.SetIPAccountResponse{
			Data: accounts,
		})
	}
}

func NewGetSetIPResolvers(graphService thegraph.TheGraphServiceBeta, httpClient xhttp.Client) func(c *gin.Context) {
	return func(c *gin.Context) {
		accounts, err := graphService.GetSetIPResolvers()
		if err != nil {
			logger.Errorf("Failed to get set IP resolvers: %v", err)
			c.JSON(http.StatusInternalServerError, ErrorMessage("Internal server error"))
			return
		}

		c.JSON(http.StatusOK, entity.SetResolverResponse{
			Data: accounts,
		})
	}
}

func NewGetRegisteredModules(graphService thegraph.TheGraphServiceBeta, httpClient xhttp.Client) func(c *gin.Context) {
	return func(c *gin.Context) {
		mods, err := graphService.GetRegisteredModules()
		if err != nil {
			logger.Errorf("Failed to get added modules: %v", err)
			c.JSON(http.StatusInternalServerError, ErrorMessage("Internal server error"))
			return
		}

		c.JSON(http.StatusOK, entity.ModuleAddedResponse{
			Data: mods,
		})
	}
}

func NewGetRemovedModules(graphService thegraph.TheGraphServiceBeta, httpClient xhttp.Client) func(c *gin.Context) {
	return func(c *gin.Context) {
		mods, err := graphService.GetRemovedModules()
		if err != nil {
			logger.Errorf("Failed to get removed modules: %v", err)
			c.JSON(http.StatusInternalServerError, ErrorMessage("Internal server error"))
			return
		}

		c.JSON(http.StatusOK, entity.ModuleRemovedResponse{
			Data: mods,
		})
	}
}
