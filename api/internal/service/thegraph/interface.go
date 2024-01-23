package thegraph

import "github.com/storyprotocol/protocol-api/api/internal/entity"

type TheGraphServiceMvp interface {
	GetFranchises() ([]*entity.FranchiseMVP, error)
	GetFranchise(franchiseId string) (*entity.FranchiseMVP, error)
	GetIpAssets(franchiseId string) ([]*entity.IpAssetMVP, error)
	GetIpAsset(franchiseId string, ipAssetId string) (*entity.IpAssetMVP, error)
	GetLicenses(franchiseId string, ipAssetId string) ([]*entity.LicenseMVP, error)
	GetLicense(licenseId string) (*entity.LicenseMVP, error)
	GetCollections(franchiseId string) ([]*entity.CollectionMVP, error)
	GetTransactions() ([]*entity.TransactionMVP, error)
	GetTransaction(transactionId string) (*entity.TransactionMVP, error)
}

type TheGraphServiceBeta interface {
	GetIPAccount(accountId string) ([]*entity.IPAccount, error)
	GetIPAccounts(limit int64, offset int64) ([]*entity.IPAccount, error)
	GetModule(moduleName string) ([]*entity.Module, error)
	GetModules(limit int64, offset int64) ([]*entity.Module, error)

	//GetIPsRegistered() ([]*entity.IPRegistered, error)
	//GetSetIPAccounts() ([]*entity.SetIPAccount, error)
	//GetSetIPResolvers() ([]*entity.SetResolver, error)
}
