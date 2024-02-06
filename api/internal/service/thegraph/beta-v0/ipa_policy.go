package beta_v0

import (
	"context"
	"fmt"
	"github.com/machinebox/graphql"
	beta_v0 "github.com/storyprotocol/protocol-api/api/internal/models/beta-v0"
	"github.com/storyprotocol/protocol-api/api/internal/service/thegraph"
)

func (c *ServiceBetaImpl) GetIPAPolicy(ipaPolicyId string) (*beta_v0.IPAPolicy, error) {
	query := fmt.Sprintf(`
	query {
		ipapolicy(id: "%s") {
			id
			policyId
			index
			active
			inherited
			blockNumber
			blockTimestamp
	  	}
	}
    `, ipaPolicyId)

	req := graphql.NewRequest(query)
	ctx := context.Background()
	var ipapRes beta_v0.IPAPolicyTheGraphResponse
	if err := c.client.Run(ctx, req, &ipapRes); err != nil {
		return nil, fmt.Errorf("failed to get ipapolicu from the graph. error: %v", err)
	}

	return ipapRes.IPAPolicy, nil

}

func (c *ServiceBetaImpl) ListIPAPolicies(options *thegraph.TheGraphQueryOptions) ([]*beta_v0.IPAPolicy, error) {
	whereString := c.buildWhereConditions(options)
	query := fmt.Sprintf(`
	query(%s){
		ipapolicies (%s, where:{%s}) {
			id
			policyId
			index
			active
			inherited
			blockNumber
			blockTimestamp
		}
	}
    `, QUERY_INTERFACE, QUERY_VALUE, whereString)

	req := c.buildNewRequest(options, query)

	ctx := context.Background()
	var ipapRes beta_v0.IPAPoliciesTheGraphResponse
	if err := c.client.Run(ctx, req, &ipapRes); err != nil {
		return nil, fmt.Errorf("failed to get ipapolicies from the graph. error: %v", err)
	}

	ipaps := []*beta_v0.IPAPolicy{}
	for _, ipap := range ipapRes.IPAPolicies {
		ipaps = append(ipaps, ipap)
	}

	return ipaps, nil
}
