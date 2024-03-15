package betav0

type RenModule struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	Module         string `json:"module"`
	BlockNumber    string `json:"block_number"`
	BlockTimestamp string `json:"block_time"`
	DeletedAt      string `json:"deleted_at"`
}

type RenModuleTheGraphResponse struct {
	Module *RenModule `json:"record"`
}

type RenModulesTheGraphResponse struct {
	Modules []*RenModule `json:"records"`
}

type RenModuleResponse struct {
	Data *RenModule `json:"data"`
}

type RenModulesResponse struct {
	Data []*RenModule `json:"data"`
}
