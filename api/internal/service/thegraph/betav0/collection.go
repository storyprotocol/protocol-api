package betav0

import (
	"context"
	"fmt"

	beta_v0 "github.com/storyprotocol/protocol-api/api/internal/models/betav0"
	"github.com/storyprotocol/protocol-api/api/internal/service/thegraph"
)

func (c *ServiceBetaImpl) GetCollection(collectionId string) (*beta_v0.Collection, error) {
	query := fmt.Sprintf(`
		query {
		  collection(id: "%s") {
			id
			assetCount
			licensesCount
			resolvedDisputeCount
			cancelledDisputeCount
			raisedDisputeCount
			judgedDisputesCount
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
		return nil, fmt.Errorf("failed to get collection from the graph. error: %v", err)
	}

	return colRes.Collection, nil
}

func (c *ServiceBetaImpl) ListCollections(options *thegraph.TheGraphQueryOptions) ([]*beta_v0.Collection, error) {
	whereString := c.buildWhereConditions(options)
	query := fmt.Sprintf(`
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
		return nil, fmt.Errorf("failed to get collections from the graph. error: %v", err)
	}

	cols := []*beta_v0.Collection{}
	for _, col := range colsRes.Collections {
		cols = append(cols, col)
	}

	return cols, nil
}
