package beta_v0

type LicenseFramework struct {
	Creator                 string                  `json:"creator,omitempty"`
	ID                      string                  `json:"id,omitempty"`
	FrameworkCreationParams FrameworkCreationParams `json:"frameworkCreationParams,omitempty"`
	BlockNumber             string                  `json:"blockNumber,omitempty"`
	BlockTimestamp          string                  `json:"blockTimestamp,omitempty"`
}

type FrameworkCreationParams struct {
	ID                           string   `json:"id,omitempty"`
	ActivationParamDefaultValues []string `json:"activationParamDefaultValues,omitempty"`
	ActivationParamVerifiers     []string `json:"activationParamVerifiers,omitempty"`
	DefaultNeedsActivation       bool     `json:"defaultNeedsActivation,omitempty"`
	LinkParentParamDefaultValues []string `json:"linkParentParamDefaultValues,omitempty"`
	LinkParentParamVerifiers     []string `json:"linkParentParamVerifiers,omitempty"`
	MintingParamDefaultValues    []string `json:"mintingParamDefaultValues,omitempty"`
	MintingParamVerifiers        []string `json:"mintingParamVerifiers,omitempty"`
	LicenseUrl                   string   `json:"licenseUrl,omitempty"`
}

type LicenseFrameworkTheGraphResponse struct {
	LicenseFramework *LicenseFramework `json:"licenseFramework"`
}

type LicenseFrameworksTheGraphResponse struct {
	LicenseFrameworks []*LicenseFramework `json:"licenseFrameworks"`
}

type LicenseFrameworkResponse struct {
	Data *LicenseFramework `json:"data"`
}

type LicensesFrameworkResponse struct {
	Data []*LicenseFramework `json:"data"`
}
