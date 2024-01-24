package thegraph

import (
	"context"
	"fmt"
	"github.com/machinebox/graphql"
	"github.com/storyprotocol/protocol-api/api/internal/models/beta-v0"
)

const (
	QUERY_INTERFACE = "$first: Int, $skip: Int, $orderBy: String, $orderDirection: String"
	QUERY_VALUE     = "first: $first, skip: $skip, orderBy: $orderBy, orderDirection: $orderDirection"
)

func NewTheGraphServiceBetaImpl(client *graphql.Client) TheGraphServiceBeta {
	return &theGraphServiceBetaImpl{
		client: client,
	}
}

type theGraphServiceBetaImpl struct {
	client *graphql.Client
}

func (c *theGraphServiceBetaImpl) GetIPAccount(accountId string) ([]*beta_v0.IPAccount, error) {
	query := fmt.Sprintf(`
	query {
		iprecord(id: "%s") {
			id
			ipId
			chainId
			tokenContract
			tokenId
			metadataResolverAddress
	  	}
	}
    `, accountId)
	//query := fmt.Sprintf(`
	//query {
	//	ipaccountRegistered(id: "%s") {
	//		account
	//		implementation
	//		chainId
	//		tokenContract
	//		tokenId
	//  	}
	//}
	//`, accountId)

	req := graphql.NewRequest(query)

	ctx := context.Background()
	var ipAccountTheGraphResponse beta_v0.IPAccountTheGraphResponse
	if err := c.client.Run(ctx, req, &ipAccountTheGraphResponse); err != nil {
		return nil, fmt.Errorf("failed to get account from the graph. error: %v", err)
	}

	accts := []*beta_v0.IPAccount{}
	accts = append(accts, ipAccountTheGraphResponse.IPAccount)

	return accts, nil
}

func (c *theGraphServiceBetaImpl) ListIPAccounts(options *TheGraphQueryOptions) ([]*beta_v0.IPAccount, error) {
	query := fmt.Sprintf(`
	query(%s) {
		iprecords (%s) {
			id
			ipId
			chainId
			tokenContract
			tokenId
			metadataResolverAddress
		}
    }
    `, QUERY_INTERFACE, QUERY_VALUE)

	req := c.buildNewRequest(options, query)

	ctx := context.Background()
	var ipAccountsTheGraphResponse beta_v0.IPAccountsTheGraphResponse
	if err := c.client.Run(ctx, req, &ipAccountsTheGraphResponse); err != nil {
		return nil, fmt.Errorf("failed to get registered ip accounts from the graph. error: %v", err)
	}

	ipAccounts := []*beta_v0.IPAccount{}
	for _, ipAccount := range ipAccountsTheGraphResponse.IPAccounts {
		ipAccounts = append(ipAccounts, ipAccount)
	}

	return ipAccounts, nil
}

func (c *theGraphServiceBetaImpl) GetModule(moduleName string) ([]*beta_v0.Module, error) {
	query := fmt.Sprintf(`
	query {
		module(id: "%s") {
			name
			module
	  	}
	}
    `, moduleName)

	req := graphql.NewRequest(query)
	ctx := context.Background()
	var modules beta_v0.ModuleTheGraphResponse
	if err := c.client.Run(ctx, req, &modules); err != nil {
		return nil, fmt.Errorf("failed to get modules from the graph. error: %v", err)
	}

	mods := []*beta_v0.Module{}
	for _, mod := range modules.Module {
		mods = append(mods, mod)
	}

	return mods, nil

}

func (c *theGraphServiceBetaImpl) ListModules(options *TheGraphQueryOptions) ([]*beta_v0.Module, error) {
	query := fmt.Sprintf(`
	query(%s){
		modules (%s) {
			name
			module
		}
	}
    `, QUERY_INTERFACE, QUERY_VALUE)

	req := c.buildNewRequest(options, query)

	ctx := context.Background()
	var modules beta_v0.ModulesTheGraphResponse
	if err := c.client.Run(ctx, req, &modules); err != nil {
		return nil, fmt.Errorf("failed to get modules from the graph. error: %v", err)
	}

	mods := []*beta_v0.Module{}
	for _, mod := range modules.Modules {
		mods = append(mods, mod)
	}

	return mods, nil
}

