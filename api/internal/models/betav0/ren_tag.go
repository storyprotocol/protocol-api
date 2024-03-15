package betav0

type RenTag struct {
	ID             string `json:"id"`
	UUID           string `json:"uuid,omitempty"`
	IPID           string `json:"ip_id"`
	Tag            string `json:"tag"`
	DeletedAt      string `json:"deleted_at"`
	BlockNumber    string `json:"block_number"`
	BlockTimestamp string `json:"block_time"`
}

type RenTagTheGraphResponse struct {
	Tag *RenTag `json:"record"`
}

type RenTagsTheGraphResponse struct {
	Tags []*RenTag `json:"records"`
}

type RenTagResponse struct {
	Data *RenTag `json:"data"`
}

type RenTagsResponse struct {
	Data []*RenTag `json:"data"`
}
