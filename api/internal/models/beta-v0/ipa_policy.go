package beta_v0

type IPAPolicy struct {
	ID             string `json:"ipId,omitempty"`
	PolicyId       string `json:"policyId,omitempty"`
	Index          string `json:"index,omitempty"`
	Active         bool   `json:"active,omitempty"`
	Inherited      bool   `json:"inherited,omitempty"`
	BlockNumber    string `json:"blockNumber,omitempty"`
	BlockTimestamp string `json:"blockTimestamp,omitempty"`
}

type IPAPoliciesTheGraphResponse struct {
	IPAPolicies []*IPAPolicy `json:"ipaPolicies"`
}

type IPAPolicyTheGraphResponse struct {
	IPAPolicy *IPAPolicy `json:"ipaPolicy"`
}

type IPAPolicyResponse struct {
	Data *IPAPolicy `json:"data"`
}

type IPAPoliciesResponse struct {
	Data []*IPAPolicy `json:"data"`
}
