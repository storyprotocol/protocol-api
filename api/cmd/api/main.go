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

	theGraphClient := graphql.NewClient(cfg.TheGraphEndpoint)
	theGraphService := thegraph.NewTheGraphServiceMvpImpl(theGraphClient)

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "Hello")
	})

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, "OK")
	})

	protocol := r.Group("/")
	protocol.Use(cors.Default())
	{
		// Endpoint to get franchises
		protocol.GET("/franchise", handler.NewGetFranchisesHandler(theGraphService, httpClient))

		// Endpoint to get a franchise
		protocol.GET("/franchise/:franchiseId", handler.NewGetFranchiseHandler(theGraphService, httpClient))

		// Endpoint to get ip assets from a franchise
		protocol.GET("/ipasset", handler.NewGetIpAssetsHandler(theGraphService, httpClient))

		// Endpoint to get a single ip asset from a franchise
		protocol.GET("/ipasset/:ipAssetId", handler.NewGetIpAssetHandler(theGraphService, httpClient))

		// Endpoint to get licenses from an ip asset
		protocol.GET("/license", handler.NewGetLicensesHandler(theGraphService))

		// Endpoint to get a single license
		protocol.GET("/license/:licenseId", handler.NewGetLicenseHandler(theGraphService))

		// Endpoint to get collections
		protocol.GET("/collection", handler.NewGetCollectionsHandler(theGraphService))

		// Endpoint to get transactions
		protocol.GET("/transaction", handler.NewGetTransactionsHandler(theGraphService))

		// Endpoint to get transaction
		protocol.GET("/transaction/:transactionId", handler.NewGetTransactionHandler(theGraphService))
	}

	port := fmt.Sprintf(":%d", cfg.Port)
	_ = r.Run(port)
}
