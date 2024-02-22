package betav0

type LicenseOwner struct {
	ID       string `json:"id"`
	PolicyId string `json:"policyId"`
	Owner    string `json:"owner"`
	Amount   string `json:"amount"`

	BlockNumber    string `json:"blockNumber"`
	BlockTimestamp string `json:"blockTimestamp"`
}

type LicenseOwnerTheGraphResponse struct {
	LicenseOwner *LicenseOwner `json:"licenseOwner"`
}

type LicenseOwnersTheGraphResponse struct {
	LicenseOwners []*LicenseOwner `json:"licenseOwners"`
}

type LicenseOwnerResponse struct {
	Data *LicenseOwner `json:"data"`
}

type LicenseOwnersResponse struct {
	Data []*LicenseOwner `json:"data"`
}

type LicenseOwnersRequestBody struct {
	Options *LicenseOwnerQueryOptions `json:"options"`
}
type LicenseOwnerQueryOptions struct {
	Pagination struct {
		Offset int `json:"offset"`
		Limit  int `json:"limit"`
	} `json:"pagination"`
	Where struct {
		PolicyId string `json:"policyId"`
		Owner    string `json:"owner"`
	} `json:"where"`
	OrderBy        string `json:"orderBy"`
	OrderDirection string `json:"orderDirection"`
}
