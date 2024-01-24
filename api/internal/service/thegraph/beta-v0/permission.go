package beta_v0

import (
	"context"
	"fmt"
	"github.com/machinebox/graphql"
	beta_v0 "github.com/storyprotocol/protocol-api/api/internal/models/beta-v0"
	"github.com/storyprotocol/protocol-api/api/internal/service/thegraph"
)

func (c *ServiceBetaImpl) GetPermission(permissionId string) (*beta_v0.Permission, error) {
	query := fmt.Sprintf(`
		query {
		  permission(id: "%s") {
			id
			ipAccount
			permission
			signer
			to
			func
			blockTimestamp
			blockNumber
		  }
		}
    `, permissionId)

	req := graphql.NewRequest(query)
	ctx := context.Background()
	var permRes beta_v0.PermissionTheGraphResponse
	if err := c.client.Run(ctx, req, &permRes); err != nil {
		return nil, fmt.Errorf("failed to get perm from the graph. error: %v", err)
	}

	return permRes.Permission, nil
}

func (c *ServiceBetaImpl) ListPermissions(options *thegraph.TheGraphQueryOptions) ([]*beta_v0.Permission, error) {
	query := fmt.Sprintf(`
	query(%s) {
	  permission(id: "%s") {
		id
		ipAccount
		permission
		signer
		to
		func
		blockTimestamp
		blockNumber
	  }
	}
    `, QUERY_INTERFACE, QUERY_VALUE)

	req := c.buildNewRequest(options, query)
	ctx := context.Background()
	var permsRes beta_v0.PermissionsTheGraphResponse
	if err := c.client.Run(ctx, req, &permsRes); err != nil {
		return nil, fmt.Errorf("failed to get permissions from the graph. error: %v", err)
	}

	perms := []*beta_v0.Permission{}
	for _, perm := range permsRes.Permissions {
		perms = append(perms, perm)
	}

	return perms, nil
}
