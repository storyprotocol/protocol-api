package betav0

type RoyaltyPay struct {
	ID             string `json:"id"`
	ReceiverIpId   string `json:"receiverIpId"`
	PayerIpId      string `json:"payerIpId"`
	Sender         string `json:"sender"`
	Token          string `json:"token"`
	Amount         string `json:"amount"`
	BlockNumber    string `json:"blockNumber"`
	BlockTimestamp string `json:"blockTimestamp"`
}

type RoyaltyPayTheGraphResponse struct {
	RoyaltyPay *RoyaltyPay `json:"royaltyPay"`
}

type RoyaltyPaysTheGraphResponse struct {
	RoyaltyPays []*RoyaltyPay `json:"royaltyPays"`
}

type RoyaltyPayResponse struct {
	Data *RoyaltyPay `json:"data"`
}

type RoyaltyPaysResponse struct {
	Data []*RoyaltyPay `json:"data"`
}

type RoyaltyPayRequestBody struct {
	Options *RoyaltyPayQueryOptions `json:"options"`
}

type RoyaltyPayQueryOptions struct {
	Pagination struct {
		Offset int `json:"offset"`
		Limit  int `json:"limit"`
	} `json:"pagination"`
	Where struct {
		ReceiverIpId string `json:"receiverIpId"`
		PayerIpId    string `json:"payerIpId"`
		Sender       string `json:"sender"`
		Token        string `json:"token"`
	} `json:"where"`
	OrderBy        string `json:"orderBy"`
	OrderDirection string `json:"orderDirection"`
}
