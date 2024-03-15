package betav0

type RenDispute struct {
	ID                string `json:"id"`
	TargetIpId        string `json:"target_ip_id"`
	TargetTag         string `json:"target_tag"`
	CurrentTag        string `json:"current_tag"`
	ArbitrationPolicy string `json:"arbitration_policy"`
	EvidenceLink      string `json:"evidence_link"`
	Initiator         string `json:"initiator"`
	Data              string `json:"data"`
	BlockNumber       string `json:"block_number"`
	BlockTimestamp    string `json:"block_time"`
}

type RenDisputeTheGraphResponse struct {
	Dispute *RenDispute `json:"record"`
}

type RenDisputesTheGraphResponse struct {
	Disputes []*RenDispute `json:"records"`
}

type RenDisputeResponse struct {
	Data *RenDispute `json:"data"`
}

type RenDisputesResponse struct {
	Data []*RenDispute `json:"data"`
}
