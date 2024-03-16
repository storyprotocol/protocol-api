package betav0

import (
	"context"
	"fmt"
	"github.com/machinebox/graphql"
	beta_v0 "github.com/storyprotocol/protocol-api/api/internal/models/betav0"
	"github.com/storyprotocol/protocol-api/api/internal/service/thegraph"
)

func (c *ServiceBetaImpl) GetPolicyFrameworkManager(pfmId string) (*beta_v0.PolicyFrameworkManager, *beta_v0.RenPolicyFrameworkManager, error) {
	if c.apiKey != "" {
		query := fmt.Sprintf(`
		query {
		  record(id: "%s") {
			id
			address
			name
			license_text_url
			block_time
			block_number
			}
		}
    `, pfmId)

		req := c.buildNewRequest(nil, query)
		ctx := context.Background()
		var pfmIdRes beta_v0.RenPolicyFrameworkManagerTheGraphResponse
		if err := c.client.Run(ctx, req, &pfmIdRes); err != nil {
			return nil, nil, fmt.Errorf("failed to get policy framework manager from the graph. error: %v", err)
		}

		pfmIdRes.PolicyFrameworkManager.LicenseUrl = pfmIdRes.PolicyFrameworkManager.LicenseTextUrl
		pfmIdRes.PolicyFrameworkManager.LicenseTextUrl = ""
		return nil, pfmIdRes.PolicyFrameworkManager, nil
	} else {
		query := fmt.Sprintf(`
		query {
		  policyFrameworkManager(id: "%s") {
			id
			address
			name
			licenseUrl
			blockTimestamp
			blockNumber
		  }
		}
    `, pfmId)

		req := graphql.NewRequest(query)
		ctx := context.Background()
		var pfmIdRes beta_v0.PolicyFrameworkManagerTheGraphResponse
		if err := c.client.Run(ctx, req, &pfmIdRes); err != nil {
			return nil, nil, fmt.Errorf("failed to get policy framework manager from the graph. error: %v", err)
		}

		return pfmIdRes.PolicyFrameworkManager, nil, nil
	}

}

func (c *ServiceBetaImpl) ListPolicyFrameworkManagers(options *thegraph.TheGraphQueryOptions) ([]*beta_v0.PolicyFrameworkManager, []*beta_v0.RenPolicyFrameworkManager, error) {
	whereString := c.buildWhereConditions(options)
	query := ""

	if c.apiKey != "" {
		VALUES := fmt.Sprintf(REN_QUERY_VALUE, ORDER_BY, ORDER_DIRECTION)

		query = fmt.Sprintf(`
			query(%s){
			  records(%s, filter:{%s}) {
					id
					address
					name
					license_text_url
					block_time
					block_number
			  }
				
			}
			`, REN_QUERY_INTERFACE, VALUES, whereString)

		req := c.buildNewRequest(options, query)

		ctx := context.Background()
		var pfwmsRes beta_v0.RenPolicyFrameworkManagersTheGraphResponse
		if err := c.client.Run(ctx, req, &pfwmsRes); err != nil {
			return nil, nil, fmt.Errorf("failed to get framework managers from the graph. error: %v", err)
		}

		pfwms := []*beta_v0.RenPolicyFrameworkManager{}
		for _, roy := range pfwmsRes.PolicyFrameworkManagers {
			roy.LicenseUrl = roy.LicenseTextUrl
			roy.LicenseTextUrl = ""
			pfwms = append(pfwms, roy)
		}

		return nil, pfwms, nil
	} else {
		query = fmt.Sprintf(`
		query(%s){
		  policyFrameworkManagers(%s, where:{%s}) {
				id
				address
				name
				licenseUrl
				blockTimestamp
				blockNumber
		  }
			
		}
		`, QUERY_INTERFACE, QUERY_VALUE, whereString)

		req := c.buildNewRequest(options, query)

		ctx := context.Background()
		var pfwmsRes beta_v0.PolicyFrameworkManagersTheGraphResponse
		if err := c.client.Run(ctx, req, &pfwmsRes); err != nil {
			return nil, nil, fmt.Errorf("failed to get framework managers from the graph. error: %v", err)
		}

		pfwms := []*beta_v0.PolicyFrameworkManager{}
		for _, roy := range pfwmsRes.PolicyFrameworkManagers {
			pfwms = append(pfwms, roy)
		}

		return pfwms, nil, nil
	}

}
