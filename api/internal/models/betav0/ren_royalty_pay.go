package betav0

type RenRoyaltyPay struct {
	ID             string `json:"id"`
	ReceiverIpId   string `json:"receiver_ip_d"`
	PayerIpId      string `json:"payer_ip_id"`
	Sender         string `json:"sender"`
	Token          string `json:"token"`
	Amount         string `json:"amount"`
	BlockNumber    string `json:"block_number"`
	BlockTimestamp string `json:"block_time"`
}

type RenRoyaltyPayTheGraphResponse struct {
	RoyaltyPay *RenRoyaltyPay `json:"records"`
}

type RenRoyaltyPaysTheGraphResponse struct {
	RoyaltyPays []*RenRoyaltyPay `json:"records"`
}

type RenRoyaltyPayResponse struct {
	Data *RenRoyaltyPay `json:"data"`
}

type RenRoyaltyPaysResponse struct {
	Data []*RenRoyaltyPay `json:"data"`
}
