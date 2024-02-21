package betav0

type Policy struct {
	ID                     string `json:"id"`
	PolicyFrameworkManager string `json:"policyFrameworkManager"`
	FrameworkData          string `json:"frameworkData"`
	RoyaltyPolicy          string `json:"royaltyPolicy"`
	RoyaltyData            string `json:"royaltyData"`
	MintingFee             string `json:"mintingFee"`
	MintingFeeToken        string `json:"mintingFeeToken"`

	BlockNumber    string    `json:"blockNumber"`
	BlockTimestamp string    `json:"blockTimestamp"`
	PIL            PILPolicy `json:"pil"`
}

type PILPolicy struct {
	ID                        string   `json:"id"`
	Attribution               bool     `json:"attribution"`
	CommercialUse             bool     `json:"commercialUse"`
	CommercialAttribution     bool     `json:"commercialAttribution"`
	CommercializerChecker     string   `json:"commercializerChecker"`
	CommercializerCheckerData string   `json:"commercializerCheckerData"`
	CommercialRevShare        string   `json:"commercialRevShare"`
	DerivativesAllowed        bool     `json:"derivativesAllowed"`
	DerivativesAttribution    bool     `json:"derivativesAttribution"`
	DerivativesApproval       bool     `json:"derivativesApproval"`
	DerivativesReciprocal     bool     `json:"derivativesReciprocal"`
	Territories               []string `json:"territories"`
	DistributionChannels      []string `json:"distributionChannels"`
	ContentRestrictions       []string `json:"contentRestrictions"`
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
		MintingFeeToken        string `json:"mintingFeeToken"`
		RoyaltyPolicy          string `json:"royaltyPolicy"`
	} `json:"where"`
	OrderBy        string `json:"orderBy"`
	OrderDirection string `json:"orderDirection"`
}
