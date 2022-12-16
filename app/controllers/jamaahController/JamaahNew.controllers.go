package jamaahcontroller

import (
	"go-mongo/app/controllers/jamaahController/request"
	"go-mongo/app/models"
	"go-mongo/config"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gookit/validate"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func JamaahNew(c *fiber.Ctx) error {
	var jamaah models.Jamaah
	var res models.Response

	p := new(request.JamaahNewValidateForm)
	if err := c.BodyParser(p); err != nil {
		log.Println(err)
		res.Status = http.StatusBadRequest
		res.Message = "Empty payloads"

		return c.Status(res.Status).JSON(res)
	}

	v := validate.Struct(p)
	if !v.Validate() {
		res.Status = http.StatusBadRequest
		res.Message = v.Errors.One()

		return c.Status(res.Status).JSON(res)
	}

	db, _ := config.Connect()

	collection := db.Collection(models.JamaahCollectionName)

	_, errIndex := collection.Indexes().CreateOne(
		c.Context(),
		mongo.IndexModel{
			Keys:    bson.D{{Key: "uuid", Value: 1}},
			Options: options.Index().SetUnique(true),
		},
	)

	if errIndex != nil {
		log.Println(errIndex)

		res.Status = http.StatusBadRequest
		res.Message = "Jamaah already exist"

		return c.Status(res.Status).JSON(res)
	}

	// force MongoDB to always set its own generated ObjectIDs
	jamaah.ID = ""
	jamaah.Name = p.Name
	jamaah.UUID = p.UUID

	// insert the record
	insertionResult, err := collection.InsertOne(c.Context(), jamaah)
	if err != nil {
		log.Println(err)
		res.Status = http.StatusInternalServerError
		res.Message = "Something went wrong"

		return c.Status(res.Status).JSON(res)
	}

	// get the just inserted record in order to return it as response
	filter := bson.D{{Key: "_id", Value: insertionResult.InsertedID}}
	createdRecord := collection.FindOne(c.Context(), filter)

	// decode the Mongo record into Jamaah
	createdJamaah := &jamaah
	createdRecord.Decode(createdJamaah)

	res.Status = http.StatusOK
	res.Message = "success"
	res.Data = createdJamaah

	return c.Status(200).JSON(res)
}
