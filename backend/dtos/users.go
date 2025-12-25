package dtos

import (
	"serverless/interfaces"
)

type UserRegisterDto struct {
	interfaces.Validatable
	Username string `json:"username" validate:"required,min=3,max=20"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type UserLoginDto struct {
	interfaces.Validatable
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}
