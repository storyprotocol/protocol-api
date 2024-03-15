package betav0

import (
	"context"
	"fmt"
	"github.com/machinebox/graphql"
	beta_v0 "github.com/storyprotocol/protocol-api/api/internal/models/betav0"
	"github.com/storyprotocol/protocol-api/api/internal/service/thegraph"
)

func (c *ServiceBetaImpl) GetRoyalty(royaltyId string) (*beta_v0.Royalty, *beta_v0.RenRoyalty, error) {
	if c.apiKey == "" {
		query := fmt.Sprintf(`
		query {
		  record(id: "%s") {
			id
			ip_id
			data
			royalty_policy
			block_timestamp
			block_number
		  }
		}
    `, royaltyId)

		req := graphql.NewRequest(query)
		ctx := context.Background()
		var royRes beta_v0.RenRoyaltyTheGraphResponse
		if err := c.client.Run(ctx, req, &royRes); err != nil {
			return nil, nil, fmt.Errorf("failed to get royalty from the graph. error: %v", err)
		}

		return nil, royRes.Royalty, nil
	} else {
		query := fmt.Sprintf(`
		query {
		  iproyalty(id: "%s") {
			id
			ipId
			data
			royaltyPolicy
			blockTimestamp
			blockNumber
		  }
		}
    `, royaltyId)

		req := graphql.NewRequest(query)
		ctx := context.Background()
		var royRes beta_v0.RoyaltyTheGraphResponse
		if err := c.client.Run(ctx, req, &royRes); err != nil {
			return nil, nil, fmt.Errorf("failed to get royalty from the graph. error: %v", err)
		}

		return royRes.Royalty, nil, nil
	}

}

func (c *ServiceBetaImpl) ListRoyalties(options *thegraph.TheGraphQueryOptions) ([]*beta_v0.Royalty, []*beta_v0.RenRoyalty, error) {
	whereString := c.buildWhereConditions(options)
	query := ""

	if c.apiKey != "" {
		VALUES := fmt.Sprintf(REN_QUERY_VALUE, ORDER_BY, ORDER_DIRECTION)

		query = fmt.Sprintf(`
		query(%s){
		  records(%s, filter:{%s}) {
				id
				royalty_policy
				ip_id
				data
				block_timestamp
				block_number
		  }
			
		}
		`, REN_QUERY_INTERFACE, VALUES, whereString)

		req := c.buildNewRequest(options, query)

		ctx := context.Background()
		var roysRes beta_v0.RenRoyaltiesTheGraphResponse
		if err := c.client.Run(ctx, req, &roysRes); err != nil {
			return nil, nil, fmt.Errorf("failed to get royalties from the graph. error: %v", err)
		}

		roys := []*beta_v0.RenRoyalty{}
		for _, roy := range roysRes.Royalties {
			roys = append(roys, roy)
		}

		return nil, roys, nil
	} else {
		query = fmt.Sprintf(`
	query(%s){
	  iproyalties(%s, where:{%s}) {
			id
			royaltyPolicy
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
			return nil, nil, fmt.Errorf("failed to get royalties from the graph. error: %v", err)
		}

		roys := []*beta_v0.Royalty{}
		for _, roy := range roysRes.Royalties {
			roys = append(roys, roy)
		}

		return roys, nil, nil
	}

}
