package models

type CreateUserRequest struct {
    FullName string     `json:"full_name"`
    Username string     `json:"username" validate:"required,notblank"`
    Email string        `json:"email" validate:"required,email"`
    Password string     `json:"password" validate:"required,min=8,strongpassword"`
}