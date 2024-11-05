package models

type ResponseWithoutData struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

type JwtUser struct {
	UserID *int64  `json:"id"`
	Email  *string `json:"email"`
	Name   *string `json:"name"`
	Role   *string `json:"role"`
}

type Admin struct {
	Id       uint    `json:"id" db:"id"`
	Email    *string `json:"email" db:"email"`
	Password string  `json:"password" db:"password"`
	Is_admin *bool   `json:"is_admin" db:"is_admin"`
}
