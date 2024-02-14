package betav0

// Get IP ACCOUNT
type IPAsset struct {
	ID             string   `json:"id"`
	ChainId        string   `json:"chainId"`
	ParentIpIds    []string `json:"parentIpIds"`
	ChildIpIds     []string `json:"childIpIds"`
	RootIpIds      []string `json:"rootIpIds"`
	TokenContract  string   `json:"tokenContract"`
	TokenId        string   `json:"tokenId"`
	Resolver       string   `json:"metadataResolverAddress"`
	Metadata       Metadata `json:"metadata"`
	BlockNumber    string   `json:"blockNumber"`
	BlockTimestamp string   `json:"blockTimestamp"`
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
type IpAssetRequestBody struct {
	Options *IpAssetQueryOptions `json:"options"`
}

type IpAssetQueryOptions struct {
	Pagination struct {
		Offset int `json:"offset"`
		Limit  int `json:"limit"`
	} `json:"pagination"`
	Where struct {
		MetadataResolverAddress string `json:"metadataResolverAddress"`
		TokenContract           string `json:"tokenContract"`
		TokenId                 string `json:"tokenId"`
		ChainId                 string `json:"chainId"`
	} `json:"where"`
	OrderBy        string `json:"orderBy"`
	OrderDirection string `json:"orderDirection"`
}
