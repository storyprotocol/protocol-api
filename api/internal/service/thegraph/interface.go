package thegraph

import (
	"github.com/storyprotocol/protocol-api/api/internal/models"
	"github.com/storyprotocol/protocol-api/api/internal/models/beta-v0"
)

type TheGraphServiceBeta interface {
	GetIPAccount(accountId string) ([]*beta_v0.IPAccount, error)
	GetModule(moduleName string) ([]*beta_v0.Module, error)
	GetLicense(licenseId string) ([]*beta_v0.License, error)
	GetLicenseFramework(licenseId string) ([]*beta_v0.LicenseFramework, error)
	GetPolicy(policyId string) ([]*beta_v0.Policy, error)

	ListIPAccounts(options *TheGraphQueryOptions) ([]*beta_v0.IPAccount, error)
	ListModules(options *TheGraphQueryOptions) ([]*beta_v0.Module, error)
	ListLicenses(options *TheGraphQueryOptions) ([]*beta_v0.License, error)
	ListLicenseFrameworks(options *TheGraphQueryOptions) ([]*beta_v0.LicenseFramework, error)
	ListAccessControlPermissions(options *TheGraphQueryOptions) ([]*beta_v0.AccessControlPermission, error)
	ListPolicies(options *TheGraphQueryOptions) ([]*beta_v0.Policy, error)
}

type TheGraphQueryOptions struct {
	First          int
	Skip           int
	OrderBy        string
	OrderDirection string
}

func FromRequestQueryOptions(options *models.QueryOptions) *TheGraphQueryOptions {
	if options == nil {
		return &TheGraphQueryOptions{
			First: 100,
			Skip:  0,
		}
	}

	if options.Pagination.Limit == 0 {
		options.Pagination.Limit = 100
	}

	return &TheGraphQueryOptions{
		First: options.Pagination.Limit,
		Skip:  options.Pagination.Offset,
	}
}
