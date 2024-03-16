package betav0

import (
	"context"
	"fmt"
	"github.com/machinebox/graphql"
	beta_v0 "github.com/storyprotocol/protocol-api/api/internal/models/betav0"
	"github.com/storyprotocol/protocol-api/api/internal/service/thegraph"
)

func (c *ServiceBetaImpl) GetLicense(licenseId string) (*beta_v0.License, *beta_v0.RenLicense, error) {
	if c.apiKey != "" {
		query := fmt.Sprintf(`
		query {
		  record(id: "%s") {
			id
			policy_id
			licensor_ip_id
			amount
			transferable
			block_number
			block_time
		  }
		}
    `, licenseId)

		req := c.buildNewRequest(nil, query)
		ctx := context.Background()
		var licensesRes beta_v0.RenLicenseTheGraphResponse
		if err := c.client.Run(ctx, req, &licensesRes); err != nil {
			return nil, nil, fmt.Errorf("failed to get license from the graph. error: %v", err)
		}

		return nil, licensesRes.License, nil
	} else {
		query := fmt.Sprintf(`
		query {
		  license(id: "%s") {
			id
			policyId
			licensorIpId
			amount
			transferable
			blockNumber
			blockTimestamp
		  }
		}
    `, licenseId)

		req := graphql.NewRequest(query)
		ctx := context.Background()
		var licensesRes beta_v0.LicenseTheGraphResponse
		if err := c.client.Run(ctx, req, &licensesRes); err != nil {
			return nil, nil, fmt.Errorf("failed to get license from the graph. error: %v", err)
		}

		return licensesRes.License, nil, nil
	}

}

func (c *ServiceBetaImpl) ListLicenses(options *thegraph.TheGraphQueryOptions) ([]*beta_v0.License, []*beta_v0.RenLicense, error) {
	whereString := c.buildWhereConditions(options)
	query := ""

	if c.apiKey != "" {
		VALUES := fmt.Sprintf(REN_QUERY_VALUE, ORDER_BY, ORDER_DIRECTION)

		query = fmt.Sprintf(`
		query(%s){
		  records(%s,filter:{%s}){
			id
			policy_id
			licensor_ip_id
			amount
			transferable
			block_time
			block_number
		  }
		}
		`, REN_QUERY_INTERFACE, VALUES, whereString)

		req := c.buildNewRequest(options, query)

		ctx := context.Background()
		var licensesRes beta_v0.RenLicensesTheGraphResponse
		if err := c.client.Run(ctx, req, &licensesRes); err != nil {
			return nil, nil, fmt.Errorf("failed to get licenses from the graph. error: %v", err)
		}

		licenses := []*beta_v0.RenLicense{}
		for _, license := range licensesRes.Licenses {
			licenses = append(licenses, license)
		}

		return nil, licenses, nil
	} else {
		query = fmt.Sprintf(`
	query(%s){
	  licenses (%s, where:{%s}) {
		id
		policyId
		licensorIpId
		amount
		transferable
		blockTimestamp
		blockNumber
	  }
	}
    `, QUERY_INTERFACE, QUERY_VALUE, whereString)

		req := c.buildNewRequest(options, query)

		ctx := context.Background()
		var licensesRes beta_v0.LicensesTheGraphResponse
		if err := c.client.Run(ctx, req, &licensesRes); err != nil {
			return nil, nil, fmt.Errorf("failed to get licenses from the graph. error: %v", err)
		}

		licenses := []*beta_v0.License{}
		for _, license := range licensesRes.Licenses {
			licenses = append(licenses, license)
		}

		return licenses, nil, nil
	}

}
