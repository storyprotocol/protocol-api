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

	// theGraphBeta Sepolia
	theGraphSepoliaBetaClient := graphql.NewClient(cfg.TheGraphSepoliaBetaEndpoint)
	theGraphSepolia0xSplitBetaClient := graphql.NewClient(cfg.TheGraphSepolia0xSplitBetaEndpoint)
	theGraphSepoliaBetaService := betav0.NewTheGraphServiceBetaImpl(
		theGraphSepoliaBetaClient,
		cfg.OpenChainLookupEndpoint,
		"")

	theGraphSepolia0xSplitBetaService := betav0.NewTheGraphServiceBetaImpl(
		theGraphSepolia0xSplitBetaClient,
		cfg.OpenChainLookupEndpoint,
		"")

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "Hello")
	})

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, "OK")
	})

	r.Use(CORSMiddleware())

	docs.SwaggerInfo.BasePath = "/"
	sepoliaV1 := r.Group("/api/sepolia/v1")
	renaissanceV1 := r.Group("/api/renaissance/v1")

	{
		protocol := sepoliaV1.Group("/")
		protocol.Use(AuthMiddleware())

		// BETA
		{
			{
				protocol.GET("/royalties/payments/:royaltyPayId", betaHandlers.NewGetRoyaltyPay(theGraphSepoliaBetaService, httpClient))
				protocol.POST("/royalties/payments", betaHandlers.NewListRoyaltyPays(theGraphSepoliaBetaService, httpClient))

				protocol.GET("/royalties/policies/:royaltyPolicyId", betaHandlers.NewGetRoyaltyPolicy(theGraphSepoliaBetaService, httpClient))
				protocol.POST("/royalties/policies", betaHandlers.NewListRoyaltyPolicies(theGraphSepoliaBetaService, httpClient))

				protocol.GET("/royalties/splits/:royaltySplitId", betaHandlers.NewGetRoyaltyLiquidSplit(theGraphSepolia0xSplitBetaService, httpClient))

			}

			{
				protocol.GET("/licenses/:licenseId", betaHandlers.NewGetLicense(theGraphSepoliaBetaService, httpClient))
				protocol.POST("/licenses", betaHandlers.NewListLicenses(theGraphSepoliaBetaService, httpClient))

				protocol.GET("/licenses/mintingfees/:licenseMintingFeePaidId", betaHandlers.NewGetLicenseMintingFeePay(theGraphSepoliaBetaService, httpClient))
				protocol.POST("/licenses/mintingfees", betaHandlers.NewListLicenseMintingFeePaids(theGraphSepoliaBetaService, httpClient))

				protocol.GET("/licenses/owners/:licenseOwnerId", betaHandlers.NewGetLicenseOwner(theGraphSepoliaBetaService, httpClient))
				protocol.POST("/licenses/owners", betaHandlers.NewListLicenseOwners(theGraphSepoliaBetaService, httpClient))
			}

			{
				protocol.GET("/policies/:policyId", betaHandlers.NewGetPolicy(theGraphSepoliaBetaService, httpClient))
				protocol.POST("/policies", betaHandlers.NewListPolicies(theGraphSepoliaBetaService, httpClient))

				protocol.GET("/policies/frameworks/:pfwmId", betaHandlers.NewGetPolicyFrameworkManager(theGraphSepoliaBetaService, httpClient))
				protocol.POST("/policies/frameworks", betaHandlers.NewListPolicyFrameworkManagers(theGraphSepoliaBetaService, httpClient))

			}

			protocol.GET("/assets/:assetId", betaHandlers.NewGetIPAsset(theGraphSepoliaBetaService, httpClient))
			protocol.GET("/modules/:moduleId", betaHandlers.NewGetModule(theGraphSepoliaBetaService, httpClient))
			protocol.GET("/ipapolicies/:ipaPolicyId", betaHandlers.NewGetIPAPolicy(theGraphSepoliaBetaService, httpClient))
			protocol.GET("/disputes/:disputeId", betaHandlers.NewGetDispute(theGraphSepoliaBetaService, httpClient))
			protocol.GET("/permissions/:permissionId", betaHandlers.NewGetPermission(theGraphSepoliaBetaService, httpClient))
			//protocol.GET("/tags/:tagId", betaHandlers.NewGetTag(theGraphSepoliaBetaService, httpClient))
			protocol.GET("/collections/:collectionId", betaHandlers.NewGetCollection(theGraphSepoliaBetaService, httpClient))
			protocol.GET("/transactions/:trxId", betaHandlers.NewGetTransaction(theGraphSepoliaBetaService, httpClient))

			protocol.POST("/assets", betaHandlers.NewListIPAssets(theGraphSepoliaBetaService, httpClient))
			protocol.POST("/modules", betaHandlers.NewListModules(theGraphSepoliaBetaService, httpClient))
			protocol.POST("/ipapolicies", betaHandlers.NewListIPAPolicies(theGraphSepoliaBetaService, httpClient))
			protocol.POST("/disputes", betaHandlers.NewListDisputes(theGraphSepoliaBetaService, httpClient))
			protocol.POST("/permissions", betaHandlers.NewListPermissions(theGraphSepoliaBetaService, httpClient))
			//protocol.POST("/tags", betaHandlers.NewListTags(theGraphSepoliaBetaService, httpClient))
			protocol.POST("/collections", betaHandlers.NewListCollections(theGraphSepoliaBetaService, httpClient))
			protocol.POST("/transactions", betaHandlers.NewListTransactions(theGraphSepoliaBetaService, httpClient))
		}

	}

	// theGraphBeta Renaissance
	theGraphRenaissance0xSplitBetaClient := graphql.NewClient(cfg.TheGraphRenaissance0xSplitBetaEndpoint)
	theGraph0xSplitService := betav0.NewTheGraphServiceBetaImpl(
		theGraphRenaissance0xSplitBetaClient,
		cfg.OpenChainLookupEndpoint,
		cfg.ZettablockAPIKey)

	// transaction endpoint
	theGraphTransactionClient := graphql.NewClient(cfg.ZettablockTransactionEndpoint)
	theGraphTransactionService := betav0.NewTheGraphServiceBetaImpl(
		theGraphTransactionClient,
		cfg.OpenChainLookupEndpoint,
		cfg.ZettablockAPIKey)

	// collection endpoint
	theGraphCollectionClient := graphql.NewClient(cfg.ZettablockCollectionEndpoint)
	theGraphCollectionService := betav0.NewTheGraphServiceBetaImpl(
		theGraphCollectionClient,
		cfg.OpenChainLookupEndpoint,
		cfg.ZettablockAPIKey)

	// dispute endpoint
	theGraphDisputeClient := graphql.NewClient(cfg.ZettablockDisputeEndpoint)
	theGraphDisputeService := betav0.NewTheGraphServiceBetaImpl(
		theGraphDisputeClient,
		cfg.OpenChainLookupEndpoint,
		cfg.ZettablockAPIKey)

	// transaction endpoint
	theGraphIPAssetClient := graphql.NewClient(cfg.ZettablockIPAssetEndpoint)
	theGraphIPAssetService := betav0.NewTheGraphServiceBetaImpl(
		theGraphIPAssetClient,
		cfg.OpenChainLookupEndpoint,
		cfg.ZettablockAPIKey)

	// ipapolicy endpoint
	theGraphIPAPolicyClient := graphql.NewClient(cfg.ZettablockIPAPolicyEndpoint)
	theGraphIPAPolicyService := betav0.NewTheGraphServiceBetaImpl(
		theGraphIPAPolicyClient,
		cfg.OpenChainLookupEndpoint,
		cfg.ZettablockAPIKey)

	// license endpoint
	theGraphLicenseClient := graphql.NewClient(cfg.ZettablockLicenseEndpoint)
	theGraphLicenseService := betav0.NewTheGraphServiceBetaImpl(
		theGraphLicenseClient,
		cfg.OpenChainLookupEndpoint,
		cfg.ZettablockAPIKey)

	// licensemfp endpoint
	theGraphLicenseMfpClient := graphql.NewClient(cfg.ZettablockLicenseMfpEndpoint)
	theGraphLicenseMfpService := betav0.NewTheGraphServiceBetaImpl(
		theGraphLicenseMfpClient,
		cfg.OpenChainLookupEndpoint,
		cfg.ZettablockAPIKey)

	// licenseowner endpoint
	theGraphLicenseOwnerClient := graphql.NewClient(cfg.ZettablockLicenseOwnerEndpoint)
	theGraphLicenseOwnerService := betav0.NewTheGraphServiceBetaImpl(
		theGraphLicenseOwnerClient,
		cfg.OpenChainLookupEndpoint,
		cfg.ZettablockAPIKey)

	// module endpoint
	theGraphModuleClient := graphql.NewClient(cfg.ZettablockModuleEndpoint)
	theGraphModuleService := betav0.NewTheGraphServiceBetaImpl(
		theGraphModuleClient,
		cfg.OpenChainLookupEndpoint,
		cfg.ZettablockAPIKey)

	// permission endpoint
	theGraphPermissionClient := graphql.NewClient(cfg.ZettablockPermissionsEndpoint)
	theGraphPermissionService := betav0.NewTheGraphServiceBetaImpl(
		theGraphPermissionClient,
		cfg.OpenChainLookupEndpoint,
		cfg.ZettablockAPIKey)

	// policy endpoint
	theGraphPolicyClient := graphql.NewClient(cfg.ZettablockPolicyEndpoint)
	theGraphPolicyService := betav0.NewTheGraphServiceBetaImpl(
		theGraphPolicyClient,
		cfg.OpenChainLookupEndpoint,
		cfg.ZettablockAPIKey)

	// policyfwm endpoint
	theGraphPolicyFwmClient := graphql.NewClient(cfg.ZettablockPolicyFwmEndpoint)
	theGraphPolicyFwmService := betav0.NewTheGraphServiceBetaImpl(
		theGraphPolicyFwmClient,
		cfg.OpenChainLookupEndpoint,
		cfg.ZettablockAPIKey)

	// royaltypay endpoint
	theGraphRoyaltyPayClient := graphql.NewClient(cfg.ZettablockRoyaltyPayEndpoint)
	theGraphRoyaltyPayService := betav0.NewTheGraphServiceBetaImpl(
		theGraphRoyaltyPayClient,
		cfg.OpenChainLookupEndpoint,
		cfg.ZettablockAPIKey)

	// royaltypolicy endpoint
	theGraphRoyaltyPolicyClient := graphql.NewClient(cfg.ZettablockRoyaltyPolicyEndpoint)
	theGraphRoyaltyPolicyService := betav0.NewTheGraphServiceBetaImpl(
		theGraphRoyaltyPolicyClient,
		cfg.OpenChainLookupEndpoint,
		cfg.ZettablockAPIKey)

	{
		protocol2 := renaissanceV1.Group("/")
		protocol2.Use(AuthMiddleware())

		// BETA
		{
			{
				protocol2.GET("/royalties/payments/:royaltyPayId", betaHandlers.NewGetRoyaltyPay(theGraphRoyaltyPayService, httpClient))
				protocol2.POST("/royalties/payments", betaHandlers.NewListRoyaltyPays(theGraphRoyaltyPayService, httpClient))

				protocol2.GET("/royalties/policies/:royaltyPolicyId", betaHandlers.NewGetRoyaltyPolicy(theGraphRoyaltyPolicyService, httpClient))
				protocol2.POST("/royalties/policies", betaHandlers.NewListRoyaltyPolicies(theGraphRoyaltyPolicyService, httpClient))

				protocol2.GET("/royalties/splits/:royaltySplitId", betaHandlers.NewGetRoyaltyLiquidSplit(theGraph0xSplitService, httpClient))

			}

			{
				protocol2.GET("/licenses/:licenseId", betaHandlers.NewGetLicense(theGraphLicenseService, httpClient))
				protocol2.POST("/licenses", betaHandlers.NewListLicenses(theGraphLicenseService, httpClient))

				protocol2.GET("/licenses/mintingfees/:licenseMintingFeePaidId", betaHandlers.NewGetLicenseMintingFeePay(theGraphLicenseMfpService, httpClient))
				protocol2.POST("/licenses/mintingfees", betaHandlers.NewListLicenseMintingFeePaids(theGraphLicenseMfpService, httpClient))

				protocol2.GET("/licenses/owners/:licenseOwnerId", betaHandlers.NewGetLicenseOwner(theGraphLicenseOwnerService, httpClient))
				protocol2.POST("/licenses/owners", betaHandlers.NewListLicenseOwners(theGraphLicenseOwnerService, httpClient))
			}

			{
				protocol2.GET("/policies/:policyId", betaHandlers.NewGetPolicy(theGraphPolicyService, httpClient))
				protocol2.POST("/policies", betaHandlers.NewListPolicies(theGraphPolicyService, httpClient))

				protocol2.GET("/policies/frameworks/:pfwmId", betaHandlers.NewGetPolicyFrameworkManager(theGraphPolicyFwmService, httpClient))
				protocol2.POST("/policies/frameworks", betaHandlers.NewListPolicyFrameworkManagers(theGraphPolicyFwmService, httpClient))

			}

			protocol2.GET("/assets/:assetId", betaHandlers.NewGetIPAsset(theGraphIPAssetService, httpClient))
			protocol2.GET("/modules/:moduleId", betaHandlers.NewGetModule(theGraphModuleService, httpClient))
			protocol2.GET("/ipapolicies/:ipaPolicyId", betaHandlers.NewGetIPAPolicy(theGraphIPAPolicyService, httpClient))
			protocol2.GET("/disputes/:disputeId", betaHandlers.NewGetDispute(theGraphDisputeService, httpClient))
			protocol2.GET("/permissions/:permissionId", betaHandlers.NewGetPermission(theGraphPermissionService, httpClient))
			protocol2.GET("/collections/:collectionId", betaHandlers.NewGetCollection(theGraphCollectionService, httpClient))
			protocol2.GET("/transactions/:trxId", betaHandlers.NewGetTransaction(theGraphTransactionService, httpClient))

			protocol2.POST("/assets", betaHandlers.NewListIPAssets(theGraphIPAssetService, httpClient))
			protocol2.POST("/modules", betaHandlers.NewListModules(theGraphModuleService, httpClient))
			protocol2.POST("/ipapolicies", betaHandlers.NewListIPAPolicies(theGraphIPAPolicyService, httpClient))
			protocol2.POST("/disputes", betaHandlers.NewListDisputes(theGraphDisputeService, httpClient))
			protocol2.POST("/permissions", betaHandlers.NewListPermissions(theGraphPermissionService, httpClient))
			protocol2.POST("/collections", betaHandlers.NewListCollections(theGraphCollectionService, httpClient))
			protocol2.POST("/transactions", betaHandlers.NewListTransactions(theGraphTransactionService, httpClient))
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
