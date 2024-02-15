package betav0

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/storyprotocol/protocol-api/api/internal/handler/betav0/messages"
	beta_v0 "github.com/storyprotocol/protocol-api/api/internal/models/betav0"
	"github.com/storyprotocol/protocol-api/api/internal/service/thegraph"
	xhttp "github.com/storyprotocol/protocol-api/pkg/http"
	"github.com/storyprotocol/protocol-api/pkg/logger"
	"net/http"
)

type TrxWhere struct {
}

// @BasePath /

// GetTransaction Example godoc
// @Summary Get a Dispute
// @Schemes
// @Description Retrieve a Transaction
// @Tags Transactions
// @Accept json
// @Produce json
// @Param        trxId   path      string  true  "Transaction ID"
// @Security ApiKeyAuth
// @param X-API-Key header string true "API Key"
// @Success 200 {object} TransactionResponse
// @Router /api/v1/transactions/{trxId} [get]
func NewGetTransaction(graphService thegraph.TheGraphServiceBeta, httpClient xhttp.Client) func(c *gin.Context) {
	return func(c *gin.Context) {
		trxId := c.Param("trxId")

		trx, err := graphService.GetTransaction(trxId)
		if err != nil {
			logger.Errorf("Failed to get transaction: %v", err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, messages.ErrorMessage("Internal server error"))
			return
		}

		c.JSON(http.StatusOK, beta_v0.TransactionResponse{
			Data: trx,
		})
	}
}

// @BasePath /

// ListTransactions Example godoc
// @Summary List Transactions
// @Schemes
// @Description Retrieve a paginated, filtered list of Transactions
// @Tags Transactions
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @param X-API-Key header string true "API Key"
// @Param data body betav0.TransactionRequestBody true "Query Parameters ("where" values are optional. Remove if not using)"
// @Success 200 {object} TransactionsResponse
// @Router /api/v1/transactions [post]
func NewListTransactions(graphService thegraph.TheGraphServiceBeta, httpClient xhttp.Client) func(c *gin.Context) {
	return func(c *gin.Context) {
		var requestBody beta_v0.TransactionRequestBody
		if err := c.BindJSON(&requestBody); err != nil {
			logger.Errorf("Failed to read request body: %v", err)
			requestBody = beta_v0.TransactionRequestBody{}
		}

		options := fromTrxRequestQueryOptions(requestBody.Options)

		fmt.Println(options)
		trxs, err := graphService.ListTransactions(options)
		if err != nil {
			logger.Errorf("Failed to get added transactions: %v", err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, messages.ErrorMessage("Internal server error"))
			return
		}

		c.JSON(http.StatusOK, beta_v0.TransactionsResponse{
			Data: trxs,
		})
	}
}

func fromTrxRequestQueryOptions(options *beta_v0.TrxQueryOptions) *thegraph.TheGraphQueryOptions {
	if options == nil {
		return &thegraph.TheGraphQueryOptions{
			First: 100,
			Skip:  0,
		}
	}

	var queryOptions = &thegraph.TheGraphQueryOptions{}

	if options.Pagination.Limit == 0 {
		options.Pagination.Limit = 100
	}

	queryOptions.First = options.Pagination.Limit
	queryOptions.Skip = options.Pagination.Offset

	queryOptions.Where.ActionType = options.Where.ActionType
	queryOptions.Where.ResourceId = options.Where.ResourceId

	return queryOptions
}