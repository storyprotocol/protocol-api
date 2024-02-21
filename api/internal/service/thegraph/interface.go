package thegraph

import (
	"github.com/storyprotocol/protocol-api/api/internal/models/betav0"
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
	GetRoyaltyPolicy(royaltyPolicyId string) (*betav0.RoyaltyPolicy, error)
	GetLicenseMintingFeePaid(licenseMintingFeePaidId string) (*betav0.LicenseMintingFeePaid, error)
	GetTransaction(trxId string) (*betav0.Transaction, error)

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
	ListRoyaltyPolicies(options *TheGraphQueryOptions) ([]*betav0.RoyaltyPolicy, error)
	ListLicenseMintingFeePaids(options *TheGraphQueryOptions) ([]*betav0.LicenseMintingFeePaid, error)
	ListTransactions(options *TheGraphQueryOptions) ([]*betav0.Transaction, error)
}

type TheGraphQueryOptions struct {
	First          int
	Skip           int
	OrderBy        string
	OrderDirection string
	Where          struct {
		Name                    string `json:"name,omitempty"`
		Module                  string `json:"module,omitempty"`
		TargetIpId              string `json:"targetIpId,omitempty"`
		TargetTag               string `json:"targetTag,omitempty"`
		CurrentTag              string `json:"currentTag,omitempty"`
		Initiator               string `json:"initiator,omitempty"`
		MetadataResolverAddress string `json:"metadataResolverAddress,omitempty"`
		TokenContract           string `json:"tokenContract,omitempty"`
		TokenId                 string `json:"tokenId,omitempty"`
		ChainId                 string `json:"chainId,omitempty"`
		PolicyId                string `json:"policyId,omitempty"`
		Active                  string `json:"active,omitempty"`
		Inherited               string `json:"inherited,omitempty"`
		LicensorIpdId           string `json:"licensorIpdId,omitempty"`
		Creator                 string `json:"creator,omitempty"`
		Signer                  string `json:"signer,omitempty"`
		To                      string `json:"to,omitempty"`
		Address                 string `json:"address,omitempty"`
		IPID                    string `json:"ipId,omitempty"`
		RoyaltyPolicy           string `json:"royaltyPolicy,omitempty"`
		ReceiverIpId            string `json:"receiverIpId,omitempty"`
		PayerIpId               string `json:"payerIpId,omitempty"`
		Sender                  string `json:"sender,omitempty"`
		Token                   string `json:"token,omitempty"`
		Payer                   string `json:"payer,omitempty"`
		Tag                     string `json:"tag,omitempty"`
		ActionType              string `json:"actionType,omitempty"`
		ResourceId              string `json:"resourceId,omitempty"`
		PolicyFrameworkManager  string `json:"policyFrameworkManager,omitempty"`
		MintingFeeToken         string `json:"mintingFeeToken,omitempty"`
	}
}
