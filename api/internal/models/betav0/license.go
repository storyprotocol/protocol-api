package betav0

type License struct {
	ID           string `json:"id"`
	PolicyID     string `json:"policyId"`
	LicensorIpId string `json:"licensorIpId"`
	Transferable bool   `json:"transferable"`
	Amount       string `json:"amount"`

	BlockNumber    string `json:"blockNumber"`
	BlockTimestamp string `json:"blockTimestamp"`
}

type LicenseTheGraphResponse struct {
	License *License `json:"license"`
}

type LicensesTheGraphResponse struct {
	Licenses []*License `json:"licenses"`
}

type LicenseResponse struct {
	Data *License `json:"data"`
}

type LicensesResponse struct {
	Data []*License `json:"data"`
}

type LicenseRequestBody struct {
	Options *LicenseQueryOptions `json:"options"`
}
type LicenseQueryOptions struct {
	Pagination struct {
		Offset int `json:"offset"`
		Limit  int `json:"limit"`
	} `json:"pagination"`
	Where struct {
		PolicyId      string `json:"policyId"`
		LicensorIpdId string `json:"licensorIpdId"`
		Transferable  bool   `json:"transferable"`
	} `json:"where"`
	OrderBy        string `json:"orderBy"`
	OrderDirection string `json:"orderDirection"`
}
