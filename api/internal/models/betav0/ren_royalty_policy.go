package betav0

type RenRoyaltyPolicy struct {
	ID                  string `json:"id"`
	SplitClone          string `json:"split_clone"`
	AncestorsVault      string `json:"ancestors_vault"`
	RoyaltyStack        string `json:"royalty_stack"`
	TargetAncestors     string `json:"target_ancestors"`
	TargetRoyaltyAmount string `json:"target_royalty_amount"`

	BlockNumber    string `json:"block_number"`
	BlockTimestamp string `json:"block_time"`
}

type RenRoyaltyPolicyTheGraphResponse struct {
	RoyaltyPolicy *RenRoyaltyPolicy `json:"record"`
}

type RenRoyaltyPoliciesTheGraphResponse struct {
	Royalties []*RenRoyaltyPolicy `json:"records"`
}

type RenRoyaltyPolicyResponse struct {
	Data *RenRoyaltyPolicy `json:"data"`
}

type RenRoyaltyPoliciesResponse struct {
	Data []*RenRoyaltyPolicy `json:"data"`
}
