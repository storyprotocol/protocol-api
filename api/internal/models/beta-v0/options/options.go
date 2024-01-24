package options

type QueryOptions struct {
	Pagination struct {
		Offset int `json:"offset"`
		Limit  int `json:"limit"`
	} `json:"pagination"`
}

type RequestBody struct {
	Options *QueryOptions `json:"messages"`
}
