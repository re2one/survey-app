package response

type UserResponse struct {
	Role     string `json:"role"`
	Username string `json:"username"`
	Token    string `json:"token"`
}
