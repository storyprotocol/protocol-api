package thegraph

import "github.com/storyprotocol/protocol-api/api/internal/entity"

type TheGraphServiceBeta interface {
	GetIPAccount(accountId string) ([]*entity.IPAccount, error)
	GetModule(moduleName string) ([]*entity.Module, error)

	ListIPAccounts(options *TheGraphQueryOptions) ([]*entity.IPAccount, error)
	ListModules(options *TheGraphQueryOptions) ([]*entity.Module, error)
}

type TheGraphQueryOptions struct {
	First          int
	Skip           int
	OrderBy        string
	OrderDirection string
}

func FromRequestQueryOptions(options *entity.QueryOptions) *TheGraphQueryOptions {
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
