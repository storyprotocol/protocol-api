package thegraph

import (
	"context"
	"fmt"
	"github.com/machinebox/graphql"
	"github.com/storyprotocol/protocol-api/api/internal/entity"
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

func (c *theGraphServiceBetaImpl) GetIPAccount(accountId string) ([]*entity.IPAccount, error) {
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
	var ipAccountTheGraphResponse entity.IPAccountTheGraphResponse
	if err := c.client.Run(ctx, req, &ipAccountTheGraphResponse); err != nil {
		return nil, fmt.Errorf("failed to get account from the graph. error: %v", err)
	}

	accts := []*entity.IPAccount{}
	accts = append(accts, ipAccountTheGraphResponse.IPAccount)

	return accts, nil
}

func (c *theGraphServiceBetaImpl) ListIPAccounts(options *TheGraphQueryOptions) ([]*entity.IPAccount, error) {
	options = c.setQueryOptions(options)

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

	req := graphql.NewRequest(query)
	req.Var("first", options.First)
	req.Var("skip", options.Skip)
	req.Var("orderBy", options.OrderBy)
	req.Var("orderDirection", options.OrderDirection)

	ctx := context.Background()
	var ipAccountsTheGraphResponse entity.IPAccountsTheGraphResponse
	if err := c.client.Run(ctx, req, &ipAccountsTheGraphResponse); err != nil {
		return nil, fmt.Errorf("failed to get registered ip accounts from the graph. error: %v", err)
	}

	ipAccounts := []*entity.IPAccount{}
	for _, ipAccount := range ipAccountsTheGraphResponse.IPAccounts {
		ipAccounts = append(ipAccounts, ipAccount)
	}

	return ipAccounts, nil
}

func (c *theGraphServiceBetaImpl) GetModule(moduleName string) ([]*entity.Module, error) {
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
	var modules entity.ModuleTheGraphResponse
	if err := c.client.Run(ctx, req, &modules); err != nil {
		return nil, fmt.Errorf("failed to get modules from the graph. error: %v", err)
	}

	mods := []*entity.Module{}
	for _, mod := range modules.Module {
		mods = append(mods, mod)
	}

	return mods, nil

}

func (c *theGraphServiceBetaImpl) ListModules(options *TheGraphQueryOptions) ([]*entity.Module, error) {
	options = c.setQueryOptions(options)

	query := fmt.Sprintf(`
	query(%s){
		modules (%s) {
			name
			module
		}
	}
    `, QUERY_INTERFACE, QUERY_VALUE)

	req := graphql.NewRequest(query)
	req.Var("first", options.First)
	req.Var("skip", options.Skip)
	req.Var("orderBy", options.OrderBy)
	req.Var("orderDirection", options.OrderDirection)

	ctx := context.Background()
	var modules entity.ModulesTheGraphResponse
	if err := c.client.Run(ctx, req, &modules); err != nil {
		return nil, fmt.Errorf("failed to get modules from the graph. error: %v", err)
	}

	mods := []*entity.Module{}
	for _, mod := range modules.Modules {
		mods = append(mods, mod)
	}

	return mods, nil
}

func (c *theGraphServiceBetaImpl) GetLicense(licenseId string) ([]*entity.License, error) {
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
	var licensesRes entity.LicenseTheGraphResponse
	if err := c.client.Run(ctx, req, &licensesRes); err != nil {
		return nil, fmt.Errorf("failed to get license from the graph. error: %v", err)
	}

	licenses := []*entity.License{}
	for _, license := range licensesRes.License {
		licenses = append(licenses, license)
	}

	return licenses, nil

}

func (c *theGraphServiceBetaImpl) ListLicenses(options *TheGraphQueryOptions) ([]*entity.License, error) {
	options = c.setQueryOptions(options)

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

	req := graphql.NewRequest(query)
	req.Var("first", options.First)
	req.Var("skip", options.Skip)
	req.Var("orderBy", options.OrderBy)
	req.Var("orderDirection", options.OrderDirection)

	ctx := context.Background()
	var licensesRes entity.LicensesTheGraphResponse
	if err := c.client.Run(ctx, req, &licensesRes); err != nil {
		return nil, fmt.Errorf("failed to get licenses from the graph. error: %v", err)
	}

	licenses := []*entity.License{}
	for _, license := range licensesRes.Licenses {
		licenses = append(licenses, license)
	}

	return licenses, nil
}

func (c *theGraphServiceBetaImpl) GetLicenseFramework(licenseId string) ([]*entity.LicenseFramework, error) {
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
	var licensesRes entity.LicenseFrameworkTheGraphResponse
	if err := c.client.Run(ctx, req, &licensesRes); err != nil {
		return nil, fmt.Errorf("failed to get license from the graph. error: %v", err)
	}

	licenses := []*entity.LicenseFramework{}
	for _, license := range licensesRes.LicenseFramework {
		licenses = append(licenses, license)
	}

	return licenses, nil

}

func (c *theGraphServiceBetaImpl) ListLicenseFrameworks(options *TheGraphQueryOptions) ([]*entity.LicenseFramework, error) {
	options = c.setQueryOptions(options)

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

	req := graphql.NewRequest(query)
	req.Var("first", options.First)
	req.Var("skip", options.Skip)
	req.Var("orderBy", options.OrderBy)
	req.Var("orderDirection", options.OrderDirection)

	ctx := context.Background()
	var licensesRes entity.LicenseFrameworksTheGraphResponse
	if err := c.client.Run(ctx, req, &licensesRes); err != nil {
		return nil, fmt.Errorf("failed to get license frameworks from the graph. error: %v", err)
	}

	licenses := []*entity.LicenseFramework{}
	for _, license := range licensesRes.LicenseFrameworks {
		licenses = append(licenses, license)
	}

	return licenses, nil
}

func (c *theGraphServiceBetaImpl) ListAccessControlPermissions(options *TheGraphQueryOptions) ([]*entity.AccessControlPermission, error) {
	options = c.setQueryOptions(options)

	query := fmt.Sprintf(`
	query(%s){
		modules (%s) {
			name
			module
		}
	}
    `, QUERY_INTERFACE, QUERY_VALUE)

	req := graphql.NewRequest(query)
	req.Var("first", options.First)
	req.Var("skip", options.Skip)
	req.Var("orderBy", options.OrderBy)
	req.Var("orderDirection", options.OrderDirection)

	ctx := context.Background()
	var acpsRes entity.AccessControlPermissionTheGraphResponse
	if err := c.client.Run(ctx, req, &acpsRes); err != nil {
		return nil, fmt.Errorf("failed to get access control permissions from the graph. error: %v", err)
	}

	acps := []*entity.AccessControlPermission{}
	for _, acp := range acpsRes.AccessControlPermissions {
		acps = append(acps, acp)
	}

	return acps, nil
}

func (c *theGraphServiceBetaImpl) GetPolicy(policyId string) ([]*entity.Policy, error) {
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
	var polRes entity.PolicyTheGraphResponse
	if err := c.client.Run(ctx, req, &polRes); err != nil {
		return nil, fmt.Errorf("failed to get policy from the graph. error: %v", err)
	}

	pols := []*entity.Policy{}
	for _, pol := range polRes.Policy {
		pols = append(pols, pol)
	}

	return pols, nil

}

func (c *theGraphServiceBetaImpl) ListPolicies(options *TheGraphQueryOptions) ([]*entity.Policy, error) {
	options = c.setQueryOptions(options)

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

	req := graphql.NewRequest(query)
	req.Var("first", options.First)
	req.Var("skip", options.Skip)
	req.Var("orderBy", options.OrderBy)
	req.Var("orderDirection", options.OrderDirection)

	ctx := context.Background()
	var polsRes entity.PoliciesTheGraphResponse
	if err := c.client.Run(ctx, req, &polsRes); err != nil {
		return nil, fmt.Errorf("failed to get policies from the graph. error: %v", err)
	}

	pols := []*entity.Policy{}
	for _, pol := range polsRes.Policies {
		pols = append(pols, pol)
	}

	return pols, nil
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

//func (c *theGraphServiceBetaImpl) GetIPsRegistered() ([]*entity.IPRegistered, error) {
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
//	var ipRegisteredsTheGraphResponse entity.IPRegisteredTheGraphResponse
//	if err := c.client.Run(ctx, req, &ipRegisteredsTheGraphResponse); err != nil {
//		return nil, fmt.Errorf("failed to get registered ip accounts from the graph. error: %v", err)
//	}
//
//	ips := []*entity.IPRegistered{}
//	for _, ip := range ipRegisteredsTheGraphResponse.IPRegistered {
//		ips = append(ips, ip)
//	}
//
//	return ips, nil
//}

//func (c *theGraphServiceBetaImpl) GetSetIPAccounts() ([]*entity.SetIPAccount, error) {
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
//	var ipAccountSets entity.SetIPAccountTheGraphResponse
//	if err := c.client.Run(ctx, req, &ipAccountSets); err != nil {
//		return nil, fmt.Errorf("failed to get set ip accounts from the graph. error: %v", err)
//	}
//
//	accs := []*entity.SetIPAccount{}
//	for _, acc := range ipAccountSets.SetIPAccount {
//		accs = append(accs, acc)
//	}
//
//	return accs, nil
//}

//func (c *theGraphServiceBetaImpl) GetSetIPResolvers() ([]*entity.SetResolver, error) {
//	req := graphql.NewRequest(`
//    {
//		ipresolverSets {
//			ipId
//			resolver
//		}
//	}`)
//
//	ctx := context.Background()
//	var ipResolverSets entity.SetResolverTheGraphResponse
//	if err := c.client.Run(ctx, req, &ipResolverSets); err != nil {
//		return nil, fmt.Errorf("failed to get set ip resolvers from the graph. error: %v", err)
//	}
//
//	rslvrs := []*entity.SetResolver{}
//	for _, rslvr := range ipResolverSets.SetResolver {
//		rslvrs = append(rslvrs, rslvr)
//	}
//
//	return rslvrs, nil
//}
