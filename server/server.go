package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	"github.com/sidikimamsetiyawan/go-project-personal-blog/database"
	"github.com/sidikimamsetiyawan/go-project-personal-blog/router"
)

func init() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error in loading .env file.")
	}
	// Part Connection Database
	database.ConnectDB()
}

func main() {

	// START - Part automaticly close connection
	sqlDb, err := database.DBConn.DB()

	if err != nil {
		panic("Error in sql connection.")
	}

	defer sqlDb.Close()

	// END - Part automaticly close connection

	app := fiber.New()
	/*
		app.Get("/", func(c *fiber.Ctx) error {

			return c.JSON(fiber.Map{"message": "Welcome to my first web application"})
			// return c.SendString("Hello World.")
		})
	*/
	app.Use(logger.New())

	router.SetupRoutes(app)

	app.Listen(":8000")
}
