package beta_v0

import (
	"github.com/machinebox/graphql"
	"github.com/storyprotocol/protocol-api/api/internal/service/thegraph"
)

const (
	QUERY_INTERFACE = "$first: Int, $skip: Int, $orderBy: String, $orderDirection: String"
	QUERY_VALUE     = "first: $first, skip: $skip, orderBy: $orderBy, orderDirection: $orderDirection"
)

func NewTheGraphServiceBetaImpl(client *graphql.Client) thegraph.TheGraphServiceBeta {
	return &ServiceBetaImpl{
		client: client,
	}
}

type ServiceBetaImpl struct {
	client *graphql.Client
}

func (s *ServiceBetaImpl) buildNewRequest(options *thegraph.TheGraphQueryOptions, query string) *graphql.Request {
	options = s.setQueryOptions(options)

	req := graphql.NewRequest(query)
	req.Var("first", options.First)
	req.Var("skip", options.Skip)
	req.Var("orderBy", options.OrderBy)
	req.Var("orderDirection", options.OrderDirection)

	return req
}

func (s *ServiceBetaImpl) setQueryOptions(options *thegraph.TheGraphQueryOptions) *thegraph.TheGraphQueryOptions {
	if options == nil {
		options = &thegraph.TheGraphQueryOptions{
			First: 100,
			Skip:  0,
		}
	}

	if options.First == 0 {
		options.First = 100
	}

	options.OrderBy = "blockTimestamp"
	options.OrderDirection = "desc"

	return options
}

//func (c *TheGraphServiceBetaImpl) GetIPsRegistered() ([]*models.IPRegistered, error) {
//	req := graphql.NewRequest(`
//    {
//		ipregistereds {
//			id
//			chainId
//			tokenContract
//			tokenId
//			resolver
//		}
//	}`)
//
//	ctx := context.Background()
//	var ipRegisteredsTheGraphResponse models.IPRegisteredTheGraphResponse
//	if err := c.client.Run(ctx, req, &ipRegisteredsTheGraphResponse); err != nil {
//		return nil, fmt.Errorf("failed to get registered ip accounts from the graph. error: %v", err)
//	}
//
//	ips := []*models.IPRegistered{}
//	for _, ip := range ipRegisteredsTheGraphResponse.IPRegistered {
//		ips = append(ips, ip)
//	}
//
//	return ips, nil
//}

//func (c *TheGraphServiceBetaImpl) GetSetIPAccounts() ([]*models.SetIPAccount, error) {
//	req := graphql.NewRequest(`
//    {
//		ipaccountSets {
//			ipId
//			chainId
//			tokenContract
//			tokenId
//		}
//	}`)
//
//	ctx := context.Background()
//	var ipAccountSets models.SetIPAccountTheGraphResponse
//	if err := c.client.Run(ctx, req, &ipAccountSets); err != nil {
//		return nil, fmt.Errorf("failed to get set ip accounts from the graph. error: %v", err)
//	}
//
//	accs := []*models.SetIPAccount{}
//	for _, acc := range ipAccountSets.SetIPAccount {
//		accs = append(accs, acc)
//	}
//
//	return accs, nil
//}

//func (c *TheGraphServiceBetaImpl) GetSetIPResolvers() ([]*models.SetResolver, error) {
//	req := graphql.NewRequest(`
//    {
//		ipresolverSets {
//			ipId
//			resolver
//		}
//	}`)
//
//	ctx := context.Background()
//	var ipResolverSets models.SetResolverTheGraphResponse
//	if err := c.client.Run(ctx, req, &ipResolverSets); err != nil {
//		return nil, fmt.Errorf("failed to get set ip resolvers from the graph. error: %v", err)
//	}
//
//	rslvrs := []*models.SetResolver{}
//	for _, rslvr := range ipResolverSets.SetResolver {
//		rslvrs = append(rslvrs, rslvr)
//	}
//
//	return rslvrs, nil
//}
