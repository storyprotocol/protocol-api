package betav0

type RenPolicyFrameworkManager struct {
	ID             string `json:"id"`
	Address        string `json:"address"`
	Name           string `json:"name"`
	LicenseUrl     string `json:"license_url"`
	LicenseTextUrl string `json:"license_text_url,omitempty"`
	BlockNumber    string `json:"block_number"`
	BlockTimestamp string `json:"block_time"`
}

type RenPolicyFrameworkManagersTheGraphResponse struct {
	PolicyFrameworkManagers []*RenPolicyFrameworkManager `json:"records"`
}

type RenPolicyFrameworkManagerTheGraphResponse struct {
	PolicyFrameworkManager *RenPolicyFrameworkManager `json:"record"`
}

type RenPolicyFrameworkManagerResponse struct {
	Data *RenPolicyFrameworkManager `json:"data"`
}

type RenPolicyFrameworkManagersResponse struct {
	Data []*RenPolicyFrameworkManager `json:"data"`
}
