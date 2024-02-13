package betav0

type Collection struct {
	ID             string `json:"id,omitempty"`
	AssetCount     string `json:"assetCount,omitempty"`
	BlockNumber    string `json:"blockNumber,omitempty"`
	BlockTimestamp string `json:"blockTimestamp,omitempty"`
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
