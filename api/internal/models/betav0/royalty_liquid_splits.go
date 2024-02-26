package betav0

type RoyaltySplit struct {
	ID                 string   `json:"id"`
	Holders            []Holder `json:"holders"`
	ClaimFromIPPoolArg string   `json:"claimFromIPPoolArg"`
}

type Holder struct {
	ID        string `json:"id"`
	Ownership string `json:"ownership"`
}

type RoyaltySplitTheGraphResponse struct {
	RoyaltySplit *RoyaltySplit `json:"liquidSplit"`
}
type RoyaltySplitResponse struct {
	Data *RoyaltySplit `json:"data"`
}
