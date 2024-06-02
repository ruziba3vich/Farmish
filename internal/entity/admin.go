// Package entity defines main entities for business logic (services), data base mapping and
// HTTP response objects if suitable. Each logic group entities in own file.
package entity

// LoginRequest -.
type LoginRequest struct {
	Email    string `json:"email"       example:"dealer@gmail.com"`
	Password string `json:"password"    example:"pass"`
}

type LoginResponse struct {
	AccessToken string `json:"access_token"`
}
