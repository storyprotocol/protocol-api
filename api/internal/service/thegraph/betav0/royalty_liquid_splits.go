package betav0

import (
	"context"
	"fmt"
	"github.com/machinebox/graphql"
	beta_v0 "github.com/storyprotocol/protocol-api/api/internal/models/betav0"
	"strings"
)

func (c *ServiceBetaImpl) GetRoyaltyLiquidSplit(royaltySplitId string) (*beta_v0.RoyaltySplit, *beta_v0.RenRoyaltySplit, error) {
	query := ""

	if c.apiKey == "" {
		query = fmt.Sprintf(`
		query {
		  record(id: "%s") {
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
		var royRes beta_v0.RenRoyaltySplitTheGraphResponse
		if err := c.splitClient.Run(ctx, req, &royRes); err != nil {
			return nil, nil, fmt.Errorf("failed to get royalty split from the graph. error: %v", err)
		}

		var holders []beta_v0.RenHolder
		claimFromIpPoolArg := "["

		for i, holder := range royRes.RoyaltySplit.Holders {
			sHolder := c.formatRenHolder(holder)
			holders = append(holders, sHolder)

			if sHolder.Ownership != "" {
				if i != 0 {
					claimFromIpPoolArg = claimFromIpPoolArg + ","
				}
				claimFromIpPoolArg = claimFromIpPoolArg + fmt.Sprintf("%v", sHolder.ID)

			}
		}

		claimFromIpPoolArg = claimFromIpPoolArg + "]"

		royRes.RoyaltySplit.ClaimFromIPPoolArg = claimFromIpPoolArg
		royRes.RoyaltySplit.Holders = holders

		return nil, royRes.RoyaltySplit, nil
	} else {
		query = fmt.Sprintf(`
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
			return nil, nil, fmt.Errorf("failed to get royalty split from the graph. error: %v", err)
		}

		var holders []beta_v0.Holder
		claimFromIpPoolArg := "["

		for i, holder := range royRes.RoyaltySplit.Holders {
			sHolder := c.formatHolder(holder)
			holders = append(holders, sHolder)

			if sHolder.Ownership != "" {
				if i != 0 {
					claimFromIpPoolArg = claimFromIpPoolArg + ","
				}
				claimFromIpPoolArg = claimFromIpPoolArg + fmt.Sprintf("%v", sHolder.ID)

			}
		}

		claimFromIpPoolArg = claimFromIpPoolArg + "]"

		royRes.RoyaltySplit.ClaimFromIPPoolArg = claimFromIpPoolArg
		royRes.RoyaltySplit.Holders = holders

		return royRes.RoyaltySplit, nil, nil
	}

}

func (c *ServiceBetaImpl) formatHolder(_holder beta_v0.Holder) beta_v0.Holder {
	var holder beta_v0.Holder
	splitId := strings.Split(_holder.ID, "-")
	holder.ID = splitId[1]

	if len(_holder.Ownership) > 4 {
		holder.Ownership = _holder.Ownership[:len(_holder.Ownership)-3]
	}
	return holder
}

func (c *ServiceBetaImpl) formatRenHolder(_holder beta_v0.RenHolder) beta_v0.RenHolder {
	var holder beta_v0.RenHolder
	splitId := strings.Split(_holder.ID, "-")
	holder.ID = splitId[1]

	if len(_holder.Ownership) > 4 {
		holder.Ownership = _holder.Ownership[:len(_holder.Ownership)-3]
	}
	return holder
}
