package betav0

import (
	"context"
	"fmt"
	"github.com/machinebox/graphql"
	beta_v0 "github.com/storyprotocol/protocol-api/api/internal/models/betav0"
	"github.com/storyprotocol/protocol-api/api/internal/service/thegraph"
)

func (c *ServiceBetaImpl) GetRoyaltyPay(royaltyPayId string) (*beta_v0.RoyaltyPay, *beta_v0.RenRoyaltyPay, error) {
	if c.apiKey != "" {
		query := fmt.Sprintf(`
		query {
		  record(id: "%s") {
			id
			receiver_ip_id
			payer_ip_id
			sender
			token
			amount
			block_time
			block_number
		  }
		}
    `, royaltyPayId)

		req := c.buildNewRequest(nil, query)
		ctx := context.Background()
		var royRes beta_v0.RenRoyaltyPayTheGraphResponse
		if err := c.client.Run(ctx, req, &royRes); err != nil {
			return nil, nil, fmt.Errorf("failed to get royalty pay from the graph. error: %v", err)
		}

		return nil, royRes.RoyaltyPay, nil
	} else {
		query := fmt.Sprintf(`
		query {
		  royaltyPay(id: "%s") {
			id
			receiverIpId
			payerIpId
			sender
			token
			amount
			blockTimestamp
			blockNumber
		  }
		}
    `, royaltyPayId)

		req := graphql.NewRequest(query)
		ctx := context.Background()
		var royRes beta_v0.RoyaltyPayTheGraphResponse
		if err := c.client.Run(ctx, req, &royRes); err != nil {
			return nil, nil, fmt.Errorf("failed to get royalty pay from the graph. error: %v", err)
		}

		return royRes.RoyaltyPay, nil, nil
	}

}

func (c *ServiceBetaImpl) ListRoyaltyPays(options *thegraph.TheGraphQueryOptions) ([]*beta_v0.RoyaltyPay, []*beta_v0.RenRoyaltyPay, error) {
	whereString := c.buildWhereConditions(options)
	query := ""

	if c.apiKey != "" {
		VALUES := fmt.Sprintf(REN_QUERY_VALUE, ORDER_BY, ORDER_DIRECTION)
		query = fmt.Sprintf(`
		query(%s){
		  records(%s, filter:{%s}) {
			id
			receiver_ip_id
			payer_ip_id
			sender
			token
			amount
			block_time
			block_number
		  }
			
		}
		`, REN_QUERY_INTERFACE, VALUES, whereString)

		req := c.buildNewRequest(options, query)

		ctx := context.Background()
		var roysRes beta_v0.RenRoyaltyPaysTheGraphResponse
		if err := c.client.Run(ctx, req, &roysRes); err != nil {
			return nil, nil, fmt.Errorf("failed to get royalty pays from the graph. error: %v", err)
		}

		roys := []*beta_v0.RenRoyaltyPay{}
		for _, roy := range roysRes.RoyaltyPays {
			roys = append(roys, roy)
		}

		return nil, roys, nil
	} else {
		query = fmt.Sprintf(`
	query(%s){
	  royaltyPays(%s, where:{%s}) {
			id
			receiverIpId
			payerIpId
			sender
			token
			amount
			blockTimestamp
			blockNumber
	  }
		
	}
    `, QUERY_INTERFACE, QUERY_VALUE, whereString)

		req := c.buildNewRequest(options, query)

		ctx := context.Background()
		var roysRes beta_v0.RoyaltyPaysTheGraphResponse
		if err := c.client.Run(ctx, req, &roysRes); err != nil {
			return nil, nil, fmt.Errorf("failed to get royalty pays from the graph. error: %v", err)
		}

		roys := []*beta_v0.RoyaltyPay{}
		for _, roy := range roysRes.RoyaltyPays {
			roys = append(roys, roy)
		}

		return roys, nil, nil
	}

}
