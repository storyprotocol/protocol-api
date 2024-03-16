package betav0

type RenLicenseOwner struct {
	ID       string `json:"id"`
	PolicyId string `json:"policy_id"`
	Owner    string `json:"owner"`
	Amount   string `json:"amount"`

	BlockNumber    string `json:"block_number"`
	BlockTimestamp string `json:"block_time"`
}

type RenLicenseOwnerTheGraphResponse struct {
	LicenseOwner *RenLicenseOwner `json:"record"`
}

type RenLicenseOwnersTheGraphResponse struct {
	LicenseOwners []*RenLicenseOwner `json:"records"`
}

type RenLicenseOwnerResponse struct {
	Data *RenLicenseOwner `json:"data"`
}

type RenLicenseOwnersResponse struct {
	Data []*RenLicenseOwner `json:"data"`
}
