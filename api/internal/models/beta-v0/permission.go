package beta_v0

type Permission struct {
	UUID           string `json:"uuid,omitempty"`
	IPAccount      string `json:"ipAccount,omitempty"`
	Permission     string `json:"permission,omitempty"`
	Signer         string `json:"signer,omitempty"`
	To             string `json:"to,omitempty"`
	Func           string `json:"func,omitempty"`
	BlockTimestamp string `json:"blockTimestamp,omitempty"`
	BlockNumber    string `json:"blockNumber,omitempty"`
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
