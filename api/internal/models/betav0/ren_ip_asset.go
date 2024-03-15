package betav0

// Get IP ACCOUNT
type RenIPAsset struct {
	ID             string      `json:"id"`
	ChainId        string      `json:"chain_id"`
	ParentIpIds    []IPAsset   `json:"parent_ip_ids"`
	ChildIpIds     []IPAsset   `json:"child_ip_ids"`
	RootIpIds      []IPAsset   `json:"root_ip_ids"`
	TokenContract  string      `json:"token_contract"`
	TokenId        string      `json:"token_id"`
	Resolver       string      `json:"metadata_resolver_address"`
	Metadata       RenMetadata `json:"metadata"`
	BlockNumber    string      `json:"block_number"`
	BlockTimestamp string      `json:"block_time"`
}

type RenMetadata struct {
	Name             string `json:"name"`
	Hash             string `json:"hash"`
	RegistrationDate string `json:"registration_date"`
	Registrant       string `json:"registrant"`
	URI              string `json:"uri"`
}

type RenIPAssetsTheGraphResponse struct {
	IPAssets []*RenIPAsset `json:"records"`
}

type RenIPAssetTheGraphResponse struct {
	IPAsset *RenIPAsset `json:"record"`
}

type RenIPAssetResponse struct {
	Data *RenIPAsset `json:"data"`
}

type RenIPAssetsResponse struct {
	Data []*RenIPAsset `json:"data"`
}
