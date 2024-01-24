package entity

type Module struct {
	Name   string `json:"name,omitempty"`
	Module string `json:"module,omitempty"`
}

type ModuleTheGraphResponse struct {
	Module []*Module `json:"module"`
}

type ModulesTheGraphResponse struct {
	Modules []*Module `json:"modules"`
}

type ModuleResponse struct {
	Data []*Module `json:"data"`
}
