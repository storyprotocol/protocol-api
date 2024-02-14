package betav0

type Policy struct {
	ID                     string    `json:"id"`
	PolicyFrameworkManager string    `json:"policyFrameworkManager"`
	BlockNumber            string    `json:"blockNumber"`
	BlockTimestamp         string    `json:"blockTimestamp"`
	UML                    UMLPolicy `json:"uml"`
}

type UMLPolicy struct {
	ID                      string   `json:"id"`
	FrameworkManagerAddress string   `json:"frameworkManagerAddress"`
	Attribution             bool     `json:"attribution"`
	Transferable            bool     `json:"transferable"`
	CommercialUse           bool     `json:"commercialUse"`
	CommercialAttribution   bool     `json:"commercialAttribution"`
	Commercializers         []string `json:"commercializers"`
	DerivativesAllowed      bool     `json:"derivativesAllowed"`
	DerivativesAttribution  bool     `json:"derivativesAttribution"`
	DerivativesApproval     bool     `json:"derivativesApproval"`
	DerivativesReciprocal   bool     `json:"derivativesReciprocal"`
	DerivativesRevShare     string   `json:"derivativesRevShare"`
	Territories             []string `json:"territories"`
	DistributionChannels    []string `json:"distributionChannels"`
	RoyaltyPolicy           string   `json:"royaltyPolicy"`
}

type PoliciesTheGraphResponse struct {
	Policies []*Policy `json:"policies"`
}

type PolicyTheGraphResponse struct {
	Policy *Policy `json:"policy"`
}

type PolicyResponse struct {
	Data *Policy `json:"data"`
}

type PoliciesResponse struct {
	Data []*Policy `json:"data"`
}

type PolicyRequestBody struct {
	Options *PolicyQueryOptions `json:"options"`
}

type PolicyQueryOptions struct {
	Pagination struct {
		Offset int `json:"offset"`
		Limit  int `json:"limit"`
	} `json:"pagination"`
	Where struct {
		PolicyFrameworkManager string `json:"policyFrameworkManager"`
	} `json:"where"`
	OrderBy        string `json:"orderBy"`
	OrderDirection string `json:"orderDirection"`
}
