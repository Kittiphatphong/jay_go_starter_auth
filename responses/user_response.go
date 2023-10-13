package responses

type LoginResponse struct {
	Token string `json:"token"`
}

type UserResponse struct {
	Name     string `json:"name"`
	Username string `json:"username"`
}
