package betav0

type PolicyFrameworkManager struct {
	ID             string `json:"id,omitempty"`
	Address        string `json:"address,omitempty"`
	Name           string `json:"name,omitempty"`
	LicenseUrl     string `json:"licenseUrl,omitempty"`
	BlockNumber    string `json:"blockNumber,omitempty"`
	BlockTimestamp string `json:"blockTimestamp,omitempty"`
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
