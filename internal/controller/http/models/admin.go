package models

type LoginRequest struct {
	Email    string `json:"email"       example:"dealer@gmail.com"`
	Password string `json:"password"    example:"pass"`
}

type LoginResponse struct {
	AccessToken string `json:"access_token"`
}
