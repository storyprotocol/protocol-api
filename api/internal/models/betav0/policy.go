package betav0

type Policy struct {
	ID                     string    `json:"id,omitempty"`
	PolicyID               string    `json:"policyId,omitempty"`
	PolicyFrameworkManager string    `json:"policyFrameworkManager,omitempty"`
	Policy                 string    `json:"policy,omitempty"`
	BlockNumber            string    `json:"blockNumber,omitempty"`
	BlockTimestamp         string    `json:"blockTimestamp,omitempty"`
	UML                    UMLPolicy `json:"uml,omitempty"`
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
