package betav0

import (
	"context"
	"fmt"
	"github.com/machinebox/graphql"
	beta_v0 "github.com/storyprotocol/protocol-api/api/internal/models/betav0"
	"github.com/storyprotocol/protocol-api/api/internal/service/thegraph"
)

func (c *ServiceBetaImpl) GetLicenseMintingFeePaid(lmfpId string) (*beta_v0.LicenseMintingFeePaid, error) {
	query := fmt.Sprintf(`
		query {
		  licenseMintingFeePaidEntity(id: "%s") {
			id
			receiverIpId
			payer
			token
			amount
			blockNumber
			blockTimestamp
		  }
		}
    `, lmfpId)

	req := graphql.NewRequest(query)
	ctx := context.Background()
	var lmfpRes beta_v0.LicenseMintingFeePaidTheGraphResponse
	if err := c.client.Run(ctx, req, &lmfpRes); err != nil {
		return nil, fmt.Errorf("failed to get license minting fee paid from the graph. error: %v", err)
	}

	return lmfpRes.LicenseMintingFeePaid, nil

}

func (c *ServiceBetaImpl) ListLicenseMintingFeePaids(options *thegraph.TheGraphQueryOptions) ([]*beta_v0.LicenseMintingFeePaid, error) {
	whereString := c.buildWhereConditions(options)
	query := fmt.Sprintf(`
	query(%s){
	  licenseMintingFeePaidEntities (%s, where:{%s}) {
		id
		receiverIpId
		payer
		token
		amount
		blockNumber
		blockTimestamp
	  }
	}
    `, QUERY_INTERFACE, QUERY_VALUE, whereString)

	fmt.Println(query)

	req := c.buildNewRequest(options, query)

	ctx := context.Background()
	var lmfpsRes beta_v0.LicenseMintingFeePaidsTheGraphResponse
	if err := c.client.Run(ctx, req, &lmfpsRes); err != nil {
		return nil, fmt.Errorf("failed to get license minting fee paids from the graph. error: %v", err)
	}

	lmfps := []*beta_v0.LicenseMintingFeePaid{}
	for _, rPol := range lmfpsRes.LicenseMintingFeePaids {
		lmfps = append(lmfps, rPol)
	}

	return lmfps, nil
}
