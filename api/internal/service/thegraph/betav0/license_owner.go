package betav0

import (
	"context"
	"fmt"
	"github.com/machinebox/graphql"
	beta_v0 "github.com/storyprotocol/protocol-api/api/internal/models/betav0"
	"github.com/storyprotocol/protocol-api/api/internal/service/thegraph"
)

func (c *ServiceBetaImpl) GetLicenseOwner(licenseOwnerId string) (*beta_v0.LicenseOwner, *beta_v0.RenLicenseOwner, error) {
	if c.apiKey != "" {
		query := fmt.Sprintf(`
			query {
			  record(id: "%s") {
				id
				policy_id
				owner
				amount
				block_number
				block_time
			  }
			}
		`, licenseOwnerId)

		req := c.buildNewRequest(nil, query)
		ctx := context.Background()
		var licensesRes beta_v0.RenLicenseOwnerTheGraphResponse
		if err := c.client.Run(ctx, req, &licensesRes); err != nil {
			return nil, nil, fmt.Errorf("failed to get license owner from the graph. error: %v", err)
		}

		return nil, licensesRes.LicenseOwner, nil
	} else {
		query := fmt.Sprintf(`
		query {
		  licenseOwner(id: "%s") {
			id
			policyId
			owner
			amount
			blockNumber
			blockTimestamp
		  }
		}
    `, licenseOwnerId)

		req := graphql.NewRequest(query)
		ctx := context.Background()
		var licensesRes beta_v0.LicenseOwnerTheGraphResponse
		if err := c.client.Run(ctx, req, &licensesRes); err != nil {
			return nil, nil, fmt.Errorf("failed to get license owner from the graph. error: %v", err)
		}

		return licensesRes.LicenseOwner, nil, nil
	}

}

func (c *ServiceBetaImpl) ListLicenseOwners(options *thegraph.TheGraphQueryOptions) ([]*beta_v0.LicenseOwner, []*beta_v0.RenLicenseOwner, error) {
	whereString := c.buildWhereConditions(options)
	query := ""

	if c.apiKey != "" {
		VALUES := fmt.Sprintf(REN_QUERY_VALUE, ORDER_BY, ORDER_DIRECTION)

		query = fmt.Sprintf(`
		query(%s){
		  records (%s, filter:{%s}) {
			id
			policy_id
			owner
			amount
			block_number
			block_time
		  }
		}
		`, REN_QUERY_INTERFACE, VALUES, whereString)

		req := c.buildNewRequest(options, query)

		ctx := context.Background()
		var licensesRes beta_v0.RenLicenseOwnersTheGraphResponse
		if err := c.client.Run(ctx, req, &licensesRes); err != nil {
			return nil, nil, fmt.Errorf("failed to get licenses owners from the graph. error: %v", err)
		}

		licenses := []*beta_v0.RenLicenseOwner{}
		for _, license := range licensesRes.LicenseOwners {
			licenses = append(licenses, license)
		}

		return nil, licenses, nil
	} else {
		query = fmt.Sprintf(`
	query(%s){
	  licenseOwners (%s, where:{%s}) {
		id
		policyId
		owner
		amount
		blockNumber
		blockTimestamp
	  }
	}
    `, QUERY_INTERFACE, QUERY_VALUE, whereString)

		req := c.buildNewRequest(options, query)

		ctx := context.Background()
		var licensesRes beta_v0.LicenseOwnersTheGraphResponse
		if err := c.client.Run(ctx, req, &licensesRes); err != nil {
			return nil, nil, fmt.Errorf("failed to get licenses owners from the graph. error: %v", err)
		}

		licenses := []*beta_v0.LicenseOwner{}
		for _, license := range licensesRes.LicenseOwners {
			licenses = append(licenses, license)
		}

		return licenses, nil, nil
	}

}
