package beta_v0

import (
	"context"
	"fmt"
	"github.com/machinebox/graphql"
	beta_v0 "github.com/storyprotocol/protocol-api/api/internal/models/beta-v0"
	"github.com/storyprotocol/protocol-api/api/internal/service/thegraph"
)

func (c *ServiceBetaImpl) GetLicense(licenseId string) ([]*beta_v0.License, error) {
	query := fmt.Sprintf(`
		query {
		  license(id: "%s") {
			id
			licenseData {
			  licensorIpIds
			  policyId
			}
			amount
			creator
			licenseId
			receiver
		  }
		}
    `, licenseId)

	req := graphql.NewRequest(query)
	ctx := context.Background()
	var licensesRes beta_v0.LicenseTheGraphResponse
	if err := c.client.Run(ctx, req, &licensesRes); err != nil {
		return nil, fmt.Errorf("failed to get license from the graph. error: %v", err)
	}

	licenses := []*beta_v0.License{}
	for _, license := range licensesRes.License {
		licenses = append(licenses, license)
	}

	return licenses, nil

}

func (c *ServiceBetaImpl) ListLicenses(options *thegraph.TheGraphQueryOptions) ([]*beta_v0.License, error) {
	query := fmt.Sprintf(`
	query(%s){
		{
		  licenses (%s) {
			amount
			creator
			licenseId
			receiver
			licenseData {
			  licensorIpIds
			  policyId
			  id
			}
		  }
		}
	}
    `, QUERY_INTERFACE, QUERY_VALUE)

	req := c.buildNewRequest(options, query)

	ctx := context.Background()
	var licensesRes beta_v0.LicensesTheGraphResponse
	if err := c.client.Run(ctx, req, &licensesRes); err != nil {
		return nil, fmt.Errorf("failed to get licenses from the graph. error: %v", err)
	}

	licenses := []*beta_v0.License{}
	for _, license := range licensesRes.Licenses {
		licenses = append(licenses, license)
	}

	return licenses, nil
}
