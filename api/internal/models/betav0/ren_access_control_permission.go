package betav0

//type GetAccessControlPermissionRequest struct {
//	Offset int `json:"offset"`
//	Limit  int `json:"limit"`
//}

type RenAccessControlPermission struct {
	IpAccount      string `json:"sender"`
	Signer         string `json:"sender"`
	To             string `json:"sender"`
	Func           string `json:"id"`
	Permission     string `json:"fw_creation"`
	BlockNumber    string `json:"block_number"`
	BlockTimestamp string `json:"block_time"`
}

type RenAccessControlPermissionTheGraphResponse struct {
	AccessControlPermissions []*RenAccessControlPermission `json:"records"`
}

type RenAccessControlPermissionResponse struct {
	Data []*RenAccessControlPermission `json:"data"`
}
