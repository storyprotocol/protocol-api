package betav0

type RenIPAPolicy struct {
	ID             string `json:"id"`
	IPID           string `json:"ip_id"`
	PolicyId       string `json:"policy_id"`
	Index          string `json:"index"`
	Active         string `json:"active"`
	Inherited      string `json:"inherited"`
	BlockNumber    string `json:"block_number"`
	BlockTimestamp string `json:"block_time"`
}

type RenIPAPoliciesTheGraphResponse struct {
	IPAPolicies []*RenIPAPolicy `json:"records"`
}

type RenIPAPolicyTheGraphResponse struct {
	IPAPolicy *RenIPAPolicy `json:"record"`
}

type RenIPAPolicyResponse struct {
	Data *RenIPAPolicy `json:"data"`
}

type RenIPAPoliciesResponse struct {
	Data []*RenIPAPolicy `json:"data"`
}
