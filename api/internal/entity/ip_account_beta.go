package entity

// Get IP ACCOUNT
type IPAccount struct {
	IPAccountAddress string `json:"account,omitempty"`
	IPAccountImpl    string `json:"implementation,omitempty"`
	ChainId          string `json:"chainId,omitempty"`
	TokenContract    string `json:"tokenContract,omitempty"`
	TokenId          string `json:"tokenId,omitempty"`
}

type IPAccountsTheGraphResponse struct {
	IPAccounts []*IPAccount `json:"ipaccountRegistereds"`
}

type IPAccountTheGraphResponse struct {
	IPAccount *IPAccount `json:"ipaccountRegistered"`
}

type IPAccountsResponse struct {
	Data []*IPAccount `json:"data"`
}
