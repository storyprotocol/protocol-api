package betav0

type RoyaltyPolicy struct {
	ID                  string   `json:"id"`
	SplitClone          string   `json:"splitClone"`
	AncestorsVault      string   `json:"ancestorsVault"`
	RoyaltyStack        string   `json:"royaltyStack"`
	TargetAncestors     []string `json:"targetAncestors"`
	TargetRoyaltyAmount []string `json:"targetRoyaltyAmount"`

	BlockNumber    string `json:"blockNumber"`
	BlockTimestamp string `json:"blockTimestamp"`
}

type RoyaltyPolicyTheGraphResponse struct {
	RoyaltyPolicy *RoyaltyPolicy `json:"royaltyPolicy"`
}

type RoyaltyPoliciesTheGraphResponse struct {
	Royalties []*RoyaltyPolicy `json:"royaltyPolicies"`
}

type RoyaltyPolicyResponse struct {
	Data *RoyaltyPolicy `json:"data"`
}

type RoyaltyPoliciesResponse struct {
	Data []*RoyaltyPolicy `json:"data"`
}

type RoyaltyPolicyRequestBody struct {
	Options *RoyaltyPolicyQueryOptions `json:"options"`
}
type RoyaltyPolicyQueryOptions struct {
	Pagination struct {
		Offset int `json:"offset"`
		Limit  int `json:"limit"`
	} `json:"pagination"`
	Where struct {
	} `json:"where"`
	OrderBy        string `json:"orderBy"`
	OrderDirection string `json:"orderDirection"`
}
