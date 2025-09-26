package auth

import (
	"cashier-be/pkg/models"
	"cashier-be/src/user"
)

type JwtToken struct {
	User  user.UserProfileData `json:"user"`
	Token string               `json:"token"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegisterRequest struct {
	Name     string      `json:"name"`
	Username string      `json:"username"`
	Password string      `json:"password"`
	Role     models.Role `json:"role"`
}
