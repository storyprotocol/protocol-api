package betav0

type RenRoyalty struct {
	ID             string `json:"id"`
	IPID           string `json:"ip_id"`
	Data           string `json:"data"`
	RoyaltyPolicy  string `json:"royalty_policy"`
	BlockNumber    string `json:"block_number"`
	BlockTimestamp string `json:"block_time"`
}

type RenRoyaltyTheGraphResponse struct {
	Royalty *RenRoyalty `json:"record"`
}

type RenRoyaltiesTheGraphResponse struct {
	Royalties []*RenRoyalty `json:"records"`
}

type RenRoyaltyResponse struct {
	Data *RenRoyalty `json:"data"`
}

type RenRoyaltiesResponse struct {
	Data []*RenRoyalty `json:"data"`
}
