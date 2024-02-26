package betav0

import (
	"context"
	"fmt"
	"github.com/machinebox/graphql"
	beta_v0 "github.com/storyprotocol/protocol-api/api/internal/models/betav0"
	"strings"
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

	claimFromIpPoolArg := "["

	for i, holder := range royRes.RoyaltySplit.Holders {
		sHolder := c.formatHolder(holder)
		if sHolder.Ownership != "1000" {
			claimFromIpPoolArg = claimFromIpPoolArg + fmt.Sprintf("%v", sHolder.ID)

			if i == len(royRes.RoyaltySplit.Holders)-1 {
				claimFromIpPoolArg = claimFromIpPoolArg + "]"
			} else {
				claimFromIpPoolArg = claimFromIpPoolArg + ","
			}
		}
	}

	royRes.RoyaltySplit.ClaimFromIPPoolArg = claimFromIpPoolArg

	fmt.Printf(claimFromIpPoolArg)
	return royRes.RoyaltySplit, nil

}

func (c *ServiceBetaImpl) formatHolder(_holder beta_v0.Holder) beta_v0.Holder {
	var holder beta_v0.Holder
	splitId := strings.Split(_holder.ID, "-")
	holder.ID = splitId[1]
	holder.Ownership = strings.Trim(_holder.Ownership, "000")

	return holder
}
