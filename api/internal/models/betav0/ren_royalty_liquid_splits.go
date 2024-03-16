package betav0

type RenRoyaltySplit struct {
	ID                 string      `json:"id"`
	Holders            []RenHolder `json:"holders"`
	ClaimFromIPPoolArg string      `json:"claim_from_ip_pool_arg"`
}

type RenHolder struct {
	ID        string `json:"id"`
	Ownership string `json:"ownership"`
}

type RenRoyaltySplitTheGraphResponse struct {
	RoyaltySplit *RenRoyaltySplit `json:"record"`
}
type RenRoyaltySplitResponse struct {
	Data *RenRoyaltySplit `json:"data"`
}
