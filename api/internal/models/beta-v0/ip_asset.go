package beta_v0

// Get IP ACCOUNT
type IPAsset struct {
	ID             string   `json:"id,omitempty"`
	ChainId        string   `json:"chainId,omitempty"`
	ParentIpIds    []string `json:"parentIpIds,omitempty"`
	TokenContract  string   `json:"tokenContract,omitempty"`
	TokenId        string   `json:"tokenId,omitempty"`
	Resolver       string   `json:"metadataResolverAddress,omitempty"`
	Metadata       Metadata `json:"metadata,omitempty"`
	BlockNumber    string   `json:"blockNumber,omitempty"`
	BlockTimestamp string   `json:"blockTimestamp,omitempty"`
}

type Metadata struct {
	Name             string `json:"name"`
	Hash             string `json:"hash"`
	RegistrationDate string `json:"registrationDate"`
	Registrant       string `json:"registrant"`
	URI              string `json:"uri"`
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
