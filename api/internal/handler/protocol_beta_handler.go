package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/storyprotocol/protocol-api/api/internal/entity"
	"github.com/storyprotocol/protocol-api/api/internal/service/thegraph"
	xhttp "github.com/storyprotocol/protocol-api/pkg/http"
	"github.com/storyprotocol/protocol-api/pkg/logger"
	"log"
	"net/http"
	"strconv"
)

func NewGetIPAccount(graphService thegraph.TheGraphServiceBeta, httpClient xhttp.Client) func(c *gin.Context) {
	return func(c *gin.Context) {
		accountId := c.Param("accountId")
		log.Println(accountId)
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
func NewGetIPAccounts(graphService thegraph.TheGraphServiceBeta, httpClient xhttp.Client) func(c *gin.Context) {
	return func(c *gin.Context) {
		limit := c.Query("limit")
		offset := c.Query("offset")

		limitI, offsetI, err := parseLimitAndOffset(limit, offset)
		if err != nil {
			logger.Errorf("Invalid limit or offset - offset: %s limit: %s", offset, limit)
			c.JSON(http.StatusBadRequest, ErrorMessage("Invalid offset or limit"))
			return
		}

		ipAccounts, err := graphService.GetIPAccounts(limitI, offsetI)
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

func NewGetModules(graphService thegraph.TheGraphServiceBeta, httpClient xhttp.Client) func(c *gin.Context) {
	return func(c *gin.Context) {
		limit := c.Query("limit")
		offset := c.Query("offset")

		limitI, offsetI, err := parseLimitAndOffset(limit, offset)
		if err != nil {
			logger.Errorf("Invalid limit or offset - offset: %s limit: %s", offset, limit)
			c.JSON(http.StatusBadRequest, ErrorMessage("Invalid offset or limit"))
			return
		}

		mods, err := graphService.GetModules(limitI, offsetI)
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

func parseLimitAndOffset(limit string, offset string) (int64, int64, error) {
	limitI, err := strconv.ParseInt(limit, 10, 64)
	if err != nil {
		logger.Errorf("Invalid limit: %s", limit)
		return 0, 0, err
	}
	offsetI, err := strconv.ParseInt(offset, 10, 64)
	if err != nil {
		logger.Errorf("Invalid offset: %s", offset)
		return 0, 0, err
	}

	return limitI, offsetI, nil
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
