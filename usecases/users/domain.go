package users

type AuthenticationResponse struct {
	Success int    `json:"success"`
	Message string `json:"message"`
}
