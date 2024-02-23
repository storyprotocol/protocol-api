package betav0

import (
	"context"
	"fmt"
	"github.com/machinebox/graphql"
	beta_v0 "github.com/storyprotocol/protocol-api/api/internal/models/betav0"
)

func (c *ServiceBetaImpl) GetRoyaltyLiquidSplit(royaltySplitId string) (*beta_v0.RoyaltySplit, error) {
	query := fmt.Sprintf(`
		query {
		  liquidSplit(id: "%s") {
			id
			holders {
				id
				ownership
			}
		  }
		}
    `, royaltySplitId)

	req := graphql.NewRequest(query)
	ctx := context.Background()
	var royRes beta_v0.RoyaltySplitTheGraphResponse
	if err := c.splitClient.Run(ctx, req, &royRes); err != nil {
		return nil, fmt.Errorf("failed to get royalty split from the graph. error: %v", err)
	}

	return royRes.RoyaltySplit, nil

}
