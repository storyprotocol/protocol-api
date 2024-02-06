package beta_v0

type Policy struct {
	ID                     string `json:"id,omitempty"`
	PolicyID               string `json:"policyId,omitempty"`
	PolicyFrameworkManager string `json:"policyFrameworkManager,omitempty"`
	Policy                 string `json:"policy,omitempty"`
	BlockNumber            string `json:"blockNumber,omitempty"`
	BlockTimestamp         string `json:"blockTimestamp,omitempty"`
}

type UMLPolicy struct {
	ID                      string `json:"id,omitempty"`
	FrameworkManagerAddress string `json:"frameworkManagerAddress,omitempty"`
	Attribution             string `json:"attribution,omitempty"`
	Transferable            string `json:"transferable,omitempty"`
	CommercialUse           string `json:"commercialUse,omitempty"`
	CommercialAttribution   string `json:"commercialAttribution,omitempty"`
	Commercializers         string `json:"commercializers,omitempty"`
	DerivativesAllowed      string `json:"derivativesAllowed,omitempty"`
	DerivativesAttribution  string `json:"derivativesAttribution,omitempty"`
	DerivativesApproval     string `json:"derivativesApproval,omitempty"`
	DerivativesReciprocal   string `json:"derivativesReciprocal,omitempty"`
	DerivativesRevShare     string `json:"derivativesRevShare,omitempty"`
	Territories             string `json:"territories,omitempty"`
	DistributionChannels    string `json:"distributionChannels,omitempty"`
	RoyaltyPolicy           string `json:"royaltyPolicy,omitempty"`
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
