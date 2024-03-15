package betav0

import (
	"context"
	"fmt"
	"github.com/machinebox/graphql"
	beta_v0 "github.com/storyprotocol/protocol-api/api/internal/models/betav0"
	"github.com/storyprotocol/protocol-api/api/internal/service/thegraph"
)

func (c *ServiceBetaImpl) GetRoyaltyPolicy(royaltyPolicyId string) (*beta_v0.RoyaltyPolicy, *beta_v0.RenRoyaltyPolicy, error) {
	if c.apiKey != "" {
		query := fmt.Sprintf(`
		query {
		  record(id: "%s") {
			id
			split_clone
			ancestors_vault
			royalty_stack
			target_ancestors
			target_royalty_amount
			block_number
			block_time
		  }
		}
    `, royaltyPolicyId)

		req := c.buildNewRequest(nil, query)
		ctx := context.Background()
		var royaltyPolRes beta_v0.RenRoyaltyPolicyTheGraphResponse
		if err := c.client.Run(ctx, req, &royaltyPolRes); err != nil {
			return nil, nil, fmt.Errorf("failed to get royalty policy from the graph. error: %v", err)
		}

		return nil, royaltyPolRes.RoyaltyPolicy, nil
	} else {
		query := fmt.Sprintf(`
		query {
		  royaltyPolicy(id: "%s") {
			id
			splitClone
			ancestorsVault
			royaltyStack
			targetAncestors
			targetRoyaltyAmount
			blockNumber
			blockTimestamp
		  }
		}
    `, royaltyPolicyId)

		req := graphql.NewRequest(query)
		ctx := context.Background()
		var royaltyPolRes beta_v0.RoyaltyPolicyTheGraphResponse
		if err := c.client.Run(ctx, req, &royaltyPolRes); err != nil {
			return nil, nil, fmt.Errorf("failed to get royalty policy from the graph. error: %v", err)
		}

		return royaltyPolRes.RoyaltyPolicy, nil, nil
	}

}

func (c *ServiceBetaImpl) ListRoyaltyPolicies(options *thegraph.TheGraphQueryOptions) ([]*beta_v0.RoyaltyPolicy, []*beta_v0.RenRoyaltyPolicy, error) {
	whereString := c.buildWhereConditions(options)
	query := ""

	if c.apiKey != "" {
		VALUES := fmt.Sprintf(REN_QUERY_VALUE, ORDER_BY, ORDER_DIRECTION)

		query = fmt.Sprintf(`
		query(%s){
		  records (%s, filter:{%s}) {
			id
			split_clone
			ancestors_vault
			royalty_stack
			target_ancestors
			target_royalty_amount
			block_number
			block_time
		  }
		}
		`, REN_QUERY_INTERFACE, VALUES, whereString)

		req := c.buildNewRequest(options, query)

		ctx := context.Background()
		var royaltyPoliciesRes beta_v0.RenRoyaltyPoliciesTheGraphResponse
		if err := c.client.Run(ctx, req, &royaltyPoliciesRes); err != nil {
			return nil, nil, fmt.Errorf("failed to get royalty policies from the graph. error: %v", err)
		}

		royaltyPolicies := []*beta_v0.RenRoyaltyPolicy{}
		for _, rPol := range royaltyPoliciesRes.Royalties {
			royaltyPolicies = append(royaltyPolicies, rPol)
		}

		return nil, royaltyPolicies, nil
	} else {
		query = fmt.Sprintf(`
		query(%s){
		  royaltyPolicies (%s, where:{%s}) {
			id
			splitClone
			ancestorsVault
			royaltyStack
			targetAncestors
			targetRoyaltyAmount
			blockNumber
			blockTimestamp
		  }
		}
		`, QUERY_INTERFACE, QUERY_VALUE, whereString)

		req := c.buildNewRequest(options, query)

		ctx := context.Background()
		var royaltyPoliciesRes beta_v0.RoyaltyPoliciesTheGraphResponse
		if err := c.client.Run(ctx, req, &royaltyPoliciesRes); err != nil {
			return nil, nil, fmt.Errorf("failed to get royalty policies from the graph. error: %v", err)
		}

		royaltyPolicies := []*beta_v0.RoyaltyPolicy{}
		for _, rPol := range royaltyPoliciesRes.Royalties {
			royaltyPolicies = append(royaltyPolicies, rPol)
		}

		return royaltyPolicies, nil, nil
	}

}
