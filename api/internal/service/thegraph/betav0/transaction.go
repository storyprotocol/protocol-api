package betav0

import (
	"context"
	"fmt"
	beta_v0 "github.com/storyprotocol/protocol-api/api/internal/models/betav0"
	"github.com/storyprotocol/protocol-api/api/internal/service/thegraph"
)

func (c *ServiceBetaImpl) GetTransaction(transactionId string) (*beta_v0.Transaction, *beta_v0.RenTransaction, error) {

	if c.apiKey != "" {
		query := fmt.Sprintf(`
		query {
		  record(id: "%s")  {
			id
			created_at
			action_type
			initiator
			ip_id
			resource_id
			resource_type
			transaction_hash
		  }
		}
    `, transactionId)
		req := c.buildNewRequest(nil, query)
		ctx := context.Background()
		var trxRes beta_v0.RenTransactionTheGraphResponse
		if err := c.client.Run(ctx, req, &trxRes); err != nil {
			return nil, nil, fmt.Errorf("failed to get transaction from the graph. error: %v", err)
		}
		return nil, trxRes.Transaction, nil

	} else {
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
			return nil, nil, fmt.Errorf("failed to get transaction from the graph. error: %v", err)
		}
		return trxRes.Transaction, nil, nil

	}

}

func (c *ServiceBetaImpl) ListTransactions(options *thegraph.TheGraphQueryOptions) ([]*beta_v0.Transaction, []*beta_v0.RenTransaction, error) {
	whereString := c.buildWhereConditions(options)
	query := ""
	if c.apiKey != "" {
		VALUES := fmt.Sprintf(REN_QUERY_VALUE, ORDER_BY, ORDER_DIRECTION)

		query = fmt.Sprintf(`
		query(%s) {
		  records (%s, filter:{%s}) {
			transaction_index
			log_index
			id
			initiator
			ip_id
			resource_type
			block_hash
			resource_id
			action_type
			created_at
			block_number
			contract_address
			tx_hash
			transaction_hash
			block_time
			block_date
		  }
		}
		`, REN_QUERY_INTERFACE, VALUES, whereString)

		req := c.buildNewRequest(options, query)

		ctx := context.Background()
		var trxRes beta_v0.RenTransactionsTheGraphResponse
		if err := c.client.Run(ctx, req, &trxRes); err != nil {
			return nil, nil, fmt.Errorf("failed to get transactions from the graph. error: %v", err)
		}

		trxs := []*beta_v0.RenTransaction{}
		for _, trx := range trxRes.Transactions {
			trxs = append(trxs, trx)
		}

		return nil, trxs, nil

	} else {
		query = fmt.Sprintf(`
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
		}`, QUERY_INTERFACE, QUERY_VALUE, whereString)

		req := c.buildNewRequest(options, query)
		ctx := context.Background()
		var trxRes beta_v0.TransactionsTheGraphResponse
		if err := c.client.Run(ctx, req, &trxRes); err != nil {
			return nil, nil, fmt.Errorf("failed to get transactions from the graph. error: %v", err)
		}

		trxs := []*beta_v0.Transaction{}
		for _, trx := range trxRes.Transactions {
			trxs = append(trxs, trx)
		}

		return trxs, nil, nil

	}
}
