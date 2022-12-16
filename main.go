package main

import (
	"go-mongo/app/routers"
	"go-mongo/config"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	errEnv := godotenv.Load()
	if errEnv != nil {
		log.Fatal("Error loading .env file")
	}

	app := fiber.New()

	_, errMg := config.Connect()
	if errMg != nil {
		log.Fatal("Mongo err")
	}

	routers.Init(app)
	app.Listen(":" + os.Getenv("APP_PORT"))
}
