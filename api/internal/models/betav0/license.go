package betav0

type License struct {
	ID           string `json:"id,omitempty"`
	PolicyID     string `json:"policyId,omitempty"`
	LicensorIpId string `json:"licensorIpId,omitempty"`
	Transferable bool   `json:"transferable,omitempty"`

	BlockNumber    string `json:"blockNumber,omitempty"`
	BlockTimestamp string `json:"blockTimestamp,omitempty"`
}

type LicenseTheGraphResponse struct {
	License *License `json:"license"`
}

type LicensesTheGraphResponse struct {
	Licenses []*License `json:"licenses"`
}

type LicenseResponse struct {
	Data *License `json:"data"`
}

type LicensesResponse struct {
	Data []*License `json:"data"`
}
