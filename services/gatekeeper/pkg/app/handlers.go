package app

import (
	"fmt"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

type ErrorResponse struct {
	FailedField string `json:"field"`
	Tag         string `json:"tag"`
	Value       string `json:"value"`
}

var Validate = validator.New()

func ValidateStruct(entity Createdto) []*ErrorResponse {
	var errors []*ErrorResponse
	err := Validate.Struct(entity)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}

func (s *Server) Create(ctx *fiber.Ctx) error {

	c := new(Createdto)

	if err := ctx.BodyParser(c); err != nil {
		fmt.Println("error = ", err)
		return ctx.JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	errors := ValidateStruct(*c)

	if errors != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(errors)
	}

	return nil
}
