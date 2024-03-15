package betav0

type RenLicense struct {
	ID           string `json:"id"`
	PolicyID     string `json:"policy_id"`
	LicensorIpId string `json:"licensor_ip_id"`
	Transferable string `json:"transferable"`
	Amount       string `json:"amount"`

	BlockNumber    string `json:"block_number"`
	BlockTimestamp string `json:"block_time"`
}

type RenLicenseTheGraphResponse struct {
	License *RenLicense `json:"record"`
}

type RenLicensesTheGraphResponse struct {
	Licenses []*RenLicense `json:"records"`
}

type RenLicenseResponse struct {
	Data *RenLicense `json:"data"`
}

type RenLicensesResponse struct {
	Data []*RenLicense `json:"data"`
}
