package betav0

type Transaction struct {
	ID             string `json:"id,omitempty"`
	UUID           string `json:"uuid,omitempty"`
	IPID           string `json:"ipId,omitempty"`
	Transaction    string `json:"tag,omitempty"`
	DeletedAt      string `json:"deletedAt,omitempty"`
	BlockNumber    string `json:"blockNumber,omitempty"`
	BlockTimestamp string `json:"blockTimestamp,omitempty"`
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
