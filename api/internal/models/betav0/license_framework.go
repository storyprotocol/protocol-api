package betav0

type LicenseFramework struct {
	Creator                 string                  `json:"creator"`
	ID                      string                  `json:"id"`
	FrameworkCreationParams FrameworkCreationParams `json:"frameworkCreationParams"`
	BlockNumber             string                  `json:"blockNumber"`
	BlockTimestamp          string                  `json:"blockTimestamp"`
}

type FrameworkCreationParams struct {
	ID                           string   `json:"id"`
	ActivationParamDefaultValues []string `json:"activationParamDefaultValues"`
	ActivationParamVerifiers     []string `json:"activationParamVerifiers"`
	DefaultNeedsActivation       bool     `json:"defaultNeedsActivation"`
	LinkParentParamDefaultValues []string `json:"linkParentParamDefaultValues"`
	LinkParentParamVerifiers     []string `json:"linkParentParamVerifiers"`
	MintingParamDefaultValues    []string `json:"mintingParamDefaultValues"`
	MintingParamVerifiers        []string `json:"mintingParamVerifiers"`
	LicenseUrl                   string   `json:"licenseUrl"`
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

type LicenseFrameworksResponse struct {
	Data []*LicenseFramework `json:"data"`
}

type LicenseFrameworkRequestBody struct {
	Options *LFWQueryOptions `json:"options"`
}

type LFWQueryOptions struct {
	Pagination struct {
		Offset int `json:"offset"`
		Limit  int `json:"limit"`
	} `json:"pagination"`
	Where struct {
		Creator string `json:"creator"`
	} `json:"where"`
	OrderBy        string `json:"orderBy"`
	OrderDirection string `json:"orderDirection"`
}
