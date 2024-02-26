package betav0

import (
	"context"
	"fmt"
	"github.com/machinebox/graphql"
	beta_v0 "github.com/storyprotocol/protocol-api/api/internal/models/betav0"
	"github.com/storyprotocol/protocol-api/api/internal/service/thegraph"
)

func (c *ServiceBetaImpl) GetModule(moduleId string) (*beta_v0.Module, error) {
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
		return nil, fmt.Errorf("failed to get modules from the graph. error: %v", err)
	}

	return modules.Module, nil

}

func (c *ServiceBetaImpl) ListModules(options *thegraph.TheGraphQueryOptions) ([]*beta_v0.Module, error) {
	whereString := c.buildWhereConditions(options)
	query := fmt.Sprintf(`
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
		return nil, fmt.Errorf("failed to get modules from the graph. error: %v", err)
	}

	mods := []*beta_v0.Module{}
	for _, mod := range modules.Modules {
		mods = append(mods, mod)
	}

	return mods, nil
}
