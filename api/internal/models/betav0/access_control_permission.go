package betav0

//type GetAccessControlPermissionRequest struct {
//	Offset int `json:"offset"`
//	Limit  int `json:"limit"`
//}

type AccessControlPermission struct {
	IpAccount      string `json:"sender"`
	Signer         string `json:"sender"`
	To             string `json:"sender"`
	Func           string `json:"id"`
	Permission     string `json:"fwCreation"`
	BlockNumber    string `json:"blockNumber"`
	BlockTimestamp string `json:"blockTimestamp"`
}

type AccessControlPermissionTheGraphResponse struct {
	AccessControlPermissions []*AccessControlPermission `json:"licenses"`
}

type AccessControlPermissionResponse struct {
	Data []*AccessControlPermission `json:"data"`
}

type AccessControlPermissionsRequestBody struct {
	Options *ACPQueryOptions `json:"options"`
}

type ACPQueryOptions struct {
	Pagination struct {
		Offset int `json:"offset"`
		Limit  int `json:"limit"`
	} `json:"pagination"`
	Where struct {
		Name   string `json:"name"`
		Module string `json:"module"`
	} `json:"where"`
	OrderBy        string `json:"orderBy"`
	OrderDirection string `json:"orderDirection"`
}

type ACPTheGraphQueryOptions struct {
	First          int
	Skip           int
	OrderBy        string
	OrderDirection string
	Where          struct {
		Name   string
		Module string
	}
}
