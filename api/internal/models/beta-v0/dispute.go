package beta_v0

type Dispute struct {
	ID                   string `json:"id,omitempty"`
	IPID                 string `json:"ipId,omitempty"`
	TargetTag            string `json:"targetTag,omitempty"`
	ArbitrationPolicy    string `json:"arbitrationPolicy,omitempty"`
	DisputeEveidenceLink string `json:"disputeEveidenceLink,omitempty"`
	Initiator            string `json:"initiator,omitempty"`
}

type DisputeTheGraphResponse struct {
	Dispute []*Module `json:"dispute"`
}

type DisputesTheGraphResponse struct {
	Disputes []*Module `json:"disputes"`
}

type DisputeResponse struct {
	Data []*Module `json:"data"`
}
