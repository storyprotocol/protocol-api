package betav0

type RenLicenseMintingFeePaid struct {
	ID             string `json:"id"`
	ReceiverIpId   string `json:"receiver_ip_id"`
	Payer          string `json:"payer"`
	Token          string `json:"token"`
	Amount         string `json:"amount"`
	BlockNumber    string `json:"block_number"`
	BlockTimestamp string `json:"block_time"`
}

type RenLicenseMintingFeePaidTheGraphResponse struct {
	LicenseMintingFeePaid *RenLicenseMintingFeePaid `json:"record"`
}

type RenLicenseMintingFeePaidsTheGraphResponse struct {
	LicenseMintingFeePaids []*RenLicenseMintingFeePaid `json:"records"`
}

type RenLicenseMintingFeePaidResponse struct {
	Data *RenLicenseMintingFeePaid `json:"data"`
}

type RenLicenseMintingFeePaidsResponse struct {
	Data []*RenLicenseMintingFeePaid `json:"data"`
}
