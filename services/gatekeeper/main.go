package main

import (
	"fmt"
	"log"
	"os"

	"github.com/abhay676/school-managment/services/gatekeeper/app/dal"
	"github.com/abhay676/school-managment/services/gatekeeper/app/routes"
	"github.com/abhay676/school-managment/services/gatekeeper/config/database"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		panic(err.Error())
	}
}

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		panic(".env file not found")
	}

	database.Connect()

	database.Migrate(&dal.Entity{})

	app := fiber.New(fiber.Config{
		AppName: os.Getenv("APP_NAME"),
	})

	app.Use(logger.New())

	routes.AuthRoutes(app)

	log.Fatalln(app.Listen(fmt.Sprintf(":%v", os.Getenv("PORT"))))

}