func (c *theGraphServiceBetaImpl) GetLicense(licenseId string) ([]*beta_v0.License, error) {
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

func (c *theGraphServiceBetaImpl) ListLicenses(options *TheGraphQueryOptions) ([]*beta_v0.License, error) {
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

func (c *theGraphServiceBetaImpl) GetLicenseFramework(licenseId string) ([]*beta_v0.LicenseFramework, error) {
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

	licenses := []*beta_v0.LicenseFramework{}
	for _, license := range licensesRes.LicenseFramework {
		licenses = append(licenses, license)
	}

	return licenses, nil

}

func (c *theGraphServiceBetaImpl) ListLicenseFrameworks(options *TheGraphQueryOptions) ([]*beta_v0.LicenseFramework, error) {
	query := fmt.Sprintf(`
		query(%s) {
		  licenseFrameworks(%s) {
			id
			creator
			frameworkCreationParams {
			  id
			  activationParamDefaultValues
			  activationParamVerifiers
			  defaultNeedsActivation
			  linkParentParamDefaultValues
			  linkParentParamVerifiers
			  mintingParamDefaultValues
			  mintingParamVerifiers
			  licenseUrl
			}
		  }
    `, QUERY_INTERFACE, QUERY_VALUE)

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

func (c *theGraphServiceBetaImpl) ListAccessControlPermissions(options *TheGraphQueryOptions) ([]*beta_v0.AccessControlPermission, error) {
	query := fmt.Sprintf(`
	query(%s){
		modules (%s) {
			name
			module
		}
	}
    `, QUERY_INTERFACE, QUERY_VALUE)

	req := c.buildNewRequest(options, query)

	ctx := context.Background()
	var acpsRes beta_v0.AccessControlPermissionTheGraphResponse
	if err := c.client.Run(ctx, req, &acpsRes); err != nil {
		return nil, fmt.Errorf("failed to get access control permissions from the graph. error: %v", err)
	}

	acps := []*beta_v0.AccessControlPermission{}
	for _, acp := range acpsRes.AccessControlPermissions {
		acps = append(acps, acp)
	}

	return acps, nil
}

func (c *theGraphServiceBetaImpl) GetPolicy(policyId string) ([]*beta_v0.Policy, error) {
	query := fmt.Sprintf(`
		query {
		  policy(id: "%s") {
			policyId
			creator
			policyData {     
				id
				frameworkId
				needsActivation
				mintingParamValues
				linkParentParamValues
				activationParamValues
			}
		  }
		}
    `, policyId)

	req := graphql.NewRequest(query)
	ctx := context.Background()
	var polRes beta_v0.PolicyTheGraphResponse
	if err := c.client.Run(ctx, req, &polRes); err != nil {
		return nil, fmt.Errorf("failed to get policy from the graph. error: %v", err)
	}

	pols := []*beta_v0.Policy{}
	for _, pol := range polRes.Policy {
		pols = append(pols, pol)
	}

	return pols, nil

}

func (c *theGraphServiceBetaImpl) ListPolicies(options *TheGraphQueryOptions) ([]*beta_v0.Policy, error) {
	query := fmt.Sprintf(`
	query(%s){
		{
		  policies(%s) {
			creator
			policyId
			policyData {
			  id
			  frameworkId
			  needsActivation
			  mintingParamValues
			  linkParentParamValues
			  activationParamValues
			}
		  }
		}
	}
    `, QUERY_INTERFACE, QUERY_VALUE)

	req := c.buildNewRequest(options, query)

	ctx := context.Background()
	var polsRes beta_v0.PoliciesTheGraphResponse
	if err := c.client.Run(ctx, req, &polsRes); err != nil {
		return nil, fmt.Errorf("failed to get policies from the graph. error: %v", err)
	}

	pols := []*beta_v0.Policy{}
	for _, pol := range polsRes.Policies {
		pols = append(pols, pol)
	}

	return pols, nil
}

func (s *theGraphServiceBetaImpl) buildNewRequest(options *TheGraphQueryOptions, query string) *graphql.Request {
	options = s.setQueryOptions(options)

	req := graphql.NewRequest(query)
	req.Var("first", options.First)
	req.Var("skip", options.Skip)
	req.Var("orderBy", options.OrderBy)
	req.Var("orderDirection", options.OrderDirection)

	return req
}

func (s *theGraphServiceBetaImpl) setQueryOptions(options *TheGraphQueryOptions) *TheGraphQueryOptions {
	if options == nil {
		options = &TheGraphQueryOptions{
			First: 100,
			Skip:  0,
		}
	}

	if options.First == 0 {
		options.First = 100
	}

	options.OrderBy = "blockTimestamp"
	options.OrderDirection = "desc"

	return options
}

//func (c *theGraphServiceBetaImpl) GetIPsRegistered() ([]*models.IPRegistered, error) {
//	req := graphql.NewRequest(`
//    {
//		ipregistereds {
//			id
//			chainId
//			tokenContract
//			tokenId
//			resolver
//		}
//	}`)
//
//	ctx := context.Background()
//	var ipRegisteredsTheGraphResponse models.IPRegisteredTheGraphResponse
//	if err := c.client.Run(ctx, req, &ipRegisteredsTheGraphResponse); err != nil {
//		return nil, fmt.Errorf("failed to get registered ip accounts from the graph. error: %v", err)
//	}
//
//	ips := []*models.IPRegistered{}
//	for _, ip := range ipRegisteredsTheGraphResponse.IPRegistered {
//		ips = append(ips, ip)
//	}
//
//	return ips, nil
//}

//func (c *theGraphServiceBetaImpl) GetSetIPAccounts() ([]*models.SetIPAccount, error) {
//	req := graphql.NewRequest(`
//    {
//		ipaccountSets {
//			ipId
//			chainId
//			tokenContract
//			tokenId
//		}
//	}`)
//
//	ctx := context.Background()
//	var ipAccountSets models.SetIPAccountTheGraphResponse
//	if err := c.client.Run(ctx, req, &ipAccountSets); err != nil {
//		return nil, fmt.Errorf("failed to get set ip accounts from the graph. error: %v", err)
//	}
//
//	accs := []*models.SetIPAccount{}
//	for _, acc := range ipAccountSets.SetIPAccount {
//		accs = append(accs, acc)
//	}
//
//	return accs, nil
//}

//func (c *theGraphServiceBetaImpl) GetSetIPResolvers() ([]*models.SetResolver, error) {
//	req := graphql.NewRequest(`
//    {
//		ipresolverSets {
//			ipId
//			resolver
//		}
//	}`)
//
//	ctx := context.Background()
//	var ipResolverSets models.SetResolverTheGraphResponse
//	if err := c.client.Run(ctx, req, &ipResolverSets); err != nil {
//		return nil, fmt.Errorf("failed to get set ip resolvers from the graph. error: %v", err)
//	}
//
//	rslvrs := []*models.SetResolver{}
//	for _, rslvr := range ipResolverSets.SetResolver {
//		rslvrs = append(rslvrs, rslvr)
//	}
//
//	return rslvrs, nil
//}
