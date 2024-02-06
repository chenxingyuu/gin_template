package auth

type PasswordLoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type Token struct {
	Access  string `json:"access"`
	Refresh string `json:"refresh"`
	Type    string `json:"type"`
}
