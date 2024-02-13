package betav0

type Dispute struct {
	ID                string `json:"id,omitempty"`
	TargetIpId        string `json:"targetIpId,omitempty"`
	TargetTag         string `json:"targetTag,omitempty"`
	CurrentTag        string `json:"currentTag,omitempty"`
	ArbitrationPolicy string `json:"arbitrationPolicy,omitempty"`
	EvidenceLink      string `json:"evidenceLink,omitempty"`
	Initiator         string `json:"initiator,omitempty"`
	Data              string `json:"data,omitempty"`
	BlockNumber       string `json:"blockNumber,omitempty"`
	BlockTimestamp    string `json:"blockTimestamp,omitempty"`
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
