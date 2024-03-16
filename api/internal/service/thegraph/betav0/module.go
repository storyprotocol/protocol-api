package betav0

import (
	"context"
	"fmt"
	"github.com/machinebox/graphql"
	beta_v0 "github.com/storyprotocol/protocol-api/api/internal/models/betav0"
	"github.com/storyprotocol/protocol-api/api/internal/service/thegraph"
)

func (c *ServiceBetaImpl) GetModule(moduleId string) (*beta_v0.Module, *beta_v0.RenModule, error) {
	if c.apiKey != "" {
		query := fmt.Sprintf(`
		query {
			record(id: "%s") {
				id
				name
				module
				block_number
				block_time
				deleted_at
			}
		}
		`, moduleId)

		req := c.buildNewRequest(nil, query)
		ctx := context.Background()
		var modules beta_v0.RenModuleTheGraphResponse
		if err := c.client.Run(ctx, req, &modules); err != nil {
			return nil, nil, fmt.Errorf("failed to get modules from the graph. error: %v", err)
		}

		return nil, modules.Module, nil

	} else {
		query := fmt.Sprintf(`
	query {
		module(id: "%s") {
			id
			name
			module
			blockNumber
			blockTimestamp
			deletedAt
	  	}
	}
    `, moduleId)

		req := graphql.NewRequest(query)
		ctx := context.Background()
		var modules beta_v0.ModuleTheGraphResponse
		if err := c.client.Run(ctx, req, &modules); err != nil {
			return nil, nil, fmt.Errorf("failed to get modules from the graph. error: %v", err)
		}

		return modules.Module, nil, nil

	}
}

func (c *ServiceBetaImpl) ListModules(options *thegraph.TheGraphQueryOptions) ([]*beta_v0.Module, []*beta_v0.RenModule, error) {
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
				block_time
				deleted_at
			}
		}
		`, REN_QUERY_INTERFACE, VALUES, whereString)

		req := c.buildNewRequest(options, query)

		ctx := context.Background()
		var modules beta_v0.RenModulesTheGraphResponse
		if err := c.client.Run(ctx, req, &modules); err != nil {
			return nil, nil, fmt.Errorf("failed to get modules from the graph. error: %v", err)
		}

		mods := []*beta_v0.RenModule{}
		for _, mod := range modules.Modules {
			mods = append(mods, mod)
		}

		return nil, mods, nil
	} else {
		query = fmt.Sprintf(`
	query(%s){
		modules (%s, where:{%s}) {
			id
			name
			module
			blockNumber
			blockTimestamp
			deletedAt
		}
	}
    `, QUERY_INTERFACE, QUERY_VALUE, whereString)

		req := c.buildNewRequest(options, query)

		ctx := context.Background()
		var modules beta_v0.ModulesTheGraphResponse
		if err := c.client.Run(ctx, req, &modules); err != nil {
			return nil, nil, fmt.Errorf("failed to get modules from the graph. error: %v", err)
		}

		mods := []*beta_v0.Module{}
		for _, mod := range modules.Modules {
			mods = append(mods, mod)
		}

		return mods, nil, nil
	}

}
