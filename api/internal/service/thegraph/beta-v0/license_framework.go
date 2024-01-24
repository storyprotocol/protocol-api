package beta_v0

import (
	"context"
	"fmt"
	"github.com/machinebox/graphql"
	beta_v0 "github.com/storyprotocol/protocol-api/api/internal/models/beta-v0"
	"github.com/storyprotocol/protocol-api/api/internal/service/thegraph"
)

func (c *ServiceBetaImpl) GetLicenseFramework(licenseId string) (*beta_v0.LicenseFramework, error) {
	query := fmt.Sprintf(`
		{
		  licenseFramework(id: "%s") {
			creator
			id
			frameworkCreationParams {
			  activationParamDefaultValues
			  activationParamVerifiers
			  defaultNeedsActivation
			  licenseUrl
			  linkParentParamDefaultValues
			  linkParentParamVerifiers
			  mintingParamDefaultValues
			  mintingParamVerifiers
			  id
			}
		  }
		}
    `, licenseId)

	req := graphql.NewRequest(query)
	ctx := context.Background()
	var licensesRes beta_v0.LicenseFrameworkTheGraphResponse
	if err := c.client.Run(ctx, req, &licensesRes); err != nil {
		return nil, fmt.Errorf("failed to get license from the graph. error: %v", err)
	}

	return licensesRes.LicenseFramework, nil

}

func (c *ServiceBetaImpl) ListLicenseFrameworks(options *thegraph.TheGraphQueryOptions) ([]*beta_v0.LicenseFramework, error) {
	query := fmt.Sprintf(`
		query(%s) {
		  licenseFrameworks(%s) {
			id
			creator
			
		  }
		}
    `, QUERY_INTERFACE, QUERY_VALUE)
	//frameworkCreationParams {
	//	id
	//	activationParamDefaultValues
	//	activationParamVerifiers
	//	defaultNeedsActivation
	//	linkParentParamDefaultValues
	//	linkParentParamVerifiers
	//	mintingParamDefaultValues
	//	mintingParamVerifiers
	//	licenseUrl
	//}
	req := c.buildNewRequest(options, query)

	ctx := context.Background()
	var licensesRes beta_v0.LicenseFrameworksTheGraphResponse
	if err := c.client.Run(ctx, req, &licensesRes); err != nil {
		return nil, fmt.Errorf("failed to get license frameworks from the graph. error: %v", err)
	}

	licenses := []*beta_v0.LicenseFramework{}
	for _, license := range licensesRes.LicenseFrameworks {
		licenses = append(licenses, license)
	}

	return licenses, nil
}
