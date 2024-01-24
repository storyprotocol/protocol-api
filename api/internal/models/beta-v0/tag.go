package beta_v0

type Tag struct {
	ID        string `json:"id,omitempty"`
	IPID      string `json:"ipId,omitempty"`
	Tag       string `json:"tag,omitempty"`
	DeletedAt string `json:"deletedAt,omitempty"`
}

type TagTheGraphResponse struct {
	Tag []*Tag `json:"tag"`
}

type TagsTheGraphResponse struct {
	Tags []*Tag `json:"tags"`
}

type TagResponse struct {
	Data []*Tag `json:"data"`
}
