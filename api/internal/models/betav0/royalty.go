package betav0

type Royalty struct {
	ID             string `json:"id"`
	IPID           string `json:"ipId"`
	Data           string `json:"data"`
	RoyaltyPolicy  string `json:"royaltyPolicy"`
	BlockNumber    string `json:"blockNumber"`
	BlockTimestamp string `json:"blockTimestamp"`
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

type RoyaltyRequestBody struct {
	Options *RoyaltyQueryOptions `json:"options"`
}
type RoyaltyQueryOptions struct {
	Pagination struct {
		Offset int `json:"offset"`
		Limit  int `json:"limit"`
	} `json:"pagination"`
	Where struct {
		IPID          string `json:"ipId"`
		RoyaltyPolicy string `json:"royaltyPolicy"`
	} `json:"where"`
	OrderBy        string `json:"orderBy"`
	OrderDirection string `json:"orderDirection"`
}
