package betav0

type Royalty struct {
	ID             string `json:"id,omitempty"`
	IPID           string `json:"ipId,omitempty"`
	Data           string `json:"data,omitempty"`
	RoyaltyPolicy  string `json:"royaltyPolicy,omitempty"`
	BlockNumber    string `json:"blockNumber,omitempty"`
	BlockTimestamp string `json:"blockTimestamp,omitempty"`
}

type RoyaltyTheGraphResponse struct {
	Royalty *Royalty `json:"iproyalty"`
}

type RoyaltiesTheGraphResponse struct {
	Royalties []*Royalty `json:"iproyalties"`
}

type RoyaltyResponse struct {
	Data *Royalty `json:"data"`
}

type RoyaltiesResponse struct {
	Data []*Royalty `json:"data"`
}
