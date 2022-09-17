package app

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Server struct {
	app *fiber.App
	db  *gorm.DB
}

type Config struct {
	DBURI    string
	APP_NAME string
}

func NewServer(app *fiber.App, db *gorm.DB) *Server {
	return &Server{
		app: app,
		db:  db,
	}
}

func (s *Server) Run() error {

	s.setupMiddlewares()

	if err := s.app.Listen(":3000"); err != nil {
		return err
	}
	// TODO: setup Routes
	return nil
}

func (s *Server) setupMiddlewares() {
	s.app.Use(compress.New())
	s.app.Use(cors.New())
	s.app.Use(requestid.New(
		requestid.Config{
			Header: "x-schl-mgmt",
			Generator: func() string {
				return uuid.NewString()
			},
		},
	))

	s.app.Use(logger.New(logger.Config{
		Format: "${pid} ${locals:requestid} ${status} - ${method} ${path}\n",
	}))

	s.app.Use(recover.New())
}

func GetConfig() *Config {
	return &Config{
		DBURI:    os.Getenv("DB_URI"),
		APP_NAME: os.Getenv("APP_NAME"),
	}
}
