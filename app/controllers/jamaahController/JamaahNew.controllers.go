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

	collection := db.Collection(models.JamaahKey())

	// force MongoDB to always set its own generated ObjectIDs
	jamaah.ID = ""
	jamaah.Name = p.Name

	// insert the record
	insertionResult, err := collection.InsertOne(c.Context(), jamaah)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	// get the just inserted record in order to return it as response
	filter := bson.D{{Key: "_id", Value: insertionResult.InsertedID}}
	createdRecord := collection.FindOne(c.Context(), filter)

	// decode the Mongo record into Employee
	createdEmployee := &jamaah
	createdRecord.Decode(createdEmployee)

	res.Status = http.StatusOK
	res.Message = "success"
	res.Data = createdEmployee

	return c.Status(200).JSON(res)
}
