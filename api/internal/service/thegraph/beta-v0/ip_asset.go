package beta_v0

import (
	"context"
	"fmt"
	"github.com/machinebox/graphql"
	beta_v0 "github.com/storyprotocol/protocol-api/api/internal/models/beta-v0"
	beta_graph "github.com/storyprotocol/protocol-api/api/internal/service/thegraph"
)

func (c *ServiceBetaImpl) GetIPAsset(assetId string) (*beta_v0.IPAsset, error) {
	query := fmt.Sprintf(`
	query {
		ipasset(id: "%s") {
			id
			ipId
			chainId
			tokenContract
			tokenId
			metadataResolverAddress
			metadata
	  	}	
	}
    `, assetId)
	//query := fmt.Sprintf(`
	//query {
	//	ipassetRegistered(id: "%s") {
	//		asset
	//		implementation
	//		chainId
	//		tokenContract
	//		tokenId
	//  	}
	//}
	//`, assetId)

	req := graphql.NewRequest(query)

	ctx := context.Background()
	var ipAssetTheGraphResponse beta_v0.IPAssetTheGraphResponse
	if err := c.client.Run(ctx, req, &ipAssetTheGraphResponse); err != nil {
		return nil, fmt.Errorf("failed to get asset from the graph. error: %v", err)
	}

	return ipAssetTheGraphResponse.IPAsset, nil
}

func (c *ServiceBetaImpl) ListIPAssets(options *beta_graph.TheGraphQueryOptions) ([]*beta_v0.IPAsset, error) {
	whereString := c.buildWhereConditions(options)
	query := fmt.Sprintf(`
	query(%s) {
		ipassets (%s, where:{%s}) {
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
	var ipAssetsTheGraphResponse beta_v0.IPAssetsTheGraphResponse
	if err := c.client.Run(ctx, req, &ipAssetsTheGraphResponse); err != nil {
		return nil, fmt.Errorf("failed to get registered ip assets from the graph. error: %v", err)
	}

	ipAssets := []*beta_v0.IPAsset{}
	for _, ipAsset := range ipAssetsTheGraphResponse.IPAssets {
		ipAssets = append(ipAssets, ipAsset)
	}

	return ipAssets, nil
}
