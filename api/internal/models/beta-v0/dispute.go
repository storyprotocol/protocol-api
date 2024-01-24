package beta_v0

type Dispute struct {
	ID                   string `json:"id,omitempty"`
	IPID                 string `json:"ipId,omitempty"`
	TargetTag            string `json:"targetTag,omitempty"`
	ArbitrationPolicy    string `json:"arbitrationPolicy,omitempty"`
	DisputeEveidenceLink string `json:"disputeEveidenceLink,omitempty"`
	Initiator            string `json:"initiator,omitempty"`
	BlockNumber          string `json:"blockNumber,omitempty"`
	BlockTimestamp       string `json:"blockTimestamp,omitempty"`
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
