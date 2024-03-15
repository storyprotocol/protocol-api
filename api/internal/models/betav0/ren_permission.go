package betav0

type RenPermission struct {
	ID             string `json:"id"`
	UUID           string `json:"uuid,omitempty"`
	Permission     string `json:"permission"`
	Signer         string `json:"signer"`
	To             string `json:"to"`
	ToAddress      string `json:"to_address,omitempty"`
	Func           string `json:"func"`
	BlockTimestamp string `json:"block_time"`
	BlockNumber    string `json:"block_number"`
}

type RenPermissionTheGraphResponse struct {
	Permission *RenPermission `json:"record"`
}

type RenPermissionsTheGraphResponse struct {
	Permissions []*RenPermission `json:"records"`
}

type RenPermissionResponse struct {
	Data *RenPermission `json:"data"`
}

type RenPermissionsResponse struct {
	Data []*RenPermission `json:"data"`
}
