package routers

import (
	"go-mongo/app/controllers"
	"go-mongo/app/middlewares"

	"github.com/gofiber/fiber/v2"
)

func Student(router fiber.Router) {
	r := router.Group("/students")

	r.Get("/", middlewares.ExampleMiddleware, controllers.StudentFind)
	r.Get("/insert", middlewares.ExampleMiddleware, controllers.StudentInsert)
}
