package entity

//type GetAccessControlPermissionRequest struct {
//	Offset int `json:"offset"`
//	Limit  int `json:"limit"`
//}

type AccessControlPermission struct {
	IpAccount  string `json:"sender,omitempty"`
	Signer     string `json:"sender,omitempty"`
	To         string `json:"sender,omitempty"`
	Func       string `json:"id,omitempty"`
	Permission string `json:"fwCreation,omitempty"`
}

type AccessControlPermissionTheGraphResponse struct {
	AccessControlPermissions []*AccessControlPermission `json:"licenses"`
}

type AccessControlPermissionResponse struct {
	Data []*AccessControlPermission `json:"data"`
}
