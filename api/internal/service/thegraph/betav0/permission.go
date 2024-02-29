package betav0

import (
	"context"
	"fmt"
	beta_v0 "github.com/storyprotocol/protocol-api/api/internal/models/betav0"
	"github.com/storyprotocol/protocol-api/api/internal/service/thegraph"
)

func (c *ServiceBetaImpl) GetPermission(permissionId string) (*beta_v0.Permission, error) {
	query := fmt.Sprintf(`
		query {
		  permissions(where: { uuid:  "%s" }) {
			id
			uuid
			permission
			signer
			to
			func
			blockTimestamp
			blockNumber
		  }
		}
    `, permissionId)

	req := c.buildNewRequest(nil, query)
	ctx := context.Background()
	var permsRes beta_v0.PermissionsTheGraphResponse
	if err := c.client.Run(ctx, req, &permsRes); err != nil {
		return nil, fmt.Errorf("failed to get permissions from the graph. error: %v", err)
	}

	perms := []*beta_v0.Permission{}
	for _, perm := range permsRes.Permissions {
		name, err := c.openChainClient.GetPermissionName(perm.Func)
		if err != nil {
			name = perm.Func
		}
		perm.Permission = name
		perm.ID = perm.UUID
		perms = append(perms, perm)
		perm.UUID = ""
	}
	return perms[0], nil
}

func (c *ServiceBetaImpl) ListPermissions(options *thegraph.TheGraphQueryOptions) ([]*beta_v0.Permission, error) {
	whereString := c.buildWhereConditions(options)
	//TODO: FIgure out why id: in the query statement causes ordering by not to work
	query := fmt.Sprintf(`
	query(%s) {
	  permissions(id: "%s", where:{%s}) {
		id
		uuid
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
		name, err := c.openChainClient.GetPermissionName(perm.Func)
		if err != nil {
			name = perm.Func
		}
		perm.Permission = name
		perm.ID = perm.UUID
		perms = append(perms, perm)
		perm.UUID = ""
	}

	return perms, nil
}
