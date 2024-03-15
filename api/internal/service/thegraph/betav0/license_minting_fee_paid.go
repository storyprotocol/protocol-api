package betav0

import (
	"context"
	"fmt"
	"github.com/machinebox/graphql"
	beta_v0 "github.com/storyprotocol/protocol-api/api/internal/models/betav0"
	"github.com/storyprotocol/protocol-api/api/internal/service/thegraph"
)

func (c *ServiceBetaImpl) GetLicenseMintingFeePaid(lmfpId string) (*beta_v0.LicenseMintingFeePaid, *beta_v0.RenLicenseMintingFeePaid, error) {
	if c.apiKey != "" {
		query := fmt.Sprintf(`
		query {
		  record(id: "%s") {
			id
			receiver_ip_id
			payer
			token
			amount
			block_number
			block_time
		  }
		}
    `, lmfpId)

		req := c.buildNewRequest(nil, query)
		ctx := context.Background()
		var lmfpRes beta_v0.RenLicenseMintingFeePaidTheGraphResponse
		if err := c.client.Run(ctx, req, &lmfpRes); err != nil {
			return nil, nil, fmt.Errorf("failed to get license minting fee paid from the graph. error: %v", err)
		}

		return nil, lmfpRes.LicenseMintingFeePaid, nil
	} else {
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
			return nil, nil, fmt.Errorf("failed to get license minting fee paid from the graph. error: %v", err)
		}

		return lmfpRes.LicenseMintingFeePaid, nil, nil
	}

}

func (c *ServiceBetaImpl) ListLicenseMintingFeePaids(options *thegraph.TheGraphQueryOptions) ([]*beta_v0.LicenseMintingFeePaid, []*beta_v0.RenLicenseMintingFeePaid, error) {
	whereString := c.buildWhereConditions(options)
	query := ""

	if c.apiKey != "" {
		VALUES := fmt.Sprintf(REN_QUERY_VALUE, ORDER_BY, ORDER_DIRECTION)

		query = fmt.Sprintf(`
		query(%s){
		  records (%s, filter:{%s}) {
			id
			receiver_ip_id
			payer
			token
			amount
			block_number
			block_time
		  }
		}
		`, REN_QUERY_INTERFACE, VALUES, whereString)

		req := c.buildNewRequest(options, query)

		ctx := context.Background()
		var lmfpsRes beta_v0.RenLicenseMintingFeePaidsTheGraphResponse
		if err := c.client.Run(ctx, req, &lmfpsRes); err != nil {
			return nil, nil, fmt.Errorf("failed to get license minting fee paids from the graph. error: %v", err)
		}

		lmfps := []*beta_v0.RenLicenseMintingFeePaid{}
		for _, rPol := range lmfpsRes.LicenseMintingFeePaids {
			lmfps = append(lmfps, rPol)
		}

		return nil, lmfps, nil
	} else {
		query = fmt.Sprintf(`
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

		req := c.buildNewRequest(options, query)

		ctx := context.Background()
		var lmfpsRes beta_v0.LicenseMintingFeePaidsTheGraphResponse
		if err := c.client.Run(ctx, req, &lmfpsRes); err != nil {
			return nil, nil, fmt.Errorf("failed to get license minting fee paids from the graph. error: %v", err)
		}

		lmfps := []*beta_v0.LicenseMintingFeePaid{}
		for _, rPol := range lmfpsRes.LicenseMintingFeePaids {
			lmfps = append(lmfps, rPol)
		}

		return lmfps, nil, nil
	}

}
