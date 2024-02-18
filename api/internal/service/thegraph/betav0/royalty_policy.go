package betav0

import (
	"context"
	"fmt"
	"github.com/machinebox/graphql"
	beta_v0 "github.com/storyprotocol/protocol-api/api/internal/models/betav0"
	"github.com/storyprotocol/protocol-api/api/internal/service/thegraph"
)

func (c *ServiceBetaImpl) GetRoyaltyPolicy(royaltyPolicyId string) (*beta_v0.RoyaltyPolicy, error) {
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
		return nil, fmt.Errorf("failed to get royalty policy from the graph. error: %v", err)
	}

	return royaltyPolRes.RoyaltyPolicy, nil

}

func (c *ServiceBetaImpl) ListRoyaltyPolicies(options *thegraph.TheGraphQueryOptions) ([]*beta_v0.RoyaltyPolicy, error) {
	whereString := c.buildWhereConditions(options)
	query := fmt.Sprintf(`
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

	fmt.Println(query)

	req := c.buildNewRequest(options, query)

	ctx := context.Background()
	var royaltyPoliciesRes beta_v0.RoyaltyPoliciesTheGraphResponse
	if err := c.client.Run(ctx, req, &royaltyPoliciesRes); err != nil {
		return nil, fmt.Errorf("failed to get royalty policies from the graph. error: %v", err)
	}

	royaltyPolicies := []*beta_v0.RoyaltyPolicy{}
	for _, rPol := range royaltyPoliciesRes.Royalties {
		royaltyPolicies = append(royaltyPolicies, rPol)
	}

	return royaltyPolicies, nil
}
