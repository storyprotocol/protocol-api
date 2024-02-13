package betav0

import (
	"context"
	"fmt"
	"github.com/machinebox/graphql"
	beta_v0 "github.com/storyprotocol/protocol-api/api/internal/models/betav0"
	"github.com/storyprotocol/protocol-api/api/internal/service/thegraph"
)

func (c *ServiceBetaImpl) GetPolicyFrameworkManager(pfmId string) (*beta_v0.PolicyFrameworkManager, error) {
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
		return nil, fmt.Errorf("failed to get policy framework manager from the graph. error: %v", err)
	}

	return pfmIdRes.PolicyFrameworkManager, nil

}

func (c *ServiceBetaImpl) ListPolicyFrameworkManagers(options *thegraph.TheGraphQueryOptions) ([]*beta_v0.PolicyFrameworkManager, error) {
	whereString := c.buildWhereConditions(options)
	query := fmt.Sprintf(`
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
		return nil, fmt.Errorf("failed to get framework managers from the graph. error: %v", err)
	}

	pfwms := []*beta_v0.PolicyFrameworkManager{}
	for _, roy := range pfwmsRes.PolicyFrameworkManagers {
		pfwms = append(pfwms, roy)
	}

	return pfwms, nil
}
