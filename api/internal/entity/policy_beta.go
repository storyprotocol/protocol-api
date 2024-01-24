package entity

type Policy struct {
	ID         string     `json:"policyId,omitempty"`
	Creator    string     `json:"creator,omitempty"`
	PolicyData PolicyData `json:"policyData,omitempty"`
}

type PolicyData struct {
	FrameworkId           string   `json:"frameworkId,omitempty"`
	NeedsActivation       bool     `json:"needsActivation,omitempty"`
	MintingParamValues    []string `json:"mintingParamValues,omitempty"`
	LinkParentParamValues []string `json:"linkParentParamValues,omitempty"`
	ActivationParamValues []string `json:"activationParamValues,omitempty"`
}

type PoliciesTheGraphResponse struct {
	Policies []*Policy `json:"policies"`
}

type PolicyTheGraphResponse struct {
	Policy []*Policy `json:"policy"`
}

type PolicyResponse struct {
	Data []*Policy `json:"data"`
}
