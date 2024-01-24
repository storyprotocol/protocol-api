package beta_v0

import (
	"context"
	"fmt"
	"github.com/machinebox/graphql"
	beta_v0 "github.com/storyprotocol/protocol-api/api/internal/models/beta-v0"
	"github.com/storyprotocol/protocol-api/api/internal/service/thegraph"
)

func (c *ServiceBetaImpl) GetModule(moduleName string) ([]*beta_v0.Module, error) {
	query := fmt.Sprintf(`
	query {
		module(id: "%s") {
			name
			module
	  	}
	}
    `, moduleName)

	req := graphql.NewRequest(query)
	ctx := context.Background()
	var modules beta_v0.ModuleTheGraphResponse
	if err := c.client.Run(ctx, req, &modules); err != nil {
		return nil, fmt.Errorf("failed to get modules from the graph. error: %v", err)
	}

	mods := []*beta_v0.Module{}
	for _, mod := range modules.Module {
		mods = append(mods, mod)
	}

	return mods, nil

}

func (c *ServiceBetaImpl) ListModules(options *thegraph.TheGraphQueryOptions) ([]*beta_v0.Module, error) {
	query := fmt.Sprintf(`
	query(%s){
		modules (%s) {
			name
			module
		}
	}
    `, QUERY_INTERFACE, QUERY_VALUE)

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
