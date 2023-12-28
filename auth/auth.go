package auth

type Auth struct {
	AuthId       *string `json:"auth-id,omitempty"`
	SubAuthId    *int    `json:"sub-auth-id,omitempty"`
	SubAuthUser  *string `json:"sub-auth-user,omitempty"`
	AuthPassword *string `json:"auth-password,omitempty"`
}
