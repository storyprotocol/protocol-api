package betav0

type Collection struct {
	ID                     string `json:"id"`
	AssetCount             string `json:"assetCount"`
	LicensesCount          string `json:"licensesCount"`
	ResolvedDisputesCount  string `json:"resolvedDisputeCount"`
	CancelledDisputesCount string `json:"cancelledDisputeCount"`
	RaisedDisputesCount    string `json:"raisedDisputeCount"`
	JudgedDisputesCount    string `json:"judgedDisputeCount"`
	BlockNumber            string `json:"blockNumber"`
	BlockTimestamp         string `json:"blockTimestamp"`
}

type CollectionTheGraphResponse struct {
	Collection *Collection `json:"collection"`
}

type CollectionsTheGraphResponse struct {
	Collections []*Collection `json:"collections"`
}

type CollectionResponse struct {
	Data *Collection `json:"data"`
}

type CollectionsResponse struct {
	Data []*Collection `json:"data"`
}

type CollectionsRequestBody struct {
	Options *CollectionQueryOptions `json:"options"`
}

type CollectionQueryOptions struct {
	Pagination struct {
		Offset int `json:"offset"`
		Limit  int `json:"limit"`
	} `json:"pagination"`
	Where struct {
	} `json:"where"`
	OrderBy        string `json:"orderBy"`
	OrderDirection string `json:"orderDirection"`
}
