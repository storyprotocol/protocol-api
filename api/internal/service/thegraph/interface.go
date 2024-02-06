package thegraph

import (
	"github.com/storyprotocol/protocol-api/api/internal/models/beta-v0"
	"github.com/storyprotocol/protocol-api/api/internal/models/beta-v0/options"
)

type TheGraphServiceBeta interface {
	// GET
	GetIPAsset(assetId string) (*beta_v0.IPAsset, error)
	GetModule(moduleId string) (*beta_v0.Module, error)
	GetLicense(licenseId string) (*beta_v0.License, error)
	GetLicenseFramework(licenseId string) (*beta_v0.LicenseFramework, error)
	GetPolicy(policyId string) (*beta_v0.Policy, error)
	GetDispute(disputeId string) (*beta_v0.Dispute, error)
	GetPermission(permissionId string) (*beta_v0.Permission, error)
	GetTag(tagId string) (*beta_v0.Tag, error)
	GetRoyalty(royaltyId string) (*beta_v0.Royalty, error)
	GetRoyaltyPay(royaltyPayId string) (*beta_v0.RoyaltyPay, error)
	GetPolicyFrameworkManager(pfwmId string) (*beta_v0.PolicyFrameworkManager, error)

	// LISTS
	ListIPAssets(options *TheGraphQueryOptions) ([]*beta_v0.IPAsset, error)
	ListModules(options *TheGraphQueryOptions) ([]*beta_v0.Module, error)
	ListLicenses(options *TheGraphQueryOptions) ([]*beta_v0.License, error)
	ListLicenseFrameworks(options *TheGraphQueryOptions) ([]*beta_v0.LicenseFramework, error)
	ListAccessControlPermissions(options *TheGraphQueryOptions) ([]*beta_v0.AccessControlPermission, error)
	ListPolicies(options *TheGraphQueryOptions) ([]*beta_v0.Policy, error)
	ListDisputes(options *TheGraphQueryOptions) ([]*beta_v0.Dispute, error)
	ListPermissions(options *TheGraphQueryOptions) ([]*beta_v0.Permission, error)
	ListTag(options *TheGraphQueryOptions) ([]*beta_v0.Tag, error)
	ListRoyalties(options *TheGraphQueryOptions) ([]*beta_v0.Royalty, error)
	ListRoyaltyPays(options *TheGraphQueryOptions) ([]*beta_v0.RoyaltyPay, error)
	ListPolicyFrameworkManagers(options *TheGraphQueryOptions) ([]*beta_v0.PolicyFrameworkManager, error)
}

type TheGraphQueryOptions struct {
	First          int
	Skip           int
	OrderBy        string
	OrderDirection string
	Where          struct {
		Creator       string
		Receiver      string
		TokenContract string
		FrameworkId   string
		IPAsset       string
		IPID          string
	}
}

func FromRequestQueryOptions(options *options.QueryOptions) *TheGraphQueryOptions {
	if options == nil {
		return &TheGraphQueryOptions{
			First: 100,
			Skip:  0,
		}
	}

	var queryOptions = &TheGraphQueryOptions{}

	if options.Pagination.Limit == 0 {
		options.Pagination.Limit = 100
	}

	queryOptions.First = options.Pagination.Limit
	queryOptions.Skip = options.Pagination.Offset

	queryOptions.Where = struct {
		Creator       string
		Receiver      string
		TokenContract string
		FrameworkId   string
		IPAsset       string
		IPID          string
	}(options.Where)

	return queryOptions
}
