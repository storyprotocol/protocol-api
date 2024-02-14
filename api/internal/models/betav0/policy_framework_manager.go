package betav0

type PolicyFrameworkManager struct {
	ID             string `json:"id"`
	Address        string `json:"address"`
	Name           string `json:"name"`
	LicenseUrl     string `json:"licenseUrl"`
	BlockNumber    string `json:"blockNumber"`
	BlockTimestamp string `json:"blockTimestamp"`
}

type PolicyFrameworkManagersTheGraphResponse struct {
	PolicyFrameworkManagers []*PolicyFrameworkManager `json:"policyFrameworkManagers"`
}

type PolicyFrameworkManagerTheGraphResponse struct {
	PolicyFrameworkManager *PolicyFrameworkManager `json:"policyFrameworkManager"`
}

type PolicyFrameworkManagerResponse struct {
	Data *PolicyFrameworkManager `json:"data"`
}

type PolicyFrameworkManagersResponse struct {
	Data []*PolicyFrameworkManager `json:"data"`
}
type PolicyFrameworkManagerRequestBody struct {
	Options *PFWMQueryOptions `json:"options"`
}

type PFWMQueryOptions struct {
	Pagination struct {
		Offset int `json:"offset"`
		Limit  int `json:"limit"`
	} `json:"pagination"`
	Where struct {
		Address string `json:"address"`
		Name    string `json:"name"`
	} `json:"where"`
	OrderBy        string `json:"orderBy"`
	OrderDirection string `json:"orderDirection"`
}
