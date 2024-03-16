package betav0

type RenLicenseFramework struct {
	Creator                 string                     `json:"creator"`
	ID                      string                     `json:"id"`
	FrameworkCreationParams RenFrameworkCreationParams `json:"framework_creation_params"`
	BlockNumber             string                     `json:"block_number"`
	BlockTimestamp          string                     `json:"block_time"`
}

type RenFrameworkCreationParams struct {
	ID                           string   `json:"id"`
	ActivationParamDefaultValues []string `json:"activation_param_default_values"`
	ActivationParamVerifiers     []string `json:"activation_param_verifiers"`
	DefaultNeedsActivation       bool     `json:"default_needs_activation"`
	LinkParentParamDefaultValues []string `json:"link_parent_param_default_values"`
	LinkParentParamVerifiers     []string `json:"link_parent_param_verifiers"`
	MintingParamDefaultValues    []string `json:"minting_param_default_values"`
	MintingParamVerifiers        []string `json:"minting_param_verifiers"`
	LicenseUrl                   string   `json:"license_url"`
}

type RenLicenseFrameworkTheGraphResponse struct {
	LicenseFramework *RenLicenseFramework `json:"record"`
}

type RenLicenseFrameworksTheGraphResponse struct {
	LicenseFrameworks []*RenLicenseFramework `json:"records"`
}

type RenLicenseFrameworkResponse struct {
	Data *RenLicenseFramework `json:"data"`
}

type RenLicenseFrameworksResponse struct {
	Data []*RenLicenseFramework `json:"data"`
}
