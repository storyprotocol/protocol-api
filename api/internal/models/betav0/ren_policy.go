package betav0

type RenPolicy struct {
	ID                     string `json:"id"`
	PolicyFrameworkManager string `json:"policy_framework_manager"`
	FrameworkData          string `json:"framework_data"`
	RoyaltyPolicy          string `json:"royalty_policy"`
	RoyaltyData            string `json:"royalty_data"`
	MintingFee             string `json:"minting_fee"`
	MintingFeeToken        string `json:"minting_fee_token"`

	BlockNumber    string       `json:"block_number"`
	BlockTimestamp string       `json:"block_time"`
	PIL            RenPILPolicy `json:"pil"`
}

type RenPILPolicy struct {
	ID                        string   `json:"id"`
	Attribution               bool     `json:"attribution"`
	CommercialUse             bool     `json:"commercial_use"`
	CommercialAttribution     bool     `json:"commercial_attribution"`
	CommercializerChecker     string   `json:"commercializer_checker"`
	CommercializerCheckerData string   `json:"commercializer_checker_data"`
	CommercialRevShare        string   `json:"commercial_rev_share"`
	DerivativesAllowed        bool     `json:"derivatives_allowed"`
	DerivativesAttribution    bool     `json:"derivatives_attribution"`
	DerivativesApproval       bool     `json:"derivatives_approval"`
	DerivativesReciprocal     bool     `json:"derivatives_reciprocal"`
	Territories               []string `json:"territories"`
	DistributionChannels      []string `json:"distribution_channels"`
	ContentRestrictions       []string `json:"content_restrictions"`
}

type RenPoliciesTheGraphResponse struct {
	Policies []*RenPolicy `json:"records"`
}

type RenPolicyTheGraphResponse struct {
	Policy *RenPolicy `json:"record"`
}

type RenPolicyResponse struct {
	Data *RenPolicy `json:"data"`
}

type RenPoliciesResponse struct {
	Data []*RenPolicy `json:"data"`
}
