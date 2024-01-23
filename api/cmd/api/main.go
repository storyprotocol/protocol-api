package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/machinebox/graphql"
	"github.com/storyprotocol/protocol-api/api/internal/config"
	"github.com/storyprotocol/protocol-api/api/internal/handler"
	"github.com/storyprotocol/protocol-api/api/internal/service/thegraph"
	xhttp "github.com/storyprotocol/protocol-api/pkg/http"
	"github.com/storyprotocol/protocol-api/pkg/logger"
)

func main() {
	r := gin.Default()
	flag.Parse()

	Logger, err := logger.InitLogger(logger.Levels.Info)
	if err != nil {
		logger.Fatalf("Failed to init logger, error: %v", err)
	}
	defer func() {
		_ = Logger.Sync()
	}()

	cfg, err := config.GetConfig()
	if err != nil {
		logger.Fatalf("Failed to init config, error: %v", err)
	}
	logger.Infof("cfg: %v", cfg)

	httpClient := xhttp.NewClient(&xhttp.ClientConfig{})

	theGraphAlphaClient := graphql.NewClient(cfg.TheGraphAlphaEndpoint)
	theGraphAlphaService := thegraph.NewTheGraphServiceMvpImpl(theGraphAlphaClient)

	// theGraphBeta
	theGraphBetaClient := graphql.NewClient(cfg.TheGraphBetaEndpoint)
	theGraphBetaService := thegraph.NewTheGraphServiceBetaImpl(theGraphBetaClient)

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "Hello")
	})

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, "OK")
	})

	protocol := r.Group("/")
	protocol.Use(cors.Default())
	{

		// BETA
		protocol.GET("/account/:accountId", handler.NewGetIPAccount(theGraphBetaService, httpClient))
		protocol.GET("/accounts", handler.NewGetIPAccounts(theGraphBetaService, httpClient))
		protocol.GET("/module/:moduleName", handler.NewGetModule(theGraphBetaService, httpClient))
		protocol.GET("/modules", handler.NewGetModules(theGraphBetaService, httpClient))

		//protocol.GET("/registeredIps", handler.NewGetIPsRegistered(theGraphBetaService, httpClient))
		//protocol.GET("/setAccounts", handler.NewGetSetIPAccounts(theGraphBetaService, httpClient))
		//protocol.GET("/setIpResolvers", handler.NewGetSetIPResolvers(theGraphBetaService, httpClient))
		//protocol.GET("/registeredModules", handler.NewGetRegisteredModules(theGraphBetaService, httpClient))
		//protocol.GET("/removedModules", handler.NewGetRemovedModules(theGraphBetaService, httpClient))

		// Endpoint to get franchises
		protocol.GET("/franchise", handler.NewGetFranchisesHandler(theGraphAlphaService, httpClient))

		// Endpoint to get a franchise
		protocol.GET("/franchise/:franchiseId", handler.NewGetFranchiseHandler(theGraphAlphaService, httpClient))

		// Endpoint to get ip assets from a franchise
		protocol.GET("/ipasset", handler.NewGetIpAssetsHandler(theGraphAlphaService, httpClient))

		// Endpoint to get a single ip asset from a franchise
		protocol.GET("/ipasset/:ipAssetId", handler.NewGetIpAssetHandler(theGraphAlphaService, httpClient))

		// Endpoint to get licenses from an ip asset
		protocol.GET("/license", handler.NewGetLicensesHandler(theGraphAlphaService))

		// Endpoint to get a single license
		protocol.GET("/license/:licenseId", handler.NewGetLicenseHandler(theGraphAlphaService))

		// Endpoint to get collections
		protocol.GET("/collection", handler.NewGetCollectionsHandler(theGraphAlphaService))

		// Endpoint to get transactions
		protocol.GET("/transaction", handler.NewGetTransactionsHandler(theGraphAlphaService))

		// Endpoint to get transaction
		protocol.GET("/transaction/:transactionId", handler.NewGetTransactionHandler(theGraphAlphaService))
	}

	port := fmt.Sprintf(":%d", cfg.Port)
	_ = r.Run(port)
}
