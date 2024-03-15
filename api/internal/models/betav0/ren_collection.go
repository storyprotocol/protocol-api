package betav0

type RenCollection struct {
	ID                     string `json:"id"`
	AssetCount             string `json:"asset_count"`
	LicensesCount          string `json:"licenses_count"`
	ResolvedDisputesCount  string `json:"resolved_dispute_count"`
	CancelledDisputesCount string `json:"cancelled_dispute_count"`
	RaisedDisputesCount    string `json:"raised_dispute_count"`
	JudgedDisputesCount    string `json:"judged_dispute_count"`
	BlockNumber            string `json:"block_number"`
	BlockTimestamp         string `json:"block_time"`
}

type RenCollectionTheGraphResponse struct {
	Collection *RenCollection `json:"record"`
}

type RenCollectionsTheGraphResponse struct {
	Collections []*RenCollection `json:"records"`
}

type RenCollectionResponse struct {
	Data *RenCollection `json:"data"`
}

type RenCollectionsResponse struct {
	Data []*RenCollection `json:"data"`
}
