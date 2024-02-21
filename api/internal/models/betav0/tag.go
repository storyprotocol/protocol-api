package betav0

type Tag struct {
	ID             string `json:"id"`
	UUID           string `json:"uuid,omitempty"`
	IPID           string `json:"ipId"`
	Tag            string `json:"tag"`
	DeletedAt      string `json:"deletedAt"`
	BlockNumber    string `json:"blockNumber"`
	BlockTimestamp string `json:"blockTimestamp"`
}

type TagTheGraphResponse struct {
	Tag *Tag `json:"tag"`
}

type TagsTheGraphResponse struct {
	Tags []*Tag `json:"tags"`
}

type TagResponse struct {
	Data *Tag `json:"data"`
}

type TagsResponse struct {
	Data []*Tag `json:"data"`
}

type TagRequestBody struct {
	Options *TagQueryOptions `json:"options"`
}
type TagQueryOptions struct {
	Pagination struct {
		Offset int `json:"offset"`
		Limit  int `json:"limit"`
	} `json:"pagination"`
	Where struct {
		IPID string `json:"ipId"`
		Tag  string `json:"tag"`
	} `json:"where"`
	OrderBy        string `json:"orderBy"`
	OrderDirection string `json:"orderDirection"`
}
