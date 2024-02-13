package betav0

import (
	"context"
	"fmt"
	"github.com/machinebox/graphql"
	beta_v0 "github.com/storyprotocol/protocol-api/api/internal/models/betav0"
	"github.com/storyprotocol/protocol-api/api/internal/service/thegraph"
)

func (c *ServiceBetaImpl) GetPolicy(policyId string) (*beta_v0.Policy, error) {
	query := fmt.Sprintf(`
		query {
		  policy(id: "%s") {
			id
			policyFrameworkManager
			uml
			blockTimestamp
			blockNumber
			uml {
			  attribution
			  commercialAttribution
			  commercialRevShare
			  commercialUse
			  commercializers
			  transferable
			  territories
			  id
			  distributionChannels
			  derivativesRevShare
			  derivativesReciprocal
			  royaltyPolicy
			  derivativesAttribution
			  derivativesApproval
			  derivativesAllowed
			  contentRestrictions
			}
		  }
		}
    `, policyId)
	//policyData {
	//	id
	//	frameworkId
	//	needsActivation
	//	mintingParamValues
	//	linkParentParamValues
	//	activationParamValues
	//}
	req := graphql.NewRequest(query)
	ctx := context.Background()
	var polRes beta_v0.PolicyTheGraphResponse
	if err := c.client.Run(ctx, req, &polRes); err != nil {
		return nil, fmt.Errorf("failed to get policy from the graph. error: %v", err)
	}

	return polRes.Policy, nil

}

func (c *ServiceBetaImpl) ListPolicies(options *thegraph.TheGraphQueryOptions) ([]*beta_v0.Policy, error) {
	whereString := c.buildWhereConditions(options)
	query := fmt.Sprintf(`
	query(%s){
	  policies(%s, where:{%s}) {
			id
			policyFrameworkManager
			blockTimestamp
			blockNumber
			
			uml {
			  attribution
			  commercialAttribution
			  commercialRevShare
			  commercialUse
			  commercializers
			  transferable
			  territories
			  id
			  distributionChannels
			  derivativesRevShare
			  derivativesReciprocal
			  royaltyPolicy
			  derivativesAttribution
			  derivativesApproval
			  derivativesAllowed
			  contentRestrictions
			}
		  }
	}
    `, QUERY_INTERFACE, QUERY_VALUE, whereString)
	//policyData {
	//	id
	//	frameworkId
	//	needsActivation
	//	mintingParamValues
	//	linkParentParamValues
	//	activationParamValues
	//}
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