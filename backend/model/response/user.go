package response

type UserResponse struct {
	Role      string `json:"role"`
	Username  string `json:"username"`
	Token     string `json:"token"`
	ExpiresAt int64  `json:"expiresAt"`
}

type Token struct {
	Token     string
	ExpiresAt int64
}
