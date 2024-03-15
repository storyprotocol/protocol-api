package betav0

import (
	"context"
	"fmt"
	"github.com/machinebox/graphql"
	beta_v0 "github.com/storyprotocol/protocol-api/api/internal/models/betav0"
	"github.com/storyprotocol/protocol-api/api/internal/service/thegraph"
)

func (c *ServiceBetaImpl) GetDispute(disputeId string) (*beta_v0.Dispute, *beta_v0.RenDispute, error) {
	if c.apiKey != "" {
		query := fmt.Sprintf(`
	query {
		record(id: "%s") {
			id
			target_ip_id
			target_tag
			current_tag
			deleted_at
			arbitration_policy
			evidence_link
			initiator
			data
			block_number
			block_time
	  	}
	}
    `, disputeId)

		req := c.buildNewRequest(nil, query)
		ctx := context.Background()
		var disputesRes beta_v0.RenDisputeTheGraphResponse
		if err := c.client.Run(ctx, req, &disputesRes); err != nil {
			return nil, nil, fmt.Errorf("failed to get dispute from the graph. error: %v", err)
		}

		return nil, disputesRes.Dispute, nil
	} else {
		query := fmt.Sprintf(`
	query {
		dispute(id: "%s") {
			id
			targetIpId
			targetTag
			currentTag
			deletedAt
			arbitrationPolicy
			evidenceLink
			initiator
			data
			blockNumber
			blockTimestamp
	  	}
	}
    `, disputeId)

		req := graphql.NewRequest(query)
		ctx := context.Background()
		var disputesRes beta_v0.DisputeTheGraphResponse
		if err := c.client.Run(ctx, req, &disputesRes); err != nil {
			return nil, nil, fmt.Errorf("failed to get dispute from the graph. error: %v", err)
		}

		return disputesRes.Dispute, nil, nil
	}

}

func (c *ServiceBetaImpl) ListDisputes(options *thegraph.TheGraphQueryOptions) ([]*beta_v0.Dispute, []*beta_v0.RenDispute, error) {
	whereString := c.buildWhereConditions(options)
	query := ""

	if c.apiKey != "" {
		VALUES := fmt.Sprintf(REN_QUERY_VALUE, ORDER_BY, ORDER_DIRECTION)

		query = fmt.Sprintf(`
		query(%s){
			records (%s, filter:{%s}) {
				id
				target_ip_id
				target_tag
				current_tag
				deleted_at
				arbitration_policy
				evidence_link
				initiator
				data
				block_number
				block_time
			}
		}
		`, REN_QUERY_INTERFACE, VALUES, whereString)

		req := c.buildNewRequest(options, query)

		ctx := context.Background()
		var disputesRes beta_v0.RenDisputesTheGraphResponse
		if err := c.client.Run(ctx, req, &disputesRes); err != nil {
			return nil, nil, fmt.Errorf("failed to get disputes from the graph. error: %v", err)
		}

		disputes := []*beta_v0.RenDispute{}
		for _, dispute := range disputesRes.Disputes {
			disputes = append(disputes, dispute)
		}

		return nil, disputes, nil
	} else {
		query = fmt.Sprintf(`
	query(%s){
		disputes (%s, where:{%s}) {
			id
			targetIpId
			targetTag
			currentTag
			deletedAt
			arbitrationPolicy
			evidenceLink
			initiator
			data
			blockNumber
			blockTimestamp
		}
	}
    `, QUERY_INTERFACE, QUERY_VALUE, whereString)

		req := c.buildNewRequest(options, query)

		ctx := context.Background()
		var disputesRes beta_v0.DisputesTheGraphResponse
		if err := c.client.Run(ctx, req, &disputesRes); err != nil {
			return nil, nil, fmt.Errorf("failed to get disputes from the graph. error: %v", err)
		}

		disputes := []*beta_v0.Dispute{}
		for _, dispute := range disputesRes.Disputes {
			disputes = append(disputes, dispute)
		}

		return disputes, nil, nil

	}

}
