package betav0

import (
	"context"
	"fmt"
	"github.com/machinebox/graphql"
	beta_v0 "github.com/storyprotocol/protocol-api/api/internal/models/betav0"
	"github.com/storyprotocol/protocol-api/api/internal/service/thegraph"
)

func (c *ServiceBetaImpl) GetLicenseFramework(licenseId string) (*beta_v0.LicenseFramework, *beta_v0.RenLicenseFramework, error) {
	if c.apiKey == "" {
		query := fmt.Sprintf(`
		{
		  record(id: "%s") {
			creator
			id
		  }
		}
    `, licenseId)

		req := graphql.NewRequest(query)
		ctx := context.Background()
		var licensesRes beta_v0.RenLicenseFrameworkTheGraphResponse
		if err := c.client.Run(ctx, req, &licensesRes); err != nil {
			return nil, nil, fmt.Errorf("failed to get license from the graph. error: %v", err)
		}

		return nil, licensesRes.LicenseFramework, nil
	} else {
		query := fmt.Sprintf(`
		{
		  licenseFramework(id: "%s") {
			creator
			id
		  }
		}
    `, licenseId)

		req := graphql.NewRequest(query)
		ctx := context.Background()
		var licensesRes beta_v0.LicenseFrameworkTheGraphResponse
		if err := c.client.Run(ctx, req, &licensesRes); err != nil {
			return nil, nil, fmt.Errorf("failed to get license from the graph. error: %v", err)
		}

		return licensesRes.LicenseFramework, nil, nil
	}

}

func (c *ServiceBetaImpl) ListLicenseFrameworks(options *thegraph.TheGraphQueryOptions) ([]*beta_v0.LicenseFramework, []*beta_v0.RenLicenseFramework, error) {
	whereString := c.buildWhereConditions(options)
	query := ""

	if c.apiKey != "" {
		VALUES := fmt.Sprintf(REN_QUERY_VALUE, ORDER_BY, ORDER_DIRECTION)

		query = fmt.Sprintf(`
		query(%s) {
		  records(%s, filter:{%s}) {
			id
			creator
		  }
		}
    `, REN_QUERY_INTERFACE, VALUES, whereString)

		req := c.buildNewRequest(options, query)

		ctx := context.Background()
		var licensesRes beta_v0.RenLicenseFrameworksTheGraphResponse
		if err := c.client.Run(ctx, req, &licensesRes); err != nil {
			return nil, nil, fmt.Errorf("failed to get license frameworks from the graph. error: %v", err)
		}

		licenses := []*beta_v0.RenLicenseFramework{}
		for _, license := range licensesRes.LicenseFrameworks {
			licenses = append(licenses, license)
		}

		return nil, licenses, nil
	} else {
		query = fmt.Sprintf(`
		query(%s) {
		  licenseFrameworks(%s, where:{%s}) {
			id
			creator
			
		  }
		}
    `, QUERY_INTERFACE, QUERY_VALUE, whereString)

		req := c.buildNewRequest(options, query)

		ctx := context.Background()
		var licensesRes beta_v0.LicenseFrameworksTheGraphResponse
		if err := c.client.Run(ctx, req, &licensesRes); err != nil {
			return nil, nil, fmt.Errorf("failed to get license frameworks from the graph. error: %v", err)
		}

		licenses := []*beta_v0.LicenseFramework{}
		for _, license := range licensesRes.LicenseFrameworks {
			licenses = append(licenses, license)
		}

		return licenses, nil, nil
	}

}
