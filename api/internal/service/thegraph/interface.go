package thegraph

import (
	"github.com/storyprotocol/protocol-api/api/internal/models/betav0"
	"github.com/storyprotocol/protocol-api/api/internal/models/betav0/options"
)

type TheGraphServiceBeta interface {
	// GET
	GetIPAsset(assetId string) (*betav0.IPAsset, error)
	GetModule(moduleId string) (*betav0.Module, error)
	GetLicense(licenseId string) (*betav0.License, error)
	GetLicenseFramework(licenseId string) (*betav0.LicenseFramework, error)
	GetPolicy(policyId string) (*betav0.Policy, error)
	GetIPAPolicy(ipaPolicyId string) (*betav0.IPAPolicy, error)
	GetDispute(disputeId string) (*betav0.Dispute, error)
	GetPermission(permissionId string) (*betav0.Permission, error)
	GetTag(tagId string) (*betav0.Tag, error)
	GetRoyalty(royaltyId string) (*betav0.Royalty, error)
	GetRoyaltyPay(royaltyPayId string) (*betav0.RoyaltyPay, error)
	GetPolicyFrameworkManager(pfwmId string) (*betav0.PolicyFrameworkManager, error)
	GetCollection(colId string) (*betav0.Collection, error)

	// LISTS
	ListIPAssets(options *TheGraphQueryOptions) ([]*betav0.IPAsset, error)
	ListModules(options *TheGraphQueryOptions) ([]*betav0.Module, error)
	ListLicenses(options *TheGraphQueryOptions) ([]*betav0.License, error)
	ListLicenseFrameworks(options *TheGraphQueryOptions) ([]*betav0.LicenseFramework, error)
	ListAccessControlPermissions(options *TheGraphQueryOptions) ([]*betav0.AccessControlPermission, error)
	ListPolicies(options *TheGraphQueryOptions) ([]*betav0.Policy, error)
	ListIPAPolicies(options *TheGraphQueryOptions) ([]*betav0.IPAPolicy, error)
	ListDisputes(options *TheGraphQueryOptions) ([]*betav0.Dispute, error)
	ListPermissions(options *TheGraphQueryOptions) ([]*betav0.Permission, error)
	ListTag(options *TheGraphQueryOptions) ([]*betav0.Tag, error)
	ListRoyalties(options *TheGraphQueryOptions) ([]*betav0.Royalty, error)
	ListRoyaltyPays(options *TheGraphQueryOptions) ([]*betav0.RoyaltyPay, error)
	ListPolicyFrameworkManagers(options *TheGraphQueryOptions) ([]*betav0.PolicyFrameworkManager, error)
	ListCollections(options *TheGraphQueryOptions) ([]*betav0.Collection, error)
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
		LicensorIpId  string
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
		LicensorIpId  string
		IPID          string
	}(options.Where)

	return queryOptions
}
