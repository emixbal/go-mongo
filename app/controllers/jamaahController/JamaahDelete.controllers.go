package jamaahcontroller

import (
	"go-mongo/app/models"
	"go-mongo/config"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func JamaahDelete(c *fiber.Ctx) error {
	var res models.Response

	jamaahId, err := primitive.ObjectIDFromHex(
		c.Params("id"),
	)

	if err != nil {
		res.Status = http.StatusBadRequest
		res.Message = "Empty payloads"

		return c.Status(res.Status).JSON(res)
	}

	db, _ := config.Connect()

	collection := db.Collection(models.JamaahCollectionName)

	// find and delete the employee with the given ID
	query := bson.D{{Key: "_id", Value: jamaahId}}
	result, err := collection.DeleteOne(c.Context(), &query)

	if err != nil {
		log.Println(err)
		res.Status = http.StatusInternalServerError
		res.Message = "Something went wrong"

		return c.Status(res.Status).JSON(res)
	}

	// the employee might not exist
	if result.DeletedCount < 1 {
		res.Status = http.StatusNotFound
		res.Message = "Jamaah not found"

		return c.Status(res.Status).JSON(res)
	}

	res.Status = http.StatusNoContent
	res.Message = "success"

	return c.Status(200).JSON(res)
}
