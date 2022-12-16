package jamaahcontroller

import (
	"go-mongo/app/models"
	"go-mongo/config"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func JamaahList(c *fiber.Ctx) error {
	var jamaahs []models.Jamaah

	db, _ := config.Connect()

	cursor, err := db.Collection(models.JamaahKey()).Find(c.Context(), bson.M{})

	if err != nil {
		panic(err)
	}

	if err = cursor.All(c.Context(), &jamaahs); err != nil {
		panic(err)
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "ok",
		"data":    jamaahs,
	})
}
