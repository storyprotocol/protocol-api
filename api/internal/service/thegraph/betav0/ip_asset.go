package betav0

import (
	"context"
	"fmt"
	"github.com/machinebox/graphql"
	beta_v0 "github.com/storyprotocol/protocol-api/api/internal/models/betav0"
	beta_graph "github.com/storyprotocol/protocol-api/api/internal/service/thegraph"
	"github.com/storyprotocol/protocol-api/pkg/logger"
)

func (c *ServiceBetaImpl) GetIPAsset(assetId string) (*beta_v0.IPAsset, *beta_v0.RenIPAsset, error) {
	if c.apiKey != "" {
		query := fmt.Sprintf(`
		query {
			record(id: "%s") {
				id
				chain_id
				token_contract
				token_id
				metadata_resolver_address
				block_number
				block_time
			}	
		}
		`, assetId)

		req := c.buildNewRequest(nil, query)
		ctx := context.Background()
		var ipAssetTheGraphResponse beta_v0.RenIPAssetTheGraphResponse
		if err := c.client.Run(ctx, req, &ipAssetTheGraphResponse); err != nil {
			return nil, nil, fmt.Errorf("failed to get asset from the graph. error: %v", err)
		}

		return nil, ipAssetTheGraphResponse.IPAsset, nil
	} else {
		query := fmt.Sprintf(`
	query {
		ipasset(id: "%s") {
			id
			chainId
			parentIpIds	{
				id
				chainId
				tokenContract
				tokenId
				metadataResolverAddress
				blockNumber
				blockTimestamp
				metadata {
					name
					hash
					registrationDate
					registrant
					uri
				}
			}
			childIpIds {
				id
				chainId
				tokenContract
				tokenId
				metadataResolverAddress
				blockNumber
				blockTimestamp
				metadata {
					name
					hash
					registrationDate
					registrant
					uri
				}
			}
			rootIpIds {
				id
				chainId
				tokenContract
				tokenId
				metadataResolverAddress
				blockNumber
				blockTimestamp
				metadata {
					name
					hash
					registrationDate
					registrant
					uri
				}
			}
			tokenContract
			tokenId
			metadataResolverAddress
			blockNumber
			blockTimestamp
			metadata {
				name
				hash
				registrationDate
				registrant
				uri
			}
			
	  	}	
	}
    `, assetId)

		req := graphql.NewRequest(query)

		ctx := context.Background()
		var ipAssetTheGraphResponse beta_v0.IPAssetTheGraphResponse
		if err := c.client.Run(ctx, req, &ipAssetTheGraphResponse); err != nil {
			return nil, nil, fmt.Errorf("failed to get asset from the graph. error: %v", err)
		}

		return ipAssetTheGraphResponse.IPAsset, nil, nil
	}

}

func (c *ServiceBetaImpl) ListIPAssets(options *beta_graph.TheGraphQueryOptions) ([]*beta_v0.IPAsset, []*beta_v0.RenIPAsset, error) {
	whereString := c.buildWhereConditions(options)
	query := ""

	logger.Info(whereString)
	if c.apiKey != "" {
		VALUES := fmt.Sprintf(REN_QUERY_VALUE, ORDER_BY, ORDER_DIRECTION)
		//TODO: What we need
		//query(%s) {
		//	records (%s, filter:{%s}) {
		//		id
		//		chain_id
		//		parent_ip_ids	{
		//			id
		//			chain_id
		//			token_contract
		//			token_id
		//			metadata_resolver_address
		//			block_number
		//			block_timestamp
		//			metadata{
		//				name
		//				hash
		//				registration_date
		//				registrant
		//				uri
		//			}
		//		}
		//		child_ip_ids{
		//			id
		//			chain_id
		//			token_contract
		//			token_id
		//			metadata_resolver_address
		//			block_number
		//			block_timestamp
		//			metadata{
		//				name
		//				hash
		//				registration_date
		//				registrant
		//				uri
		//			}
		//		}
		//		root_ip_ids{
		//			id
		//			chain_id
		//			token_contract
		//			token_id
		//			metadata_resolver_address
		//			block_number
		//			block_timestamp
		//			metadata{
		//				name
		//				hash
		//				registration_date
		//				registrant
		//				uri
		//			}
		//		}
		//		token_contract
		//		token_id
		//		metadata_resolver_address
		//		block_number
		//		block_timestamp
		//		metadata{
		//			name
		//			hash
		//			registration_date
		//			registrant
		//			uri
		//		}
		//	}
		//}
		query = fmt.Sprintf(`
		query(%s) {
			records (%s, filter:{%s}) {
				id
				chain_id
				token_contract
				token_id
				metadata_resolver_address
				block_number
				block_time
			}
	}
    `, REN_QUERY_INTERFACE, VALUES, whereString)

		req := c.buildNewRequest(options, query)

		ctx := context.Background()
		var ipAssetsTheGraphResponse beta_v0.RenIPAssetsTheGraphResponse
		if err := c.client.Run(ctx, req, &ipAssetsTheGraphResponse); err != nil {
			return nil, nil, fmt.Errorf("failed to get registered ip assets from the graph. error: %v", err)
		}

		ipAssets := []*beta_v0.RenIPAsset{}
		for _, ipAsset := range ipAssetsTheGraphResponse.IPAssets {
			ipAssets = append(ipAssets, ipAsset)
		}

		return nil, ipAssets, nil
	} else {
		query = fmt.Sprintf(`
	query(%s) {
		ipassets (%s, where:{%s}) {
			id
			chainId
			parentIpIds	{
				id
				chainId
				tokenContract
				tokenId
				metadataResolverAddress
				blockNumber
				blockTimestamp
				metadata {
					name
					hash
					registrationDate
					registrant
					uri
				}
			}
			childIpIds {
				id
				chainId
				tokenContract
				tokenId
				metadataResolverAddress
				blockNumber
				blockTimestamp
				metadata {
					name
					hash
					registrationDate
					registrant
					uri
				}
			}
			rootIpIds {
				id
				chainId
				tokenContract
				tokenId
				metadataResolverAddress
				blockNumber
				blockTimestamp
				metadata {
					name
					hash
					registrationDate
					registrant
					uri
				}
			}
			tokenContract
			tokenId
			metadataResolverAddress
			blockNumber
			blockTimestamp
			metadata {
				name
				hash
				registrationDate
				registrant
				uri
			}
		}
    }
    `, QUERY_INTERFACE, QUERY_VALUE, whereString)

		req := c.buildNewRequest(options, query)

		ctx := context.Background()
		var ipAssetsTheGraphResponse beta_v0.IPAssetsTheGraphResponse
		if err := c.client.Run(ctx, req, &ipAssetsTheGraphResponse); err != nil {
			return nil, nil, fmt.Errorf("failed to get registered ip assets from the graph. error: %v", err)
		}

		ipAssets := []*beta_v0.IPAsset{}
		for _, ipAsset := range ipAssetsTheGraphResponse.IPAssets {
			ipAssets = append(ipAssets, ipAsset)
		}

		return ipAssets, nil, nil
	}

}
