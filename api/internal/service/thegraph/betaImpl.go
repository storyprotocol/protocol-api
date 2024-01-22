package thegraph

import (
	"context"
	"fmt"
	"github.com/machinebox/graphql"
	"github.com/storyprotocol/protocol-api/api/internal/entity"
)

func NewTheGraphServiceBetaImpl(client *graphql.Client) TheGraphServiceBeta {
	return &theGraphServiceBetaImpl{
		client: client,
	}
}

type theGraphServiceBetaImpl struct {
	client *graphql.Client
}

func (c *theGraphServiceBetaImpl) GetIPAccountsRegistered() ([]*entity.IPAccountRegistered, error) {
	req := graphql.NewRequest(`
    {
		ipaccountRegistereds {
			account
			chainId
			implementation
			tokenContract
			tokenId
		}
    }`)

	ctx := context.Background()
	var ipAccountsTheGraphResponse entity.IPAccountsRegisteredTheGraphResponse
	if err := c.client.Run(ctx, req, &ipAccountsTheGraphResponse); err != nil {
		return nil, fmt.Errorf("failed to get registered ip accounts from the graph. error: %v", err)
	}

	ipAccounts := []*entity.IPAccountRegistered{}
	for _, ipAccount := range ipAccountsTheGraphResponse.IPAccountsRegistered {
		ipAccounts = append(ipAccounts, ipAccount)
	}

	return ipAccounts, nil
}

func (c *theGraphServiceBetaImpl) GetIPsRegistered() ([]*entity.IPRegistered, error) {
	req := graphql.NewRequest(`
    {
		ipregistereds {
			id
			chainId
			tokenContract
			tokenId
			resolver
		}
	}`)

	ctx := context.Background()
	var ipRegisteredsTheGraphResponse entity.IPRegisteredTheGraphResponse
	if err := c.client.Run(ctx, req, &ipRegisteredsTheGraphResponse); err != nil {
		return nil, fmt.Errorf("failed to get registered ip accounts from the graph. error: %v", err)
	}

	ips := []*entity.IPRegistered{}
	for _, ip := range ipRegisteredsTheGraphResponse.IPRegistered {
		ips = append(ips, ip)
	}

	return ips, nil
}

func (c *theGraphServiceBetaImpl) GetSetIPAccounts() ([]*entity.SetIPAccount, error) {
	req := graphql.NewRequest(`
    {
		ipaccountSets {
			ipId
			chainId
			tokenContract
			tokenId
		}
	}`)

	ctx := context.Background()
	var ipAccountSets entity.SetIPAccountTheGraphResponse
	if err := c.client.Run(ctx, req, &ipAccountSets); err != nil {
		return nil, fmt.Errorf("failed to get set ip accounts from the graph. error: %v", err)
	}

	accs := []*entity.SetIPAccount{}
	for _, acc := range ipAccountSets.SetIPAccount {
		accs = append(accs, acc)
	}

	return accs, nil
}

func (c *theGraphServiceBetaImpl) GetSetIPResolvers() ([]*entity.SetResolver, error) {
	req := graphql.NewRequest(`
    {
		ipresolverSets {
			ipId
			resolver
		}
	}`)

	ctx := context.Background()
	var ipResolverSets entity.SetResolverTheGraphResponse
	if err := c.client.Run(ctx, req, &ipResolverSets); err != nil {
		return nil, fmt.Errorf("failed to get set ip resolvers from the graph. error: %v", err)
	}

	rslvrs := []*entity.SetResolver{}
	for _, rslvr := range ipResolverSets.SetResolver {
		rslvrs = append(rslvrs, rslvr)
	}

	return rslvrs, nil
}

func (c *theGraphServiceBetaImpl) GetRegisteredModules() ([]*entity.ModuleAdded, error) {
	req := graphql.NewRequest(`
    {
		moduleAddeds {
			name
			module
		}
	}`)

	ctx := context.Background()
	var moduleAddeds entity.ModuleAddedTheGraphResponse
	if err := c.client.Run(ctx, req, &moduleAddeds); err != nil {
		return nil, fmt.Errorf("failed to get added modules from the graph. error: %v", err)
	}

	mods := []*entity.ModuleAdded{}
	for _, mod := range moduleAddeds.ModuleAdded {
		mods = append(mods, mod)
	}

	return mods, nil
}

func (c *theGraphServiceBetaImpl) GetRemovedModules() ([]*entity.ModuleRemoved, error) {
	req := graphql.NewRequest(`
    {
		moduleRemoveds {
			name
			module
		}
	}`)

	ctx := context.Background()
	var moduleRemoveds entity.ModuleRemovedTheGraphResponse
	if err := c.client.Run(ctx, req, &moduleRemoveds); err != nil {
		return nil, fmt.Errorf("failed to get removed modules from the graph. error: %v", err)
	}

	mods := []*entity.ModuleRemoved{}
	for _, mod := range moduleRemoveds.ModuleRemoved {
		mods = append(mods, mod)
	}

	return mods, nil
}
