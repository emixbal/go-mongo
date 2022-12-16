package routers

import (
	jamaahcontroller "go-mongo/app/controllers/jamaahController"

	"github.com/gofiber/fiber/v2"
)

func Jamaah(router fiber.Router) {
	r := router.Group("/jamaah")

	r.Get("/", jamaahcontroller.JamaahList)
	r.Post("/", jamaahcontroller.JamaahNew)
	r.Delete("/:id", jamaahcontroller.JamaahDelete)
	r.Get("/:uuid", jamaahcontroller.JamaahDetail)
}
