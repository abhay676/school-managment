package main

import (
	"log"
	"os"

	"github.com/abhay676/school-managment/services/gatekeeper/pkg/app"
	"github.com/abhay676/school-managment/services/gatekeeper/pkg/repository"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	if err := run(); err != nil {
		log.Fatalf("Error occured in Application %s", err.Error())
		os.Exit(1)
	}
}

func run() error {

	err := godotenv.Load()
	if err != nil {
		return err
	}
	config := app.GetConfig()

	db, err := repository.SetupDBConnection(config.DBURI)

	if err != nil {
		return err
	}

	server := app.NewServer(fiber.New(fiber.Config{
		CaseSensitive: true,
		AppName:       config.APP_NAME,
	}), db)

	if err != server.Run() {
		return err
	}

	return nil
}
