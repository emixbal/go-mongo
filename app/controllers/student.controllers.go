package controllers

import (
	"context"
	"go-mongo/app/models"
	"go-mongo/config"
	"log"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

var ctx = context.Background()

func StudentFind(c *fiber.Ctx) error {
	var students []models.Student

	db, err := config.Connect()

	if err != nil {
		log.Fatal(err.Error())
	}

	cursor, err := db.Collection("student").Find(ctx, bson.M{})

	if err != nil {
		panic(err)
	}

	if err = cursor.All(ctx, &students); err != nil {
		panic(err)
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "ok",
		"ctx":     ctx,
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

	_, err = db.Collection("student").InsertOne(ctx, student)

	if err != nil {
		log.Fatal(err.Error())
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "success",
	})
}
