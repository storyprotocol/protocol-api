package entity

// Get IP ACCOUNT
type IPAccount struct {
	ID            string `json:"id,omitempty"`
	IPID          string `json:"ipId,omitempty"`
	ChainId       string `json:"chainId,omitempty"`
	TokenContract string `json:"tokenContract,omitempty"`
	TokenId       string `json:"tokenId,omitempty"`
	Resolver      string `json:"metadataResolverAddress,omitempty"`
}

type IPAccountsTheGraphResponse struct {
	IPAccounts []*IPAccount `json:"iprecords"`
}

type IPAccountTheGraphResponse struct {
	IPAccount *IPAccount `json:"iprecord"`
}

type IPAccountsResponse struct {
	Data []*IPAccount `json:"data"`
}

// Get IP ACCOUNT
//type IPAccount struct {
//	IPAccountAddress string `json:"account,omitempty"`
//	IPAccountImpl    string `json:"implementation,omitempty"`
//	ChainId          string `json:"chainId,omitempty"`
//	TokenContract    string `json:"tokenContract,omitempty"`
//	TokenId          string `json:"tokenId,omitempty"`
//}
//
//type IPAccountsTheGraphResponse struct {
//	IPAccounts []*IPAccount `json:"ipaccountRegistereds"`
//}
//
//type IPAccountTheGraphResponse struct {
//	IPAccount *IPAccount `json:"ipaccountRegistered"`
//}
//
//type IPAccountsResponse struct {
//	Data []*IPAccount `json:"data"`
//}
