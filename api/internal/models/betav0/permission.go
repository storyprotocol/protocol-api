package betav0

type Permission struct {
	ID             string `json:"id"`
	UUID           string `json:"uuid"`
	Permission     string `json:"permission"`
	Signer         string `json:"signer"`
	To             string `json:"to"`
	Func           string `json:"func"`
	BlockTimestamp string `json:"blockTimestamp"`
	BlockNumber    string `json:"blockNumber"`
}

type PermissionTheGraphResponse struct {
	Permission *Permission `json:"permission"`
}

type PermissionsTheGraphResponse struct {
	Permissions []*Permission `json:"permissions"`
}

type PermissionResponse struct {
	Data *Permission `json:"data"`
}

type PermissionsResponse struct {
	Data []*Permission `json:"data"`
}

type PermissionRequestBody struct {
	Options *PermissionQueryOptions `json:"options"`
}

type PermissionQueryOptions struct {
	Pagination struct {
		Offset int `json:"offset"`
		Limit  int `json:"limit"`
	} `json:"pagination"`
	Where struct {
		Signer string `json:"signer"`
		To     string `json:"to"`
	} `json:"where"`
	OrderBy        string `json:"orderBy"`
	OrderDirection string `json:"orderDirection"`
}
