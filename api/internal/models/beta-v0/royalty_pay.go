package beta_v0

type RoyaltyPay struct {
	ID             string `json:"id,omitempty"`
	ReceiverIpId   string `json:"receiverIpId,omitempty"`
	PayerIpId      string `json:"payerIpId,omitempty"`
	Sender         string `json:"sender,omitempty"`
	Token          string `json:"token,omitempty"`
	Amount         string `json:"amount,omitempty"`
	BlockNumber    string `json:"blockNumber,omitempty"`
	BlockTimestamp string `json:"blockTimestamp,omitempty"`
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
