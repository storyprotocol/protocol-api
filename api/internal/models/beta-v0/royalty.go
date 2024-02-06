package beta_v0

type Royalty struct {
	ID             string `json:"id,omitempty"`
	IPID           string `json:"ipId,omitempty"`
	Data           string `json:"data,omitempty"`
	BlockNumber    string `json:"blockNumber,omitempty"`
	BlockTimestamp string `json:"blockTimestamp,omitempty"`
}

type RoyaltyTheGraphResponse struct {
	Royalty *Royalty `json:"royalty"`
}

type RoyaltiesTheGraphResponse struct {
	Royalties []*Royalty `json:"royalties"`
}

type RoyaltyResponse struct {
	Data *Royalty `json:"data"`
}

type RoyaltiesResponse struct {
	Data []*Royalty `json:"data"`
}
