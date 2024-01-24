package beta_v0

import (
	"context"
	"fmt"
	"github.com/machinebox/graphql"
	beta_v0 "github.com/storyprotocol/protocol-api/api/internal/models/beta-v0"
	"github.com/storyprotocol/protocol-api/api/internal/service/thegraph"
)

func (c *ServiceBetaImpl) GetPolicy(policyId string) ([]*beta_v0.Policy, error) {
	query := fmt.Sprintf(`
		query {
		  policy(id: "%s") {
			policyId
			creator
			policyData {     
				id
				frameworkId
				needsActivation
				mintingParamValues
				linkParentParamValues
				activationParamValues
			}
		  }
		}
    `, policyId)

	req := graphql.NewRequest(query)
	ctx := context.Background()
	var polRes beta_v0.PolicyTheGraphResponse
	if err := c.client.Run(ctx, req, &polRes); err != nil {
		return nil, fmt.Errorf("failed to get policy from the graph. error: %v", err)
	}

	pols := []*beta_v0.Policy{}
	for _, pol := range polRes.Policy {
		pols = append(pols, pol)
	}

	return pols, nil

}

func (c *ServiceBetaImpl) ListPolicies(options *thegraph.TheGraphQueryOptions) ([]*beta_v0.Policy, error) {
	query := fmt.Sprintf(`
	query(%s){
		{
		  policies(%s) {
			creator
			policyId
			policyData {
			  id
			  frameworkId
			  needsActivation
			  mintingParamValues
			  linkParentParamValues
			  activationParamValues
			}
		  }
		}
	}
    `, QUERY_INTERFACE, QUERY_VALUE)

	req := c.buildNewRequest(options, query)

	ctx := context.Background()
	var polsRes beta_v0.PoliciesTheGraphResponse
	if err := c.client.Run(ctx, req, &polsRes); err != nil {
		return nil, fmt.Errorf("failed to get policies from the graph. error: %v", err)
	}

	pols := []*beta_v0.Policy{}
	for _, pol := range polsRes.Policies {
		pols = append(pols, pol)
	}

	return pols, nil
}
