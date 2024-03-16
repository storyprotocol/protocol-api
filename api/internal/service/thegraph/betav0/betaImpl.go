package betav0

import (
	"fmt"
	"github.com/machinebox/graphql"
	"github.com/storyprotocol/protocol-api/api/internal/service/openapi"
	"github.com/storyprotocol/protocol-api/api/internal/service/thegraph"
)

const (
	REN_QUERY_INTERFACE = "$limit: Int, $skip: Int"
	REN_QUERY_VALUE     = "limit: $limit, skip: $skip, orderBy: %s, orderDirection: %s"
	QUERY_INTERFACE     = "$first: Int, $skip: Int, $orderBy: String, $orderDirection: String"
	QUERY_VALUE         = "first: $first, skip: $skip, orderBy: $orderBy, orderDirection: $orderDirection"
	QUERY_PLACEHOLDER   = "string"
	ORDER_BY            = "block_number"
	ORDER_DIRECTION     = "desc"
	RENAISSANCE         = "renaissance"
	SEPOLIA             = "sepolia"
)

func NewTheGraphServiceBetaImpl(
	client *graphql.Client,
	openChainUrl string,
	apiKey string) thegraph.TheGraphServiceBeta {
	openChainClient := openapi.NewOpenChainClient(openChainUrl)
	return &ServiceBetaImpl{
		client:          client,
		openChainClient: openChainClient,
		apiKey:          apiKey,
	}
}

type ServiceBetaImpl struct {
	client          *graphql.Client
	splitClient     *graphql.Client
	openChainClient *openapi.OpenchainClient
	apiKey          string
}

func (s *ServiceBetaImpl) buildNewRequest(options *thegraph.TheGraphQueryOptions, query string) *graphql.Request {
	options = s.setQueryOptions(options)
	req := graphql.NewRequest(query)
	req.Var("skip", options.Skip)
	req.Var("orderBy", options.OrderBy)
	req.Var("orderDirection", options.OrderDirection)

	//whereString := ""

	if s.apiKey != "" {
		req.Var("limit", options.Limit)
		req.Header.Set("X-API-KEY", s.apiKey)
		req.Header.Set("accept", "application/json")
		req.Header.Set("accept", "application/json")
		//req.Var("filter", fmt.Sprintf("{%s}", whereString))

		//if options.Where.IPID != "" {
		//	whereString = whereString + fmt.Sprintf("ip_d: \"%s\"", options.Where.IPID)
		//}
	} else {
		req.Var("first", options.First)
		//req.Var("where", fmt.Sprintf("{%s}", whereString))
		//
		//if options.Where.IPID != "" {
		//	whereString = whereString + fmt.Sprintf("ipId: \"%s\"", options.Where.IPID)
		//}
	}

	return req
}

