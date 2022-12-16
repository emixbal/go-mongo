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
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	app := fiber.New()

	config.Connect()

	routers.Init(app)
	app.Listen(":" + os.Getenv("APP_PORT"))
}
