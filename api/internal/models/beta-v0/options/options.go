package options

type QueryOptions struct {
	Pagination struct {
		Offset int `json:"offset"`
		Limit  int `json:"limit"`
	} `json:"pagination"`
	Where struct {
		Creator       string `json:"creator"`
		Receiver      string `json:"receiver"`
		TokenContract string `json:"tokenContract"`
		FrameworkId   string `json:"frameworkId"`
		IPAccount     string `json:"ipAccount"`
		IPID          string `json:"ipId"`
	} `json:"where"`
	OrderBy        string `json:"orderBy"`
	OrderDirection string `json:"orderDirection"`
}

type RequestBody struct {
	Options *QueryOptions `json:"options"`
}
