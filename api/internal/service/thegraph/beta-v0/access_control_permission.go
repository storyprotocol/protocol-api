package beta_v0

import (
	"context"
	"fmt"
	beta_v0 "github.com/storyprotocol/protocol-api/api/internal/models/beta-v0"
	"github.com/storyprotocol/protocol-api/api/internal/service/thegraph"
)

func (c *ServiceBetaImpl) ListAccessControlPermissions(options *thegraph.TheGraphQueryOptions) ([]*beta_v0.AccessControlPermission, error) {
	whereString := c.buildWhereConditions(options)
	query := fmt.Sprintf(`
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
		return nil, fmt.Errorf("failed to get access control permissions from the graph. error: %v", err)
	}

	acps := []*beta_v0.AccessControlPermission{}
	for _, acp := range acpsRes.AccessControlPermissions {
		acps = append(acps, acp)
	}

	return acps, nil
}
