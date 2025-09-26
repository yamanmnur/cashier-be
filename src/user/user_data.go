package user

import "cashier-be/pkg/models"

type UserData struct {
	Id       uint        `json:"id"`
	Username string      `json:"username"`
	Password string      `json:"password"`
	Name     string      `json:"name"`
	Role     models.Role `json:"role"`
}

type UserProfileData struct {
	Id       uint        `json:"id"`
	Username string      `json:"username"`
	Name     string      `json:"name"`
	Role     models.Role `json:"role"`
}
