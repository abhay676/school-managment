package app

type Createdto struct {
	Name        string `jsono:"name" validate:"required"`
	Role        string `json:"role" validate:"required"`
	Password    string `json:"password" validate:"required"`
	Email       string `json:"email" validate:"email"`
	Phone       int    `json:"phone" validate:"required"`
	CountryCode string `json:"countryCode" validate:"required"`
}
