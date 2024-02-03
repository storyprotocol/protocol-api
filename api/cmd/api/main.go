package main

import (
	"flag"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	ginserver "github.com/go-oauth2/gin-server"
	"github.com/go-oauth2/oauth2/v4/manage"
	"github.com/go-oauth2/oauth2/v4/models"
	"github.com/go-oauth2/oauth2/v4/server"
	"github.com/go-oauth2/oauth2/v4/store"
	"github.com/machinebox/graphql"
	"github.com/storyprotocol/protocol-api/api/cmd/docs"
	"github.com/storyprotocol/protocol-api/api/internal/config"
	betaHandlers "github.com/storyprotocol/protocol-api/api/internal/handler/beta-v0"
	"github.com/storyprotocol/protocol-api/api/internal/service/thegraph/beta-v0"
	xhttp "github.com/storyprotocol/protocol-api/pkg/http"
	"github.com/storyprotocol/protocol-api/pkg/logger"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"net/http"
)

func main() {
	cfg, err := config.GetConfig()
	if err != nil {
		logger.Fatalf("Failed to init config, error: %v", err)
	}

	manager := manage.NewDefaultManager()

	// token store
	manager.MustTokenStorage(store.NewFileTokenStore("data.db"))
	// client store
	clientStore := store.NewClientStore()
	clientStore.Set(cfg.OAuthID, &models.Client{
		ID:     cfg.OAuthID,
		Secret: cfg.OAuthSecret,
		Domain: "*",
	})
	manager.MapClientStorage(clientStore)

	// Initialize the oauth2 service
	ginserver.InitServer(manager)
	ginserver.SetAllowGetAccessRequest(true)
	ginserver.SetClientInfoHandler(server.ClientFormHandler)

	r := gin.Default()
	flag.Parse()

	Logger, err := logger.InitLogger(logger.Levels.Info)
	if err != nil {
		logger.Fatalf("Failed to init logger, error: %v", err)
	}
	defer func() {
		_ = Logger.Sync()
	}()

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

	auth := r.Group("/oauth2")
	{
		auth.GET("/token", ginserver.HandleTokenRequest)
	}

	{
		// swagger
		docs.SwaggerInfo.BasePath = "/api/v1"
		v1 := r.Group("/api/v1")
		{
			protocol := v1.Group("/")
			protocol.Use(cors.Default())
			{
				protocol.Use(ginserver.HandleTokenVerify())
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
			}
		}
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	port := fmt.Sprintf(":%d", cfg.Port)
	_ = r.Run(port)
}
