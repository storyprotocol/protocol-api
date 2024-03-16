package betav0

import (
	"context"
	"fmt"
	"github.com/machinebox/graphql"
	beta_v0 "github.com/storyprotocol/protocol-api/api/internal/models/betav0"
	"github.com/storyprotocol/protocol-api/api/internal/service/thegraph"
)

func (c *ServiceBetaImpl) GetPolicy(policyId string) (*beta_v0.Policy, *beta_v0.RenPolicy, error) {
	if c.apiKey != "" {
		query := fmt.Sprintf(`
		query {
		  record(id: "%s") {
			id
			policy_framework_manager
			framework_data
			royalty_policy
			royalty_data
			minting_fee
			minting_fee_token
			block_time
			block_number
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
		req := c.buildNewRequest(nil, query)
		ctx := context.Background()
		var polRes beta_v0.RenPolicyTheGraphResponse
		if err := c.client.Run(ctx, req, &polRes); err != nil {
			return nil, nil, fmt.Errorf("failed to get policy from the graph. error: %v", err)
		}

		return nil, polRes.Policy, nil
	} else {
		query := fmt.Sprintf(`
		query {
		  policy(id: "%s") {
			id
			policyFrameworkManager
			frameworkData
			royaltyPolicy
			royaltyData
			mintingFee
			mintingFeeToken
			blockTimestamp
			blockNumber
			pil {
			  attribution
			  commercialAttribution
			  commercialRevShare
			  commercialUse
			  commercializerChecker
			  commercializerCheckerData
			  territories
			  id
			  distributionChannels
			  derivativesReciprocal
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
			return nil, nil, fmt.Errorf("failed to get policy from the graph. error: %v", err)
		}

		return polRes.Policy, nil, nil
	}

}

func (c *ServiceBetaImpl) ListPolicies(options *thegraph.TheGraphQueryOptions) ([]*beta_v0.Policy, []*beta_v0.RenPolicy, error) {
	whereString := c.buildWhereConditions(options)
	query := ""

	if c.apiKey != "" {
		VALUES := fmt.Sprintf(REN_QUERY_VALUE, ORDER_BY, ORDER_DIRECTION)

		//TODO: What we need
		//id
		//policy_framework_manager
		//framework_data
		//royalty_policy
		//royalty_data
		//minting_fee
		//minting_fee_token
		//block_timestamp
		//block_number
		//pil{
		//	attribution
		//	commercial_attribution
		//	commercial_rev_share
		//	commercial_use
		//	commercializer_checker
		//	commercializer_checker_data
		//	territories
		//	id
		//	distribution_channels
		//	derivatives_reciprocal
		//	derivatives_attribution
		//	derivatives_approval
		//	derivatives_allowed
		//	content_restrictions
		//}
		query = fmt.Sprintf(`
		query(%s){
		  records(%s, filter:{%s}) {
				id
				policy_framework_manager
				framework_data
				royalty_policy
				royalty_data
				minting_fee
				minting_fee_token
				block_time
				block_number
			  }
			}
		`, REN_QUERY_INTERFACE, VALUES, whereString)
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
		var polsRes beta_v0.RenPoliciesTheGraphResponse
		if err := c.client.Run(ctx, req, &polsRes); err != nil {
			return nil, nil, fmt.Errorf("failed to get policies from the graph. error: %v", err)
		}

		pols := []*beta_v0.RenPolicy{}
		for _, pol := range polsRes.Policies {
			pols = append(pols, pol)
		}

		return nil, pols, nil
	} else {
		query = fmt.Sprintf(`
	query(%s){
	  policies(%s, where:{%s}) {
			id
			policyFrameworkManager
			frameworkData
			royaltyPolicy
			royaltyData
			mintingFee
			mintingFeeToken
			blockTimestamp
			blockNumber
			pil {
			  attribution
			  commercialAttribution
			  commercialRevShare
			  commercialUse
			  commercializerChecker
			  commercializerCheckerData
			  territories
			  id
			  distributionChannels
			  derivativesReciprocal
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
			return nil, nil, fmt.Errorf("failed to get policies from the graph. error: %v", err)
		}

		pols := []*beta_v0.Policy{}
		for _, pol := range polsRes.Policies {
			pols = append(pols, pol)
		}

		return pols, nil, nil
	}

}
