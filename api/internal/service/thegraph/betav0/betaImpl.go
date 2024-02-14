package betav0

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
	if options.Where.Module != "" && options.Where.Module != QUERY_PLACEHOLDER {
		whereString = whereString + fmt.Sprintf("module: \"%s\",", options.Where.Module)
	}
	if options.Where.Name != "" && options.Where.Name != QUERY_PLACEHOLDER {
		whereString = whereString + fmt.Sprintf("name: \"%s\",", options.Where.Name)
	}
	if options.Where.TargetTag != "" && options.Where.TargetTag != QUERY_PLACEHOLDER {
		whereString = whereString + fmt.Sprintf("targetTag: \"%s\",", options.Where.TargetTag)
	}
	if options.Where.TargetIpId != "" && options.Where.TargetIpId != QUERY_PLACEHOLDER {
		whereString = whereString + fmt.Sprintf("targetIpId: \"%s\",", options.Where.TargetIpId)
	}
	if options.Where.CurrentTag != "" && options.Where.CurrentTag != QUERY_PLACEHOLDER {
		whereString = whereString + fmt.Sprintf("currentTag: \"%s\",", options.Where.CurrentTag)
	}
	if options.Where.Initiator != "" && options.Where.Initiator != QUERY_PLACEHOLDER {
		whereString = whereString + fmt.Sprintf("initiator: \"%s\",", options.Where.Initiator)
	}
	if options.Where.MetadataResolverAddress != "" && options.Where.MetadataResolverAddress != QUERY_PLACEHOLDER {
		whereString = whereString + fmt.Sprintf("metadataResolverAddress: \"%s\",", options.Where.MetadataResolverAddress)
	}
	if options.Where.TokenContract != "" && options.Where.TokenContract != QUERY_PLACEHOLDER {
		whereString = whereString + fmt.Sprintf("tokenContract: \"%s\",", options.Where.TokenContract)
	}
	if options.Where.TokenId != "" && options.Where.TokenId != QUERY_PLACEHOLDER {
		whereString = whereString + fmt.Sprintf("tokenId: \"%s\",", options.Where.TokenId)
	}
	if options.Where.ChainId != "" && options.Where.ChainId != QUERY_PLACEHOLDER {
		whereString = whereString + fmt.Sprintf("chainId: \"%s\",", options.Where.ChainId)
	}
	if options.Where.PolicyId != "" && options.Where.PolicyId != QUERY_PLACEHOLDER {
		whereString = whereString + fmt.Sprintf("policyId: \"%s\",", options.Where.PolicyId)
	}
	if options.Where.Active != "" && options.Where.Active != QUERY_PLACEHOLDER {
		whereString = whereString + fmt.Sprintf("active: \"%s\",", options.Where.Active)
	}
	if options.Where.Inherited != "" && options.Where.Inherited != QUERY_PLACEHOLDER {
		whereString = whereString + fmt.Sprintf("inherited: \"%s\",", options.Where.Inherited)
	}
	if options.Where.LicensorIpdId != "" && options.Where.LicensorIpdId != QUERY_PLACEHOLDER {
		whereString = whereString + fmt.Sprintf("licensorIpId: \"%s\",", options.Where.LicensorIpdId)
	}
	if options.Where.Creator != "" && options.Where.Creator != QUERY_PLACEHOLDER {
		whereString = whereString + fmt.Sprintf("creator: \"%s\",", options.Where.Creator)
	}
	if options.Where.Signer != "" && options.Where.Signer != QUERY_PLACEHOLDER {
		whereString = whereString + fmt.Sprintf("signer: \"%s\",", options.Where.Signer)
	}
	if options.Where.To != "" && options.Where.To != QUERY_PLACEHOLDER {
		whereString = whereString + fmt.Sprintf("to: \"%s\",", options.Where.To)
	}
	if options.Where.Address != "" && options.Where.Address != QUERY_PLACEHOLDER {
		whereString = whereString + fmt.Sprintf("address: \"%s\",", options.Where.Address)
	}
	if options.Where.IPID != "" && options.Where.IPID != QUERY_PLACEHOLDER {
		whereString = whereString + fmt.Sprintf("ipId: \"%s\",", options.Where.IPID)
	}
	if options.Where.RoyaltyPolicy != "" && options.Where.RoyaltyPolicy != QUERY_PLACEHOLDER {
		whereString = whereString + fmt.Sprintf("royaltyPolicy: \"%s\",", options.Where.RoyaltyPolicy)
	}
	if options.Where.ReceiverIpId != "" && options.Where.ReceiverIpId != QUERY_PLACEHOLDER {
		whereString = whereString + fmt.Sprintf("receiverIpId: \"%s\",", options.Where.ReceiverIpId)
	}
	if options.Where.PayerIpId != "" && options.Where.PayerIpId != QUERY_PLACEHOLDER {
		whereString = whereString + fmt.Sprintf("payerIpId: \"%s\",", options.Where.PayerIpId)
	}
	if options.Where.Sender != "" && options.Where.Sender != QUERY_PLACEHOLDER {
		whereString = whereString + fmt.Sprintf("sender: \"%s\",", options.Where.Sender)
	}
	if options.Where.Token != "" && options.Where.Token != QUERY_PLACEHOLDER {
		whereString = whereString + fmt.Sprintf("token: \"%s\",", options.Where.Token)
	}
	if options.Where.Tag != "" && options.Where.Tag != QUERY_PLACEHOLDER {
		whereString = whereString + fmt.Sprintf("tag: \"%s\",", options.Where.Tag)
	}
	if options.Where.ActionType != "" && options.Where.ActionType != QUERY_PLACEHOLDER {
		whereString = whereString + fmt.Sprintf("actionType: \"%s\",", options.Where.ActionType)
	}
	if options.Where.ResourceId != "" && options.Where.ResourceId != QUERY_PLACEHOLDER {
		whereString = whereString + fmt.Sprintf("resourceId: \"%s\",", options.Where.ResourceId)
	}
	if options.Where.PolicyFrameworkManager != "" && options.Where.PolicyFrameworkManager != QUERY_PLACEHOLDER {
		whereString = whereString + fmt.Sprintf("policyFrameworkManager: \"%s\",", options.Where.PolicyFrameworkManager)
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
