package routes

import (
	"github.com/abhay676/school-managment/services/gatekeeper/app/services"
	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(app fiber.Router) {
	r := app.Group("/gk/v1/api")

	r.Post("/create", services.CreateEntity)
	r.Post("/login", services.Login)
}
