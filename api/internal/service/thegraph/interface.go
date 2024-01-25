package thegraph

import (
	"github.com/storyprotocol/protocol-api/api/internal/models/beta-v0"
	"github.com/storyprotocol/protocol-api/api/internal/models/beta-v0/options"
)

type TheGraphServiceBeta interface {
	// GET
	GetIPAccount(accountId string) (*beta_v0.IPAccount, error)
	GetModule(moduleId string) (*beta_v0.Module, error)
	GetLicense(licenseId string) (*beta_v0.License, error)
	GetLicenseFramework(licenseId string) (*beta_v0.LicenseFramework, error)
	GetTag(tagId string) (*beta_v0.Tag, error)
	GetPolicy(policyId string) (*beta_v0.Policy, error)
	GetDispute(disputeId string) (*beta_v0.Dispute, error)
	GetPermission(permissionId string) (*beta_v0.Permission, error)

	// LISTS
	ListIPAccounts(options *TheGraphQueryOptions) ([]*beta_v0.IPAccount, error)
	ListModules(options *TheGraphQueryOptions) ([]*beta_v0.Module, error)
	ListLicenses(options *TheGraphQueryOptions) ([]*beta_v0.License, error)
	ListLicenseFrameworks(options *TheGraphQueryOptions) ([]*beta_v0.LicenseFramework, error)
	ListAccessControlPermissions(options *TheGraphQueryOptions) ([]*beta_v0.AccessControlPermission, error)
	ListPolicies(options *TheGraphQueryOptions) ([]*beta_v0.Policy, error)
	ListDisputes(options *TheGraphQueryOptions) ([]*beta_v0.Dispute, error)
	ListTag(options *TheGraphQueryOptions) ([]*beta_v0.Tag, error)
	ListPermissions(options *TheGraphQueryOptions) ([]*beta_v0.Permission, error)
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
		IPAccount     string
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
		IPAccount     string
		IPID          string
	}(options.Where)

	return queryOptions
}
