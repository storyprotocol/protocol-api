package betav0

type Module struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	Module         string `json:"module"`
	BlockNumber    string `json:"blockNumber"`
	BlockTimestamp string `json:"blockTimestamp"`
	DeletedAt      string `json:"deletedAt"`
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

type ModuleRequestBody struct {
	Options *ModuleQueryOptions `json:"options"`
}

type ModuleQueryOptions struct {
	Pagination struct {
		Offset int `json:"offset"`
		Limit  int `json:"limit"`
	} `json:"pagination"`
	Where struct {
		Name string `json:"name"`
	} `json:"where"`
	OrderBy        string `json:"orderBy"`
	OrderDirection string `json:"orderDirection"`
}
