package betav0

type Collection struct {
	ID                     string `json:"id,omitempty"`
	AssetCount             string `json:"assetCount,omitempty"`
	LicensesCount          string `json:"licensesCount,omitempty"`
	ResolvedDisputesCount  string `json:"resolvedDisputeCount,omitempty"`
	CancelledDisputesCount string `json:"cancelledDisputeCount,omitempty"`
	RaisedDisputesCount    string `json:"raisedDisputeCount,omitempty"`
	JudgedDisputesCount    string `json:"judgedDisputesCount,omitempty"`
	BlockNumber            string `json:"blockNumber,omitempty"`
	BlockTimestamp         string `json:"blockTimestamp,omitempty"`
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
