package beta_v0

import (
	"context"
	"fmt"
	"github.com/machinebox/graphql"
	beta_v0 "github.com/storyprotocol/protocol-api/api/internal/models/beta-v0"
	"github.com/storyprotocol/protocol-api/api/internal/service/thegraph"
)

func (c *ServiceBetaImpl) GetDispute(disputeId string) ([]*beta_v0.Dispute, error) {
	query := fmt.Sprintf(`
	query {
		dispute(id: "%s") {
			name
			module
	  	}
	}
    `, disputeId)

	req := graphql.NewRequest(query)
	ctx := context.Background()
	var disputesRes beta_v0.DisputeTheGraphResponse
	if err := c.client.Run(ctx, req, &disputesRes); err != nil {
		return nil, fmt.Errorf("failed to get dispute from the graph. error: %v", err)
	}

	disputes := []*beta_v0.Dispute{}
	for _, dispute := range disputesRes.Dispute {
		disputes = append(disputes, dispute)
	}

	return disputes, nil

}

func (c *ServiceBetaImpl) ListDisputes(options *thegraph.TheGraphQueryOptions) ([]*beta_v0.Dispute, error) {
	query := fmt.Sprintf(`
	query(%s){
		disputes (%s) {
			name
			module
		}
	}
    `, QUERY_INTERFACE, QUERY_VALUE)

	req := c.buildNewRequest(options, query)

	ctx := context.Background()
	var disputesRes beta_v0.DisputesTheGraphResponse
	if err := c.client.Run(ctx, req, &disputesRes); err != nil {
		return nil, fmt.Errorf("failed to get disputes from the graph. error: %v", err)
	}

	disputes := []*beta_v0.Dispute{}
	for _, dispute := range disputesRes.Disputes {
		disputes = append(disputes, dispute)
	}

	return disputes, nil
}
