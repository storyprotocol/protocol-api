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
		protocol.GET("/accounts/:accountId", handler.NewGetIPAccount(theGraphBetaService, httpClient))
		protocol.GET("/modules/:moduleName", handler.NewGetModule(theGraphBetaService, httpClient))
		protocol.GET("/licenseframeworks/:frameworkId", handler.NewGetLicenseFramework(theGraphBetaService, httpClient))
		protocol.GET("/licenses/:licenseId", handler.NewGetLicense(theGraphBetaService, httpClient))
		protocol.GET("/policies/:policyId", handler.NewGetPolicy(theGraphBetaService, httpClient))

		//protocol.GET("/permissions/:permissionId", handler.NewGetLicenseFramework(theGraphBetaService, httpClient))
		//protocol.GET("/tags/:tagId", handler.NewGetLicenseFramework(theGraphBetaService, httpClient))

		protocol.POST("/accounts", handler.NewListIPAccounts(theGraphBetaService, httpClient))
		protocol.POST("/modules", handler.NewListModules(theGraphBetaService, httpClient))
		protocol.POST("/licenseframeworks", handler.NewListLicenseFrameworks(theGraphBetaService, httpClient))
		protocol.POST("/licenses", handler.NewListLicenses(theGraphBetaService, httpClient))
		protocol.POST("/policies", handler.NewListPolicies(theGraphBetaService, httpClient))
		//protocol.POST("/permissions", handler.NewListAccessControlPermissions(theGraphBetaService, httpClient))
		//protocol.POST("/tags", handler.NewListAccessControlPermissions(theGraphBetaService, httpClient))
		// disputes
		// royalties
		// policy?
		// policyCreated <- licenseRegistry
	}

	port := fmt.Sprintf(":%d", cfg.Port)
	_ = r.Run(port)
}