func (s *ServiceBetaImpl) buildWhereConditions(options *thegraph.TheGraphQueryOptions) string {
	whereString := ""

	if s.apiKey != "" {
		if options.Where.Module != "" && options.Where.Module != QUERY_PLACEHOLDER {
			whereString = whereString + fmt.Sprintf("module:  {eq:\"%s\"},", options.Where.Module)
		}
		if options.Where.Name != "" && options.Where.Name != QUERY_PLACEHOLDER {
			whereString = whereString + fmt.Sprintf("name: {eq:\"%s\"},", options.Where.Name)
		}
		if options.Where.TargetTag != "" && options.Where.TargetTag != QUERY_PLACEHOLDER {
			whereString = whereString + fmt.Sprintf("target_tag: {eq:\"%s\"},", options.Where.TargetTag)
		}
		if options.Where.TargetIpId != "" && options.Where.TargetIpId != QUERY_PLACEHOLDER {
			whereString = whereString + fmt.Sprintf("target_ip_id: {eq:\"%s\"},", options.Where.TargetIpId)
		}
		if options.Where.CurrentTag != "" && options.Where.CurrentTag != QUERY_PLACEHOLDER {
			whereString = whereString + fmt.Sprintf("current_tag: {eq:\"%s\"},", options.Where.CurrentTag)
		}
		if options.Where.Initiator != "" && options.Where.Initiator != QUERY_PLACEHOLDER {
			whereString = whereString + fmt.Sprintf("initiator: {eq:\"%s\"},", options.Where.Initiator)
		}
		if options.Where.MetadataResolverAddress != "" && options.Where.MetadataResolverAddress != QUERY_PLACEHOLDER {
			whereString = whereString + fmt.Sprintf("metadata_resolver_address: {eq:\"%s\"},", options.Where.MetadataResolverAddress)
		}
		if options.Where.TokenContract != "" && options.Where.TokenContract != QUERY_PLACEHOLDER {
			whereString = whereString + fmt.Sprintf("token_contract: {eq:\"%s\"},", options.Where.TokenContract)
		}
		if options.Where.TokenId != "" && options.Where.TokenId != QUERY_PLACEHOLDER {
			whereString = whereString + fmt.Sprintf("token_id: {eq:\"%s\"},", options.Where.TokenId)
		}
		if options.Where.ChainId != "" && options.Where.ChainId != QUERY_PLACEHOLDER {
			whereString = whereString + fmt.Sprintf("chain_id: {eq:\"%s\"},", options.Where.ChainId)
		}
		if options.Where.PolicyId != "" && options.Where.PolicyId != QUERY_PLACEHOLDER {
			whereString = whereString + fmt.Sprintf("policy_id: {eq:\"%s\"},", options.Where.PolicyId)
		}
		if options.Where.Active != "" && options.Where.Active != QUERY_PLACEHOLDER {
			whereString = whereString + fmt.Sprintf("active: %v,", options.Where.Active)
		}
		if options.Where.Inherited != "" && options.Where.Inherited != QUERY_PLACEHOLDER {
			whereString = whereString + fmt.Sprintf("inherited: %v,", options.Where.Inherited)
		}
		if options.Where.LicensorIpdId != "" && options.Where.LicensorIpdId != QUERY_PLACEHOLDER {
			whereString = whereString + fmt.Sprintf("licensor_ip_id: {eq:\"%s\"},", options.Where.LicensorIpdId)
		}
		if options.Where.Creator != "" && options.Where.Creator != QUERY_PLACEHOLDER {
			whereString = whereString + fmt.Sprintf("creator: {eq:\"%s\"},", options.Where.Creator)
		}
		if options.Where.Signer != "" && options.Where.Signer != QUERY_PLACEHOLDER {
			whereString = whereString + fmt.Sprintf("signer: {eq:\"%s\"},", options.Where.Signer)
		}
		if options.Where.To != "" && options.Where.To != QUERY_PLACEHOLDER {
			whereString = whereString + fmt.Sprintf("to: {eq:\"%s\"},", options.Where.To)
		}
		if options.Where.Address != "" && options.Where.Address != QUERY_PLACEHOLDER {
			whereString = whereString + fmt.Sprintf("address: {eq:\"%s\"},", options.Where.Address)
		}
		if options.Where.IPID != "" && options.Where.IPID != QUERY_PLACEHOLDER {
			whereString = whereString + fmt.Sprintf("ip_id: {eq:\"%s\"},", options.Where.IPID)
		}
		if options.Where.RoyaltyPolicy != "" && options.Where.RoyaltyPolicy != QUERY_PLACEHOLDER {
			whereString = whereString + fmt.Sprintf("royalty_policy: {eq:\"%s\"},", options.Where.RoyaltyPolicy)
		}
		if options.Where.ReceiverIpId != "" && options.Where.ReceiverIpId != QUERY_PLACEHOLDER {
			whereString = whereString + fmt.Sprintf("receiver_ip_id: {eq:\"%s\"},", options.Where.ReceiverIpId)
		}
		if options.Where.PayerIpId != "" && options.Where.PayerIpId != QUERY_PLACEHOLDER {
			whereString = whereString + fmt.Sprintf("payer_ip_id: {eq:\"%s\"},", options.Where.PayerIpId)
		}
		if options.Where.Sender != "" && options.Where.Sender != QUERY_PLACEHOLDER {
			whereString = whereString + fmt.Sprintf("sender: {eq:\"%s\"},", options.Where.Sender)
		}
		if options.Where.Payer != "" && options.Where.Payer != QUERY_PLACEHOLDER {
			whereString = whereString + fmt.Sprintf("payer: {eq:\"%s\"},", options.Where.Payer)
		}
		if options.Where.Token != "" && options.Where.Token != QUERY_PLACEHOLDER {
			whereString = whereString + fmt.Sprintf("token: {eq:\"%s\"},", options.Where.Token)
		}
		if options.Where.Tag != "" && options.Where.Tag != QUERY_PLACEHOLDER {
			whereString = whereString + fmt.Sprintf("tag: {eq:\"%s\"},", options.Where.Tag)
		}
		if options.Where.ActionType != "" && options.Where.ActionType != QUERY_PLACEHOLDER {
			whereString = whereString + fmt.Sprintf("action_type: {eq:\"%s\"},", options.Where.ActionType)
		}
		if options.Where.ResourceId != "" && options.Where.ResourceId != QUERY_PLACEHOLDER {
			whereString = whereString + fmt.Sprintf("resource_id: {eq:\"%s\"},", options.Where.ResourceId)
		}
		if options.Where.PolicyFrameworkManager != "" && options.Where.PolicyFrameworkManager != QUERY_PLACEHOLDER {
			whereString = whereString + fmt.Sprintf("policy_framework_manager: {eq:\"%s\"},", options.Where.PolicyFrameworkManager)
		}
		if options.Where.MintingFeeToken != "" && options.Where.MintingFeeToken != QUERY_PLACEHOLDER {
			whereString = whereString + fmt.Sprintf("minting_fee_token: {eq:\"%s\"},", options.Where.MintingFeeToken)
		}
		if options.Where.Owner != "" && options.Where.Owner != QUERY_PLACEHOLDER {
			whereString = whereString + fmt.Sprintf("owner: {eq:\"%s\"},", options.Where.Owner)
		}
	} else {
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
			whereString = whereString + fmt.Sprintf("active: %v,", options.Where.Active)
		}
		if options.Where.Inherited != "" && options.Where.Inherited != QUERY_PLACEHOLDER {
			whereString = whereString + fmt.Sprintf("inherited: %v,", options.Where.Inherited)
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
		if options.Where.Payer != "" && options.Where.Payer != QUERY_PLACEHOLDER {
			whereString = whereString + fmt.Sprintf("payer: \"%s\",", options.Where.Payer)
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
		if options.Where.MintingFeeToken != "" && options.Where.MintingFeeToken != QUERY_PLACEHOLDER {
			whereString = whereString + fmt.Sprintf("mintingFeeToken: \"%s\",", options.Where.MintingFeeToken)
		}
		if options.Where.Owner != "" && options.Where.Owner != QUERY_PLACEHOLDER {
			whereString = whereString + fmt.Sprintf("owner: \"%s\",", options.Where.Owner)
		}
	}

	return whereString
}

func (s *ServiceBetaImpl) setQueryOptions(options *thegraph.TheGraphQueryOptions) *thegraph.TheGraphQueryOptions {
	if options == nil {
		options = &thegraph.TheGraphQueryOptions{
			First: 100,
			Limit: 100,
			Skip:  0,
		}
	}

	if options.First == 0 {
		options.First = 100
	}

	if options.Limit == 0 {
		options.Limit = 100
	}

	if options.OrderBy == "" {
		if s.apiKey != "" {
			options.OrderBy = ORDER_BY
		} else {
			options.OrderBy = "blockNumber"
		}
	}

	if options.OrderDirection == "" {
		options.OrderDirection = ORDER_DIRECTION
	}

	return options
}
