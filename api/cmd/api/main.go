package main

import (
	"flag"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/machinebox/graphql"
	"github.com/storyprotocol/protocol-api/api/cmd/docs"
	"github.com/storyprotocol/protocol-api/api/internal/config"
	betaHandlers "github.com/storyprotocol/protocol-api/api/internal/handler/beta-v0"
	"github.com/storyprotocol/protocol-api/api/internal/service/thegraph/beta-v0"
	xhttp "github.com/storyprotocol/protocol-api/pkg/http"
	"github.com/storyprotocol/protocol-api/pkg/logger"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
	"slices"
)

var ApiKeys []string

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
	ApiKeys = cfg.ApiKeys

	// theGraphBeta
	theGraphBetaClient := graphql.NewClient(cfg.TheGraphBetaEndpoint)
	theGraphBetaService := beta_v0.NewTheGraphServiceBetaImpl(theGraphBetaClient)

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "Hello")
	})

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, "OK")
	})

	docs.SwaggerInfo.BasePath = "/"
	v1 := r.Group("/api/v1")

	{
		protocol := v1.Group("/")
		protocol.Use(cors.Default())
		protocol.Use(AuthMiddleware())
		// BETA
		{
			protocol.GET("/assets/:assetId", betaHandlers.NewGetIPAsset(theGraphBetaService, httpClient))
			protocol.GET("/modules/:moduleId", betaHandlers.NewGetModule(theGraphBetaService, httpClient))
			//protocol.GET("/licenseframeworks/:frameworkId", betaHandlers.NewGetLicenseFramework(theGraphBetaService, httpClient))
			protocol.GET("/licenses/:licenseId", betaHandlers.NewGetLicense(theGraphBetaService, httpClient))
			protocol.GET("/policies/:policyId", betaHandlers.NewGetPolicy(theGraphBetaService, httpClient))
			protocol.GET("/ipapolicies/:ipaPolicyId", betaHandlers.NewGetIPAPolicy(theGraphBetaService, httpClient))
			protocol.GET("/disputes/:disputeId", betaHandlers.NewGetDispute(theGraphBetaService, httpClient))
			protocol.GET("/permissions/:permissionId", betaHandlers.NewGetPermission(theGraphBetaService, httpClient))
			protocol.GET("/tags/:tagId", betaHandlers.NewGetTag(theGraphBetaService, httpClient))
			protocol.GET("/royalties/:royaltyId", betaHandlers.NewGetRoyalty(theGraphBetaService, httpClient))
			protocol.GET("/royaltypays/:royaltyPayId", betaHandlers.NewGetRoyaltyPay(theGraphBetaService, httpClient))
			protocol.GET("/policyframeworks/:pfwmId", betaHandlers.NewGetPolicyFrameworkManager(theGraphBetaService, httpClient))
			protocol.GET("/collections/:collectionId", betaHandlers.NewGetCollection(theGraphBetaService, httpClient))

			protocol.POST("/assets", betaHandlers.NewListIPAssets(theGraphBetaService, httpClient))
			protocol.POST("/modules", betaHandlers.NewListModules(theGraphBetaService, httpClient))
			//protocol.POST("/licenseframeworks", betaHandlers.NewListLicenseFrameworks(theGraphBetaService, httpClient))
			protocol.POST("/licenses", betaHandlers.NewListLicenses(theGraphBetaService, httpClient))
			protocol.POST("/policies", betaHandlers.NewListPolicies(theGraphBetaService, httpClient))
			protocol.POST("/ipapolicies", betaHandlers.NewListIPAPolicies(theGraphBetaService, httpClient))
			protocol.POST("/disputes", betaHandlers.NewListDisputes(theGraphBetaService, httpClient))
			protocol.POST("/permissions", betaHandlers.NewListPermissions(theGraphBetaService, httpClient))
			protocol.POST("/tags", betaHandlers.NewListTags(theGraphBetaService, httpClient))
			protocol.POST("/royalties", betaHandlers.NewListRoyalties(theGraphBetaService, httpClient))
			protocol.POST("/royaltypays", betaHandlers.NewListRoyaltyPays(theGraphBetaService, httpClient))
			protocol.POST("/policyframeworks", betaHandlers.NewListPolicyFrameworkManagers(theGraphBetaService, httpClient))
			protocol.POST("/collections", betaHandlers.NewListCollections(theGraphBetaService, httpClient))
		}

	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	port := fmt.Sprintf(":%d", cfg.Port)
	_ = r.Run(port)
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		apiKey := c.GetHeader("X-API-Key")

		if !slices.Contains(ApiKeys, apiKey) {
			c.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized"})
			return
		}

		c.Next()
	}
}
