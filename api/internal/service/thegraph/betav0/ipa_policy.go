package betav0

import (
	"context"
	"fmt"
	"github.com/machinebox/graphql"
	beta_v0 "github.com/storyprotocol/protocol-api/api/internal/models/betav0"
	"github.com/storyprotocol/protocol-api/api/internal/service/thegraph"
)

func (c *ServiceBetaImpl) GetIPAPolicy(ipaPolicyId string) (*beta_v0.IPAPolicy, *beta_v0.RenIPAPolicy, error) {

	if c.apiKey != "" {
		query := fmt.Sprintf(`
	query {
		record(id: "%s") {
			id
			ip_id
			policy_id
			index
			active
			inherited
			block_number
			block_time
	  	}
	}
    `, ipaPolicyId)

		req := c.buildNewRequest(nil, query)
		ctx := context.Background()
		var ipapRes beta_v0.RenIPAPolicyTheGraphResponse
		if err := c.client.Run(ctx, req, &ipapRes); err != nil {
			return nil, nil, fmt.Errorf("failed to get ipapolicu from the graph. error: %v", err)
		}

		return nil, ipapRes.IPAPolicy, nil
	} else {
		query := fmt.Sprintf(`
	query {
		ipapolicy(id: "%s") {
			id
			ipId
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
			return nil, nil, fmt.Errorf("failed to get ipapolicu from the graph. error: %v", err)
		}

		return ipapRes.IPAPolicy, nil, nil
	}

}

func (c *ServiceBetaImpl) ListIPAPolicies(options *thegraph.TheGraphQueryOptions) ([]*beta_v0.IPAPolicy, []*beta_v0.RenIPAPolicy, error) {
	whereString := c.buildWhereConditions(options)
	query := ""

	if c.apiKey != "" {
		VALUES := fmt.Sprintf(REN_QUERY_VALUE, ORDER_BY, ORDER_DIRECTION)

		query = fmt.Sprintf(`
		query(%s){
			records (%s, filter:{%s}) {
				id
				ip_id
				policy_id
				index
				active
				inherited
				block_number
				block_time
			}
		}
		`, REN_QUERY_INTERFACE, VALUES, whereString)

		req := c.buildNewRequest(options, query)

		ctx := context.Background()
		var ipapRes beta_v0.RenIPAPoliciesTheGraphResponse
		if err := c.client.Run(ctx, req, &ipapRes); err != nil {
			return nil, nil, fmt.Errorf("failed to get ipapolicies from the graph. error: %v", err)
		}

		ipaps := []*beta_v0.RenIPAPolicy{}
		for _, ipap := range ipapRes.IPAPolicies {
			ipaps = append(ipaps, ipap)
		}

		return nil, ipaps, nil
	} else {
		query = fmt.Sprintf(`
	query(%s){
		ipapolicies (%s, where:{%s}) {
			id
			ipId
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
			return nil, nil, fmt.Errorf("failed to get ipapolicies from the graph. error: %v", err)
		}

		ipaps := []*beta_v0.IPAPolicy{}
		for _, ipap := range ipapRes.IPAPolicies {
			ipaps = append(ipaps, ipap)
		}

		return ipaps, nil, nil
	}

}
