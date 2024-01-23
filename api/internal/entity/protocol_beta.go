package entity

type RequestBody struct {
	Options *QueryOptions `json:"options"`
}

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

// GET IP Record Registered
type IPRegistered struct {
	ID              string `json:"id,omitempty"`
	ChainId         string `json:"chainId,omitempty"`
	TokenContract   string `json:"tokenContract,omitempty"`
	TokenId         string `json:"tokenId,omitempty"`
	ResolverAddress string `json:"resolver,omitempty"`
}

type IPRegisteredTheGraphResponse struct {
	IPRegistered []*IPRegistered `json:"ipRegistereds"`
}

type IPRegisteredResponse struct {
	Data []*IPRegistered `json:"data"`
}

// GET Create IP Account
type SetIPAccount struct {
	IPID          string `json:"ipId,omitempty"`
	ChainId       string `json:"chainId,omitempty"`
	TokenContract string `json:"tokenContract,omitempty"`
	TokenId       string `json:"tokenId,omitempty"`
}

type SetIPAccountTheGraphResponse struct {
	SetIPAccount []*SetIPAccount `json:"setIpAccounts"`
}

type SetIPAccountResponse struct {
	Data []*SetIPAccount `json:"data"`
}

// GET Set Resolver
type SetResolver struct {
	IPID            string `json:"ipId,omitempty"`
	ResolverAddress string `json:"resolver,omitempty"`
}

type SetResolverTheGraphResponse struct {
	SetResolver []*SetResolver `json:"ipresolverSets"`
}

type SetResolverResponse struct {
	Data []*SetResolver `json:"data"`
}

// GET Registered Modules
type Module struct {
	Name   string `json:"name,omitempty"`
	Module string `json:"module,omitempty"`
}

type ModuleTheGraphResponse struct {
	Module []*Module `json:"modules"`
}

type ModuleResponse struct {
	Data []*Module `json:"data"`
}
