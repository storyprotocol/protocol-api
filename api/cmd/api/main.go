package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/machinebox/graphql"
	"github.com/storyprotocol/protocol-api/api/cmd/docs"
	"github.com/storyprotocol/protocol-api/api/internal/config"
	betaHandlers "github.com/storyprotocol/protocol-api/api/internal/handler/betav0"
	"github.com/storyprotocol/protocol-api/api/internal/service/thegraph/betav0"
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
	theGraphBetaService := betav0.NewTheGraphServiceBetaImpl(theGraphBetaClient, cfg.OpenChainLookupEndpoint)

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "Hello")
	})

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, "OK")
	})

	r.Use(CORSMiddleware())

	docs.SwaggerInfo.BasePath = "/"
	v1 := r.Group("/api/v1")

	{
		protocol := v1.Group("/")
		protocol.Use(AuthMiddleware())
		// BETA
		{
			royalties := protocol.Group("/royalties")
			{
				royalties.GET("/:royaltyId", betaHandlers.NewGetRoyalty(theGraphBetaService, httpClient))
				protocol.POST("/", betaHandlers.NewListRoyalties(theGraphBetaService, httpClient))

				royalties.GET("/payments/:royaltyPayId", betaHandlers.NewGetRoyaltyPay(theGraphBetaService, httpClient))
				protocol.POST("/payments", betaHandlers.NewListRoyaltyPays(theGraphBetaService, httpClient))

				royalties.GET("/policies/:royaltyPolicyId", betaHandlers.NewGetRoyaltyPolicy(theGraphBetaService, httpClient))
				protocol.POST("/policies", betaHandlers.NewListRoyaltyPolicies(theGraphBetaService, httpClient))

			}

			licenses := protocol.Group("/licenses")
			{
				licenses.GET("/:licenseId", betaHandlers.NewGetLicense(theGraphBetaService, httpClient))
				protocol.POST("/", betaHandlers.NewListLicenses(theGraphBetaService, httpClient))

				licenses.GET("/mintingfees/:licenseMintingFeePaidId", betaHandlers.NewGetLicenseMintingFeePay(theGraphBetaService, httpClient))
				protocol.POST("/mintingfees", betaHandlers.NewListLicenseMintingFeePaids(theGraphBetaService, httpClient))
			}

			policies := protocol.Group("/policies")
			{
				policies.GET("/:policyId", betaHandlers.NewGetPolicy(theGraphBetaService, httpClient))
				policies.POST("/", betaHandlers.NewListPolicies(theGraphBetaService, httpClient))

				policies.GET("/frameworks/:pfwmId", betaHandlers.NewGetPolicyFrameworkManager(theGraphBetaService, httpClient))
				policies.POST("/frameworks", betaHandlers.NewListPolicyFrameworkManagers(theGraphBetaService, httpClient))

			}

			protocol.GET("/assets/:assetId", betaHandlers.NewGetIPAsset(theGraphBetaService, httpClient))
			protocol.GET("/modules/:moduleId", betaHandlers.NewGetModule(theGraphBetaService, httpClient))
			protocol.GET("/ipapolicies/:ipaPolicyId", betaHandlers.NewGetIPAPolicy(theGraphBetaService, httpClient))
			protocol.GET("/disputes/:disputeId", betaHandlers.NewGetDispute(theGraphBetaService, httpClient))
			protocol.GET("/permissions/:permissionId", betaHandlers.NewGetPermission(theGraphBetaService, httpClient))
			protocol.GET("/tags/:tagId", betaHandlers.NewGetTag(theGraphBetaService, httpClient))
			protocol.GET("/collections/:collectionId", betaHandlers.NewGetCollection(theGraphBetaService, httpClient))
			protocol.GET("/transactions/:trxId", betaHandlers.NewGetTransaction(theGraphBetaService, httpClient))

			protocol.POST("/assets", betaHandlers.NewListIPAssets(theGraphBetaService, httpClient))
			protocol.POST("/modules", betaHandlers.NewListModules(theGraphBetaService, httpClient))
			protocol.POST("/ipapolicies", betaHandlers.NewListIPAPolicies(theGraphBetaService, httpClient))
			protocol.POST("/disputes", betaHandlers.NewListDisputes(theGraphBetaService, httpClient))
			protocol.POST("/permissions", betaHandlers.NewListPermissions(theGraphBetaService, httpClient))
			protocol.POST("/tags", betaHandlers.NewListTags(theGraphBetaService, httpClient))
			protocol.POST("/collections", betaHandlers.NewListCollections(theGraphBetaService, httpClient))
			protocol.POST("/transactions", betaHandlers.NewListTransactions(theGraphBetaService, httpClient))
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

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, X-API-Key, x-api-key, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
