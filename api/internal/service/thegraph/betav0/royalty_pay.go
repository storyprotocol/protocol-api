package betav0

import (
	"context"
	"fmt"
	"github.com/machinebox/graphql"
	beta_v0 "github.com/storyprotocol/protocol-api/api/internal/models/betav0"
	"github.com/storyprotocol/protocol-api/api/internal/service/thegraph"
)

func (c *ServiceBetaImpl) GetRoyaltyPay(royaltyPayId string) (*beta_v0.RoyaltyPay, error) {
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
		return nil, fmt.Errorf("failed to get royalty pay from the graph. error: %v", err)
	}

	return royRes.RoyaltyPay, nil

}

func (c *ServiceBetaImpl) ListRoyaltyPays(options *thegraph.TheGraphQueryOptions) ([]*beta_v0.RoyaltyPay, error) {
	whereString := c.buildWhereConditions(options)
	query := fmt.Sprintf(`
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
		return nil, fmt.Errorf("failed to get royalty pays from the graph. error: %v", err)
	}

	roys := []*beta_v0.RoyaltyPay{}
	for _, roy := range roysRes.RoyaltyPays {
		roys = append(roys, roy)
	}

	return roys, nil
}
