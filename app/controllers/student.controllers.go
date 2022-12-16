package controllers

import (
	"go-mongo/app/models"
	"go-mongo/config"
	"log"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func StudentFind(c *fiber.Ctx) error {
	var students []models.Student

	db, err := config.Connect()

	if err != nil {
		log.Fatal(err.Error())
	}

	cursor, err := db.Collection("student").Find(c.Context(), bson.M{})

	if err != nil {
		panic(err)
	}

	if err = cursor.All(c.Context(), &students); err != nil {
		panic(err)
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "ok",
		"data":    students,
	})
}

func StudentInsert(c *fiber.Ctx) error {

	db, err := config.Connect()

	if err != nil {
		log.Fatal(err.Error())
	}

	student := models.Student{
		Name:  "Lentera",
		Grade: 6,
	}

	_, err = db.Collection("student").InsertOne(c.Context(), student)

	if err != nil {
		log.Fatal(err.Error())
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "success",
	})
}
