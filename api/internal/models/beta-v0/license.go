package beta_v0

type License struct {
	ID          string      `json:"id,omitempty"`
	Amount      string      `json:"amount,omitempty"`
	Creator     string      `json:"creator,omitempty"`
	LicenseId   string      `json:"licenseId,omitempty"`
	Receiver    string      `json:"receiver,omitempty"`
	LicenseData LicenseData `json:"licenseData,omitempty"`
}

type LicenseData struct {
	PolicyID      string   `json:"policyId,omitempty"`
	LicensorIpIds []string `json:"licensorIpIds,omitempty"`
}

type LicenseTheGraphResponse struct {
	License []*License `json:"license"`
}

type LicensesTheGraphResponse struct {
	Licenses []*License `json:"licenses"`
}

type LicenseResponse struct {
	Data []*License `json:"data"`
}
