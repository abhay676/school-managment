package app

import "github.com/gofiber/fiber/v2"

func (s *Server) InitV1Router() fiber.Router {
	return s.app.Group("/api/v1/gk")
}

func (s *Server) InitRoutes(router fiber.Router) {
	router.Post("/create", s.Create)
}
