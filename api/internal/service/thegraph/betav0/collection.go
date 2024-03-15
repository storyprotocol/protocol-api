package betav0

import (
	"context"
	"fmt"

	beta_v0 "github.com/storyprotocol/protocol-api/api/internal/models/betav0"
	"github.com/storyprotocol/protocol-api/api/internal/service/thegraph"
)

func (c *ServiceBetaImpl) GetCollection(collectionId string) (*beta_v0.Collection, *beta_v0.RenCollection, error) {
	if c.apiKey != "" {
		query := fmt.Sprintf(`
		query {
		  record(id: "%s") {
			id
			asset_count
			licenses_count
			resolved_dispute_count
			cancelled_dispute_count
			raised_dispute_count
			judged_dispute_count
			asset_count
			block_number
			block_time
		  }
		}
    `, collectionId)

		req := c.buildNewRequest(nil, query)
		ctx := context.Background()
		var colRes beta_v0.RenCollectionTheGraphResponse
		if err := c.client.Run(ctx, req, &colRes); err != nil {
			return nil, nil, fmt.Errorf("failed to get collection from the graph. error: %v", err)
		}

		return nil, colRes.Collection, nil
	} else {
		query := fmt.Sprintf(`
		query {
		  collection(id: "%s") {
			id
			assetCount
			licensesCount
			resolvedDisputeCount
			cancelledDisputeCount
			raisedDisputeCount
			judgedDisputeCount
			assetCount
			blockNumber
			blockTimestamp
		  }
		}
    `, collectionId)

		req := c.buildNewRequest(nil, query)
		ctx := context.Background()
		var colRes beta_v0.CollectionTheGraphResponse
		if err := c.client.Run(ctx, req, &colRes); err != nil {
			return nil, nil, fmt.Errorf("failed to get collection from the graph. error: %v", err)
		}

		return colRes.Collection, nil, nil
	}

}

func (c *ServiceBetaImpl) ListCollections(options *thegraph.TheGraphQueryOptions) ([]*beta_v0.Collection, []*beta_v0.RenCollection, error) {
	whereString := c.buildWhereConditions(options)
	query := ""

	if c.apiKey != "" {
		VALUES := fmt.Sprintf(REN_QUERY_VALUE, ORDER_BY, ORDER_DIRECTION)
		query = fmt.Sprintf(`
		query(%s){
			records (%s, filter:{%s}) {
				id
				asset_count
				licenses_count
				resolved_dispute_count
				cancelled_dispute_count
				raised_dispute_count
				judged_dispute_count
				block_number
				block_time
			}
		}
		`, REN_QUERY_INTERFACE, VALUES, whereString)

		req := c.buildNewRequest(options, query)

		ctx := context.Background()
		var colsRes beta_v0.RenCollectionsTheGraphResponse
		if err := c.client.Run(ctx, req, &colsRes); err != nil {
			return nil, nil, fmt.Errorf("failed to get collections from the graph. error: %v", err)
		}

		cols := []*beta_v0.RenCollection{}
		for _, col := range colsRes.Collections {
			cols = append(cols, col)
		}
		return nil, cols, nil
	} else {
		query = fmt.Sprintf(`
		query(%s){
			collections (%s, where:{%s}) {
				id
				assetCount
				licensesCount
				resolvedDisputeCount
				cancelledDisputeCount
				raisedDisputeCount
				judgedDisputeCount
				blockNumber
				blockTimestamp
			}
		}
		`, QUERY_INTERFACE, QUERY_VALUE, whereString)

		req := c.buildNewRequest(options, query)

		ctx := context.Background()
		var colsRes beta_v0.CollectionsTheGraphResponse
		if err := c.client.Run(ctx, req, &colsRes); err != nil {
			return nil, nil, fmt.Errorf("failed to get collections from the graph. error: %v", err)
		}

		cols := []*beta_v0.Collection{}
		for _, col := range colsRes.Collections {
			cols = append(cols, col)
		}
		return cols, nil, nil

	}

}
