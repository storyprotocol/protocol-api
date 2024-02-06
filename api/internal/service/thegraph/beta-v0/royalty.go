package beta_v0

import (
	"context"
	"fmt"
	"github.com/machinebox/graphql"
	beta_v0 "github.com/storyprotocol/protocol-api/api/internal/models/beta-v0"
	"github.com/storyprotocol/protocol-api/api/internal/service/thegraph"
)

func (c *ServiceBetaImpl) GetRoyalty(royaltyId string) (*beta_v0.Royalty, error) {
	query := fmt.Sprintf(`
		query {
		  royalty(id: "%s") {
			ipId
			data
			blockTimestamp
			blockNumber
		  }
		}
    `, royaltyId)

	req := graphql.NewRequest(query)
	ctx := context.Background()
	var royRes beta_v0.RoyaltyTheGraphResponse
	if err := c.client.Run(ctx, req, &royRes); err != nil {
		return nil, fmt.Errorf("failed to get royalty from the graph. error: %v", err)
	}

	return royRes.Royalty, nil

}

func (c *ServiceBetaImpl) ListRoyalties(options *thegraph.TheGraphQueryOptions) ([]*beta_v0.Royalty, error) {
	whereString := c.buildWhereConditions(options)
	query := fmt.Sprintf(`
	query(%s){
	  royalties(%s, where:{%s}) {
			ipId
			data
			blockTimestamp
			blockNumber
	  }
		
	}
    `, QUERY_INTERFACE, QUERY_VALUE, whereString)

	req := c.buildNewRequest(options, query)

	ctx := context.Background()
	var roysRes beta_v0.RoyaltiesTheGraphResponse
	if err := c.client.Run(ctx, req, &roysRes); err != nil {
		return nil, fmt.Errorf("failed to get royalties from the graph. error: %v", err)
	}

	roys := []*beta_v0.Royalty{}
	for _, roy := range roysRes.Royalties {
		roys = append(roys, roy)
	}

	return roys, nil
}
