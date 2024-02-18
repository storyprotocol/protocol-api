package betav0

type LicenseMintingFeePaid struct {
	ID             string `json:"id"`
	ReceiverIpId   string `json:"receiverIpId"`
	Payer          string `json:"payer"`
	Token          string `json:"token"`
	Amount         string `json:"amount"`
	BlockNumber    string `json:"blockNumber"`
	BlockTimestamp string `json:"blockTimestamp"`
}

type LicenseMintingFeePaidTheGraphResponse struct {
	LicenseMintingFeePaid *LicenseMintingFeePaid `json:"licenseMintingFeePaidEntity"`
}

type LicenseMintingFeePaidsTheGraphResponse struct {
	LicenseMintingFeePaids []*LicenseMintingFeePaid `json:"licenseMintingFeePaidEntities"`
}

type LicenseMintingFeePaidResponse struct {
	Data *LicenseMintingFeePaid `json:"data"`
}

type LicenseMintingFeePaidsResponse struct {
	Data []*LicenseMintingFeePaid `json:"data"`
}

type LicenseMintingFeePaidRequestBody struct {
	Options *LicenseMintingFeePaidQueryOptions `json:"options"`
}
type LicenseMintingFeePaidQueryOptions struct {
	Pagination struct {
		Offset int `json:"offset"`
		Limit  int `json:"limit"`
	} `json:"pagination"`
	Where struct {
		IPID string `json:"ipId"`
	} `json:"where"`
	OrderBy        string `json:"orderBy"`
	OrderDirection string `json:"orderDirection"`
}
