package types

import "time"

type EntityResponse struct {
	EID       string    `json:"e_id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	IsActive  bool      `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
	Password  string    `json:"-"`
}

type CreateEntityDTO struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
	Role     string `json:"role" validate:"required"`
}

type EntityCreateResponse struct {
	Entity *EntityResponse `json:"entity"`
}

type LoginDTO struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type AccessResponse struct {
	Token string `json:"token"`
}

type LoginResponse struct {
	Entity *EntityResponse `json:"entity"`
	Auth   *AccessResponse `json:"auth"`
}
