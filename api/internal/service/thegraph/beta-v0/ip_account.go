package beta_v0

import (
	"context"
	"fmt"
	"github.com/machinebox/graphql"
	beta_v0 "github.com/storyprotocol/protocol-api/api/internal/models/beta-v0"
	beta_graph "github.com/storyprotocol/protocol-api/api/internal/service/thegraph"
)

func (c *ServiceBetaImpl) GetIPAccount(accountId string) (*beta_v0.IPAccount, error) {
	query := fmt.Sprintf(`
	query {
		iprecord(id: "%s") {
			id
			ipId
			chainId
			tokenContract
			tokenId
			metadataResolverAddress
	  	}
	}
    `, accountId)
	//query := fmt.Sprintf(`
	//query {
	//	ipaccountRegistered(id: "%s") {
	//		account
	//		implementation
	//		chainId
	//		tokenContract
	//		tokenId
	//  	}
	//}
	//`, accountId)

	req := graphql.NewRequest(query)

	ctx := context.Background()
	var ipAccountTheGraphResponse beta_v0.IPAccountTheGraphResponse
	if err := c.client.Run(ctx, req, &ipAccountTheGraphResponse); err != nil {
		return nil, fmt.Errorf("failed to get account from the graph. error: %v", err)
	}

	return ipAccountTheGraphResponse.IPAccount, nil
}

func (c *ServiceBetaImpl) ListIPAccounts(options *beta_graph.TheGraphQueryOptions) ([]*beta_v0.IPAccount, error) {
	whereString := c.buildWhereConditions(options)
	query := fmt.Sprintf(`
	query(%s) {
		iprecords (%s, where:{%s}) {
			id
			ipId
			chainId
			tokenContract
			tokenId
			metadataResolverAddress
		}
    }
    `, QUERY_INTERFACE, QUERY_VALUE, whereString)

	req := c.buildNewRequest(options, query)

	ctx := context.Background()
	var ipAccountsTheGraphResponse beta_v0.IPAccountsTheGraphResponse
	if err := c.client.Run(ctx, req, &ipAccountsTheGraphResponse); err != nil {
		return nil, fmt.Errorf("failed to get registered ip accounts from the graph. error: %v", err)
	}

	ipAccounts := []*beta_v0.IPAccount{}
	for _, ipAccount := range ipAccountsTheGraphResponse.IPAccounts {
		ipAccounts = append(ipAccounts, ipAccount)
	}

	return ipAccounts, nil
}
