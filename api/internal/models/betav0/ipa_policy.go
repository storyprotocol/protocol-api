package betav0

type IPAPolicy struct {
	ID             string `json:"id"`
	IPID           string `json:"ipId"`
	PolicyId       string `json:"policyId"`
	Index          string `json:"index"`
	Active         bool   `json:"active"`
	Inherited      bool   `json:"inherited"`
	BlockNumber    string `json:"blockNumber"`
	BlockTimestamp string `json:"blockTimestamp"`
}

type IPAPoliciesTheGraphResponse struct {
	IPAPolicies []*IPAPolicy `json:"ipapolicies"`
}

type IPAPolicyTheGraphResponse struct {
	IPAPolicy *IPAPolicy `json:"ipapolicy"`
}

type IPAPolicyResponse struct {
	Data *IPAPolicy `json:"data"`
}

type IPAPoliciesResponse struct {
	Data []*IPAPolicy `json:"data"`
}

type IPAPolicyRequestBody struct {
	Options *IPAPQueryOptions `json:"options"`
}
type IPAPQueryOptions struct {
	Pagination struct {
		Offset int `json:"offset"`
		Limit  int `json:"limit"`
	} `json:"pagination"`
	Where struct {
		PolicyId  string `json:"policyId"`
		Active    string `json:"active"`
		Inherited string `json:"inherited"`
	} `json:"where"`
	OrderBy        string `json:"orderBy"`
	OrderDirection string `json:"orderDirection"`
}
