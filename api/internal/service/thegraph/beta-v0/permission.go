package beta_v0

import (
	"context"
	"fmt"

	"github.com/machinebox/graphql"
	encoder "github.com/storyprotocol/protocol-api/api/internal/helpers"
	beta_v0 "github.com/storyprotocol/protocol-api/api/internal/models/beta-v0"
	"github.com/storyprotocol/protocol-api/api/internal/service/thegraph"
)

func (c *ServiceBetaImpl) GetPermission(permissionId string) (*beta_v0.Permission, error) {
	id, err := encoder.Decrypt(permissionId)
	if err != nil {
		return nil, fmt.Errorf("failed to decode id. error: %v", err)
	}

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
    `, id)

	req := graphql.NewRequest(query)
	ctx := context.Background()
	var permRes beta_v0.PermissionTheGraphResponse
	if err := c.client.Run(ctx, req, &permRes); err != nil {
		return nil, fmt.Errorf("failed to get perm from the graph. error: %v", err)
	}

	permRes.Permission.ID = permissionId

	return permRes.Permission, nil
}

func (c *ServiceBetaImpl) ListPermissions(options *thegraph.TheGraphQueryOptions) ([]*beta_v0.Permission, error) {
	whereString := c.buildWhereConditions(options)
	query := fmt.Sprintf(`
	query(%s) {
	  permissions(id: "%s", where:{%s}) {
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
    `, QUERY_INTERFACE, QUERY_VALUE, whereString)

	req := c.buildNewRequest(options, query)
	ctx := context.Background()
	var permsRes beta_v0.PermissionsTheGraphResponse
	if err := c.client.Run(ctx, req, &permsRes); err != nil {
		return nil, fmt.Errorf("failed to get permissions from the graph. error: %v", err)
	}

	perms := []*beta_v0.Permission{}
	for _, perm := range permsRes.Permissions {
		id, err := encoder.Encrypt(perm.ID)
		if err != nil {
			return nil, fmt.Errorf("failed to encode id. error: %v", err)
		}

		perm.ID = id
		perms = append(perms, perm)
	}

	return perms, nil
}
