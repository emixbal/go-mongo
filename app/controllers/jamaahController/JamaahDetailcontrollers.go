package jamaahcontroller

import (
	"go-mongo/app/models"
	"go-mongo/config"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func JamaahDetail(c *fiber.Ctx) error {
	var jamaah models.Jamaah
	var res models.Response

	jamaahUuid := c.Params("uuid")

	db, _ := config.Connect()
	collection := db.Collection(models.JamaahCollectionName)

	filter := bson.D{{Key: "uuid", Value: jamaahUuid}}
	if err := collection.FindOne(c.Context(), filter).Decode(&jamaah); err != nil {
		res.Status = http.StatusNotFound
		res.Message = "not found"

		return c.Status(res.Status).JSON(res)
	}

	res.Status = http.StatusOK
	res.Message = "success"
	res.Data = jamaah

	return c.Status(res.Status).JSON(res)
}
