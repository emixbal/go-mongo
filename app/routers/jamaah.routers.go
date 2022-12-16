package routers

import (
	jamaahcontroller "go-mongo/app/controllers/jamaahController"
	"go-mongo/app/middlewares"

	"github.com/gofiber/fiber/v2"
)

func Jamaah(router fiber.Router) {
	r := router.Group("/jamaah")

	r.Get("/", middlewares.ExampleMiddleware, jamaahcontroller.JamaahList)
	r.Post("/", middlewares.ExampleMiddleware, jamaahcontroller.JamaahNew)

}
