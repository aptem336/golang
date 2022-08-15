package api

type AuthRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AuthResponse struct {
	Username string `json:"username"`
	Token    string `json:"token"`
	Id       int    `json:"id"`
	Email    string `json:"email,omitempty"`
}
