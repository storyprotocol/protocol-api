package betav0

type RenTransaction struct {
	ID           string `json:"id"`
	CreatedAt    string `json:"created_at"`
	ActionType   string `json:"action_type"`
	Initiator    string `json:"initiator"`
	IpId         string `json:"ip_id"`
	ResourceId   string `json:"resource_id"`
	ResourceType string `json:"resource_type"`
}

type RenTransactionTheGraphResponse struct {
	Transaction *RenTransaction `json:"record"`
}

type RenTransactionsTheGraphResponse struct {
	Transactions []*RenTransaction `json:"records"`
}

type RenTransactionResponse struct {
	Data *RenTransaction `json:"data"`
}

type RenTransactionsResponse struct {
	Data []*RenTransaction `json:"data"`
}
