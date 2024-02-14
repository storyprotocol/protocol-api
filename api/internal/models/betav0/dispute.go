package betav0

type Dispute struct {
	ID                string `json:"id"`
	TargetIpId        string `json:"targetIpId"`
	TargetTag         string `json:"targetTag"`
	CurrentTag        string `json:"currentTag"`
	ArbitrationPolicy string `json:"arbitrationPolicy"`
	EvidenceLink      string `json:"evidenceLink"`
	Initiator         string `json:"initiator"`
	Data              string `json:"data"`
	BlockNumber       string `json:"blockNumber"`
	BlockTimestamp    string `json:"blockTimestamp"`
}

type DisputeTheGraphResponse struct {
	Dispute *Dispute `json:"dispute"`
}

type DisputesTheGraphResponse struct {
	Disputes []*Dispute `json:"disputes"`
}

type DisputeResponse struct {
	Data *Dispute `json:"data"`
}

type DisputesResponse struct {
	Data []*Dispute `json:"data"`
}

type DisputeRequestBody struct {
	Options *DisputeQueryOptions `json:"options"`
}
type DisputeQueryOptions struct {
	Pagination struct {
		Offset int `json:"offset"`
		Limit  int `json:"limit"`
	} `json:"pagination"`
	Where struct {
		TargetIpId string `json:"targetIpId"`
		TargetTag  string `json:"targetTag"`
		CurrentTag string `json:"currentTag"`
		Initiator  string `json:"initiator"`
	} `json:"where"`
	OrderBy        string `json:"orderBy"`
	OrderDirection string `json:"orderDirection"`
}
