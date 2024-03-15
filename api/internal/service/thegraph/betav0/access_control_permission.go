package betav0

import (
	"context"
	"fmt"
	beta_v0 "github.com/storyprotocol/protocol-api/api/internal/models/betav0"
	"github.com/storyprotocol/protocol-api/api/internal/service/thegraph"
)

func (c *ServiceBetaImpl) ListAccessControlPermissions(options *thegraph.TheGraphQueryOptions) ([]*beta_v0.AccessControlPermission, []*beta_v0.RenAccessControlPermission, error) {
	whereString := c.buildWhereConditions(options)
	query := ""

	if c.apiKey != "" {
		VALUES := fmt.Sprintf(REN_QUERY_VALUE, ORDER_BY, ORDER_DIRECTION)
		query = fmt.Sprintf(`
		query(%s){
			records (%s, filter:{%s}) {
				id
				name
				module
				block_number
				blockTimestamp
			}
		}
		`, REN_QUERY_INTERFACE, VALUES, whereString)

		req := c.buildNewRequest(options, query)

		ctx := context.Background()
		var acpsRes beta_v0.RenAccessControlPermissionTheGraphResponse
		if err := c.client.Run(ctx, req, &acpsRes); err != nil {
			return nil, nil, fmt.Errorf("failed to get access control permissions from the graph. error: %v", err)
		}

		acps := []*beta_v0.RenAccessControlPermission{}
		for _, acp := range acpsRes.AccessControlPermissions {
			acps = append(acps, acp)
		}

		return nil, acps, nil
	} else {
		query = fmt.Sprintf(`
		query(%s){
			modules (%s, where:{%s}) {
				id
				name
				module
				blockNumber
				blockTimestamp
			}
		}
		`, QUERY_INTERFACE, QUERY_VALUE, whereString)

		req := c.buildNewRequest(options, query)

		ctx := context.Background()
		var acpsRes beta_v0.AccessControlPermissionTheGraphResponse
		if err := c.client.Run(ctx, req, &acpsRes); err != nil {
			return nil, nil, fmt.Errorf("failed to get access control permissions from the graph. error: %v", err)
		}

		acps := []*beta_v0.AccessControlPermission{}
		for _, acp := range acpsRes.AccessControlPermissions {
			acps = append(acps, acp)
		}

		return acps, nil, nil
	}

}
