package betav0

type Transaction struct {
	ID           string `json:"id"`
	CreatedAt    string `json:"createdAt"`
	ActionType   string `json:"actionType"`
	Initiator    string `json:"initiator"`
	IpId         string `json:"ipId"`
	ResourceId   string `json:"resourceId"`
	ResourceType string `json:"resourceType"`
}

type TransactionTheGraphResponse struct {
	Transaction *Transaction `json:"transaction"`
}

type TransactionsTheGraphResponse struct {
	Transactions []*Transaction `json:"transactions"`
}

type TransactionResponse struct {
	Data *Transaction `json:"data"`
}

type TransactionsResponse struct {
	Data []*Transaction `json:"data"`
}

type TransactionsQueryOptions struct {
	Creator       string
	Receiver      string
	TokenContract string
	FrameworkId   string
	LicensorIpId  string
	IPID          string
}
type TransactionRequestBody struct {
	Options *TrxQueryOptions `json:"options"`
}

type TrxQueryOptions struct {
	Pagination struct {
		Offset int `json:"offset"`
		Limit  int `json:"limit"`
	} `json:"pagination"`
	Where struct {
		ActionType string `json:"actionType"`
		ResourceId string `json:"resourceId"`
		IPID       string `json:"ipId"`
	} `json:"where"`
	OrderBy        string `json:"orderBy"`
	OrderDirection string `json:"orderDirection"`
}
