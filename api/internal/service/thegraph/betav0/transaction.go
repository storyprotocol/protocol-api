package betav0

import (
	"context"
	"fmt"

	beta_v0 "github.com/storyprotocol/protocol-api/api/internal/models/betav0"
	"github.com/storyprotocol/protocol-api/api/internal/service/thegraph"
)

func (c *ServiceBetaImpl) GetTransaction(transactionId string) (*beta_v0.Transaction, error) {
	query := fmt.Sprintf(`
		query {
		  transaction(id: "%s")  {
			id
			createdAt
			actionType
			initiator
			ipId
			resourceId
			resourceType
			txHash
		  }
		}
    `, transactionId)

	req := c.buildNewRequest(nil, query)
	ctx := context.Background()
	var trxRes beta_v0.TransactionTheGraphResponse
	if err := c.client.Run(ctx, req, &trxRes); err != nil {
		return nil, fmt.Errorf("failed to get transaction from the graph. error: %v", err)
	}

	return trxRes.Transaction, nil
}

func (c *ServiceBetaImpl) ListTransactions(options *thegraph.TheGraphQueryOptions) ([]*beta_v0.Transaction, error) {
	whereString := c.buildWhereConditions(options)
	query := fmt.Sprintf(`
	query(%s) {
	  transactions (%s, where:{%s}) {
		id
		createdAt
		actionType
		initiator
		ipId
		resourceId
		resourceType
		txHash
	  }
	}
    `, QUERY_INTERFACE, QUERY_VALUE, whereString)

	req := c.buildNewRequest(options, query)
	ctx := context.Background()
	var trxRes beta_v0.TransactionsTheGraphResponse
	if err := c.client.Run(ctx, req, &trxRes); err != nil {
		return nil, fmt.Errorf("failed to get transactions from the graph. error: %v", err)
	}

	trxs := []*beta_v0.Transaction{}
	for _, trx := range trxRes.Transactions {
		trxs = append(trxs, trx)
	}

	return trxs, nil
}
