package betav0

import (
	"context"
	"fmt"
	beta_v0 "github.com/storyprotocol/protocol-api/api/internal/models/betav0"
	"github.com/storyprotocol/protocol-api/api/internal/service/thegraph"
)

func (c *ServiceBetaImpl) GetPermission(permissionId string) (*beta_v0.Permission, *beta_v0.RenPermission, error) {
	if c.apiKey != "" {
		query := fmt.Sprintf(`
		query {
		  records(filter: { uuid:  {eq: "%s"} }) {
			id
			uuid
			permission
			signer
			to_address
			func
			block_time
			block_number
		  }
		}
    `, permissionId)

		req := c.buildNewRequest(nil, query)
		ctx := context.Background()
		var permsRes beta_v0.RenPermissionsTheGraphResponse
		if err := c.client.Run(ctx, req, &permsRes); err != nil {
			return nil, nil, fmt.Errorf("failed to get permissions from the graph. error: %v", err)
		}

		perms := []*beta_v0.RenPermission{}
		for _, perm := range permsRes.Permissions {
			name, err := c.openChainClient.GetPermissionName(perm.Func)
			if err != nil {
				name = perm.Func
			}
			perm.Permission = name
			perm.ID = perm.UUID
			perms = append(perms, perm)
			perm.UUID = ""
			perm.To = perm.ToAddress
			perm.ToAddress = ""
		}
		return nil, perms[0], nil
	} else {
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
			return nil, nil, fmt.Errorf("failed to get permissions from the graph. error: %v", err)
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
		return perms[0], nil, nil
	}
}

func (c *ServiceBetaImpl) ListPermissions(options *thegraph.TheGraphQueryOptions) ([]*beta_v0.Permission, []*beta_v0.RenPermission, error) {
	whereString := c.buildWhereConditions(options)
	query := ""

	if c.apiKey != "" {
		VALUES := fmt.Sprintf(REN_QUERY_VALUE, ORDER_BY, ORDER_DIRECTION)

		query = fmt.Sprintf(`
		query(%s) {
		  records(%s, filter:{%s}) {
			id
			uuid
			permission
			signer
			to_address
			func
			block_time
			block_number
		  }
		}
		`, REN_QUERY_INTERFACE, VALUES, whereString)

		req := c.buildNewRequest(options, query)

		ctx := context.Background()
		var permsRes beta_v0.RenPermissionsTheGraphResponse
		if err := c.client.Run(ctx, req, &permsRes); err != nil {
			return nil, nil, fmt.Errorf("failed to get permissions from the graph. error: %v", err)
		}

		perms := []*beta_v0.RenPermission{}
		for _, perm := range permsRes.Permissions {
			name, err := c.openChainClient.GetPermissionName(perm.Func)
			if err != nil {
				name = perm.Func
			}
			perm.Permission = name
			perm.ID = perm.UUID
			perms = append(perms, perm)
			perm.UUID = ""
			perm.To = perm.ToAddress
			perm.ToAddress = ""
		}

		return nil, perms, nil
	} else {
		query = fmt.Sprintf(`
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
			return nil, nil, fmt.Errorf("failed to get permissions from the graph. error: %v", err)
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

		return perms, nil, nil
	}

}
