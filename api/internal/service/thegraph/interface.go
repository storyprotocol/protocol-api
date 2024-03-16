package thegraph

import (
	"github.com/storyprotocol/protocol-api/api/internal/models/betav0"
)

type TheGraphServiceBeta interface {
	// GET
	GetIPAsset(assetId string) (*betav0.IPAsset, *betav0.RenIPAsset, error)
	GetModule(moduleId string) (*betav0.Module, *betav0.RenModule, error)
	GetLicense(licenseId string) (*betav0.License, *betav0.RenLicense, error)
	GetLicenseFramework(licenseId string) (*betav0.LicenseFramework, *betav0.RenLicenseFramework, error)
	GetPolicy(policyId string) (*betav0.Policy, *betav0.RenPolicy, error)
	GetIPAPolicy(ipaPolicyId string) (*betav0.IPAPolicy, *betav0.RenIPAPolicy, error)
	GetDispute(disputeId string) (*betav0.Dispute, *betav0.RenDispute, error)
	GetPermission(permissionId string) (*betav0.Permission, *betav0.RenPermission, error)
	GetTag(tagId string) (*betav0.Tag, *betav0.RenTag, error)
	GetRoyalty(royaltyId string) (*betav0.Royalty, *betav0.RenRoyalty, error)
	GetRoyaltyLiquidSplit(royaltySplitId string) (*betav0.RoyaltySplit, *betav0.RenRoyaltySplit, error)
	GetRoyaltyPay(royaltyPayId string) (*betav0.RoyaltyPay, *betav0.RenRoyaltyPay, error)
	GetPolicyFrameworkManager(pfwmId string) (*betav0.PolicyFrameworkManager, *betav0.RenPolicyFrameworkManager, error)
	GetCollection(colId string) (*betav0.Collection, *betav0.RenCollection, error)
	GetRoyaltyPolicy(royaltyPolicyId string) (*betav0.RoyaltyPolicy, *betav0.RenRoyaltyPolicy, error)
	GetLicenseMintingFeePaid(licenseMintingFeePaidId string) (*betav0.LicenseMintingFeePaid, *betav0.RenLicenseMintingFeePaid, error)
	GetLicenseOwner(licenseOwnerId string) (*betav0.LicenseOwner, *betav0.RenLicenseOwner, error)
	GetTransaction(trxId string) (*betav0.Transaction, *betav0.RenTransaction, error)

	// LISTS
	ListIPAssets(options *TheGraphQueryOptions) ([]*betav0.IPAsset, []*betav0.RenIPAsset, error)
	ListModules(options *TheGraphQueryOptions) ([]*betav0.Module, []*betav0.RenModule, error)
	ListLicenses(options *TheGraphQueryOptions) ([]*betav0.License, []*betav0.RenLicense, error)
	ListLicenseFrameworks(options *TheGraphQueryOptions) ([]*betav0.LicenseFramework, []*betav0.RenLicenseFramework, error)
	ListAccessControlPermissions(options *TheGraphQueryOptions) ([]*betav0.AccessControlPermission, []*betav0.RenAccessControlPermission, error)
	ListPolicies(options *TheGraphQueryOptions) ([]*betav0.Policy, []*betav0.RenPolicy, error)
	ListIPAPolicies(options *TheGraphQueryOptions) ([]*betav0.IPAPolicy, []*betav0.RenIPAPolicy, error)
	ListDisputes(options *TheGraphQueryOptions) ([]*betav0.Dispute, []*betav0.RenDispute, error)
	ListPermissions(options *TheGraphQueryOptions) ([]*betav0.Permission, []*betav0.RenPermission, error)
	ListTag(options *TheGraphQueryOptions) ([]*betav0.Tag, []*betav0.RenTag, error)
	ListRoyalties(options *TheGraphQueryOptions) ([]*betav0.Royalty, []*betav0.RenRoyalty, error)
	ListRoyaltyPays(options *TheGraphQueryOptions) ([]*betav0.RoyaltyPay, []*betav0.RenRoyaltyPay, error)
	ListPolicyFrameworkManagers(options *TheGraphQueryOptions) ([]*betav0.PolicyFrameworkManager, []*betav0.RenPolicyFrameworkManager, error)
	ListCollections(options *TheGraphQueryOptions) ([]*betav0.Collection, []*betav0.RenCollection, error)
	ListRoyaltyPolicies(options *TheGraphQueryOptions) ([]*betav0.RoyaltyPolicy, []*betav0.RenRoyaltyPolicy, error)
	ListLicenseMintingFeePaids(options *TheGraphQueryOptions) ([]*betav0.LicenseMintingFeePaid, []*betav0.RenLicenseMintingFeePaid, error)
	ListLicenseOwners(options *TheGraphQueryOptions) ([]*betav0.LicenseOwner, []*betav0.RenLicenseOwner, error)
	ListTransactions(options *TheGraphQueryOptions) ([]*betav0.Transaction, []*betav0.RenTransaction, error)
}

type TheGraphQueryOptions struct {
	First          int
	Limit          int
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
		Transferable            string `json:"transferable,omitempty"`
		Owner                   string `json:"owner,omitempty"`
	}
}
