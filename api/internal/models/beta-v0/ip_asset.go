package beta_v0

// Get IP ACCOUNT
type IPAsset struct {
	ID             string `json:"id,omitempty"`
	IPID           string `json:"ipId,omitempty"`
	ChainId        string `json:"chainId,omitempty"`
	TokenContract  string `json:"tokenContract,omitempty"`
	TokenId        string `json:"tokenId,omitempty"`
	Resolver       string `json:"metadataResolverAddress,omitempty"`
	Metadata       string `json:"metadata,omitempty"`
	BlockNumber    string `json:"blockNumber,omitempty"`
	BlockTimestamp string `json:"blockTimestamp,omitempty"`
}

type Metadata struct {
	Name             string `json:"name,omitempty"`
	Hash             string `json:"hash,omitempty"`
	RegistrationDate string `json:"registrationDate,omitempty"`
	Registrant       string `json:"registrant,omitempty"`
	URI              string `json:"uri,omitempty"`
}

type IPAssetsTheGraphResponse struct {
	IPAssets []*IPAsset `json:"ipassets"`
}

type IPAssetTheGraphResponse struct {
	IPAsset *IPAsset `json:"ipasset"`
}

type IPAssetResponse struct {
	Data *IPAsset `json:"data"`
}

type IPAssetsResponse struct {
	Data []*IPAsset `json:"data"`
}

// Get IP ACCOUNT
//type IPAsset struct {
//	IPAssetAddress string `json:"asset,omitempty"`
//	IPAssetImpl    string `json:"implementation,omitempty"`
//	ChainId          string `json:"chainId,omitempty"`
//	TokenContract    string `json:"tokenContract,omitempty"`
//	TokenId          string `json:"tokenId,omitempty"`
//}
//
//type IPAssetsTheGraphResponse struct {
//	IPAssets []*IPAsset `json:"ipassetRegistereds"`
//}
//
//type IPAssetTheGraphResponse struct {
//	IPAsset *IPAsset `json:"ipassetRegistered"`
//}
//
//type IPAssetsResponse struct {
//	Data []*IPAsset `json:"data"`
//}
