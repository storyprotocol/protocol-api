package main

import (
	"flag"
	"fmt"
	"github.com/storyprotocol/protocol-api/api/internal/service/thegraph/beta-v0"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/machinebox/graphql"
	"github.com/storyprotocol/protocol-api/api/internal/config"
	betaHandlers "github.com/storyprotocol/protocol-api/api/internal/handler/beta-v0"
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
	theGraphBetaService := beta_v0.NewTheGraphServiceBetaImpl(theGraphBetaClient)

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
		protocol.GET("/accounts/:accountId", betaHandlers.NewGetIPAccount(theGraphBetaService, httpClient))
		protocol.GET("/modules/:moduleId", betaHandlers.NewGetModule(theGraphBetaService, httpClient))
		protocol.GET("/licenseframeworks/:frameworkId", betaHandlers.NewGetLicenseFramework(theGraphBetaService, httpClient))
		protocol.GET("/licenses/:licenseId", betaHandlers.NewGetLicense(theGraphBetaService, httpClient))
		protocol.GET("/policies/:policyId", betaHandlers.NewGetPolicy(theGraphBetaService, httpClient))
		//protocol.GET("/disputes/:disputeId", betaHandlers.NewGetDispute(theGraphBetaService, httpClient))
		protocol.GET("/permissions/:permissionId", betaHandlers.NewGetPermission(theGraphBetaService, httpClient))
		protocol.GET("/tags/:tagId", betaHandlers.NewGetTag(theGraphBetaService, httpClient))

		protocol.POST("/accounts", betaHandlers.NewListIPAccounts(theGraphBetaService, httpClient))
		protocol.POST("/modules", betaHandlers.NewListModules(theGraphBetaService, httpClient))
		protocol.POST("/licenseframeworks", betaHandlers.NewListLicenseFrameworks(theGraphBetaService, httpClient))
		protocol.POST("/licenses", betaHandlers.NewListLicenses(theGraphBetaService, httpClient))
		protocol.POST("/policies", betaHandlers.NewListPolicies(theGraphBetaService, httpClient))
		//protocol.POST("/disputes", betaHandlers.NewListDisputes(theGraphBetaService, httpClient))
		protocol.POST("/permissions", betaHandlers.NewListPermissions(theGraphBetaService, httpClient))
		protocol.POST("/tags", betaHandlers.NewListTags(theGraphBetaService, httpClient))
		// royalties
		// policyCreated <- licenseRegistry
	}

	port := fmt.Sprintf(":%d", cfg.Port)
	_ = r.Run(port)
}
