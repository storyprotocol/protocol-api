package betav0

type Module struct {
	ID             string `json:"id,omitempty"`
	Name           string `json:"name,omitempty"`
	Module         string `json:"module,omitempty"`
	BlockNumber    string `json:"blockNumber,omitempty"`
	BlockTimestamp string `json:"blockTimestamp,omitempty"`
	DeletedAt      string `json:"deletedAt,omitempty"`
}

type ModuleTheGraphResponse struct {
	Module *Module `json:"module"`
}

type ModulesTheGraphResponse struct {
	Modules []*Module `json:"modules"`
}

type ModuleResponse struct {
	Data *Module `json:"data"`
}

type ModulesResponse struct {
	Data []*Module `json:"data"`
}
