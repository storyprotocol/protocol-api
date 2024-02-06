package beta_v0

import (
	"fmt"
	"github.com/machinebox/graphql"
	"github.com/storyprotocol/protocol-api/api/internal/service/thegraph"
)

const (
	QUERY_INTERFACE   = "$first: Int, $skip: Int, $orderBy: String, $orderDirection: String"
	QUERY_VALUE       = "first: $first, skip: $skip, orderBy: $orderBy, orderDirection: $orderDirection"
	QUERY_PLACEHOLDER = "string"
)

func NewTheGraphServiceBetaImpl(client *graphql.Client) thegraph.TheGraphServiceBeta {
	return &ServiceBetaImpl{
		client: client,
	}
}

type ServiceBetaImpl struct {
	client *graphql.Client
}

func (s *ServiceBetaImpl) buildNewRequest(options *thegraph.TheGraphQueryOptions, query string) *graphql.Request {
	options = s.setQueryOptions(options)

	req := graphql.NewRequest(query)
	req.Var("first", options.First)
	req.Var("skip", options.Skip)
	req.Var("orderBy", options.OrderBy)
	req.Var("orderDirection", options.OrderDirection)

	whereString := ""
	if options.Where.IPID != "" {
		whereString = whereString + fmt.Sprintf("ipId: \"%s\"", options.Where.IPID)
	}

	req.Var("where", fmt.Sprintf("{%s}", whereString))

	return req
}

func (s *ServiceBetaImpl) buildWhereConditions(options *thegraph.TheGraphQueryOptions) string {
	whereString := ""
	if options.Where.IPID != "" && options.Where.IPID != QUERY_PLACEHOLDER {
		whereString = whereString + fmt.Sprintf("ipId: \"%s\",", options.Where.IPID)
	}
	if options.Where.IPAsset != "" && options.Where.IPAsset != QUERY_PLACEHOLDER {
		whereString = whereString + fmt.Sprintf("ipAsset: \"%s\",", options.Where.IPAsset)
	}
	if options.Where.TokenContract != "" && options.Where.TokenContract != QUERY_PLACEHOLDER {
		whereString = whereString + fmt.Sprintf("tokenContract: \"%s\",", options.Where.TokenContract)
	}
	if options.Where.Creator != "" && options.Where.Creator != QUERY_PLACEHOLDER {
		whereString = whereString + fmt.Sprintf("creator: \"%s\",", options.Where.Creator)
	}
	if options.Where.FrameworkId != "" && options.Where.FrameworkId != QUERY_PLACEHOLDER {
		whereString = whereString + fmt.Sprintf("frameworkId: \"%s\",", options.Where.FrameworkId)
	}
	if options.Where.Receiver != "" && options.Where.Receiver != QUERY_PLACEHOLDER {
		whereString = whereString + fmt.Sprintf("receiver: \"%s\",", options.Where.Receiver)
	}

	return whereString
}

func (s *ServiceBetaImpl) setQueryOptions(options *thegraph.TheGraphQueryOptions) *thegraph.TheGraphQueryOptions {
	if options == nil {
		options = &thegraph.TheGraphQueryOptions{
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
