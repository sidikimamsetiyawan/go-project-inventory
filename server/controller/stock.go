package controller

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/sidikimamsetiyawan/go-project-inventory/database"
	"github.com/sidikimamsetiyawan/go-project-inventory/model"
)

func StockList(c *fiber.Ctx) error {
	context := fiber.Map{
		"statusText": "Ok",
		"msg":        "Stock List",
	}

	db := database.DBConn

	var records []model.Stocks

	db.Find(&records)

	context["stock_records"] = records

	c.Status(200)
	return c.JSON(context)
}

func StockCreate(c *fiber.Ctx) error {
	context := fiber.Map{
		"statusText": "Ok",
		"msg":        "Add a stock data into database",
	}

	var stocks []model.Stocks

	// Parse the incoming JSON body into the slice of Category structs
	if err := c.BodyParser(&stocks); err != nil {
		log.Println("Error in parsing request.")
		context["statusText"] = ""
		context["msg"] = "Something went wrong."
	}

	// Print the parsed data to the console
	// printDataJSON(stocks)

	// Insert stocks into the database
	result := database.DBConn.Create(stocks)

	if result.Error != nil {
		log.Println("error in saving data.")
		context["statusText"] = ""
		context["msg"] = "Something went wrong"
	}

	context["msg"] = "Record is saved successfully."
	context["data"] = stocks

	c.Status(201)
	return c.JSON(context)

}

// Update a blog
func StockUpdate(c *fiber.Ctx) error {
	context := fiber.Map{
		"statusText": "Ok",
		"msg":        "Update Blog",
	}

	id := c.Params("id")

	var record []model.Stocks

	database.DBConn.First(&record, id)

	if record[0].StockID == 0 {
		log.Println("Record not found.")
		context[""] = ""
		context["msg"] = "Record not found."
		c.Status(400)
		return c.JSON(context)
	}

	if err := c.BodyParser(&record); err != nil {
		log.Println("Error in parsing request.")
	}

	result := database.DBConn.Save(record)

	if result.Error != nil {
		log.Println("Error saving data.")
	}

	context["msg"] = "Record update successfuly."
	context["data"] = record

	c.Status(200)
	return c.JSON(context)

}

// Delete a blog
func StockDelete(c *fiber.Ctx) error {

	c.Status(400)

	context := fiber.Map{
		"statusText": "",
		"msg":        "",
	}

	id := c.Params("id")

	var record []model.Stocks

	database.DBConn.First(&record, id)

	if record[0].StockID == 0 {
		log.Println("Record not found.")
		context["msg"] = "Record not found."

		c.Status(400)
		return c.JSON(context)
	}

	result := database.DBConn.Delete(record)

	if result.Error != nil {
		context["msg"] = "Something went wrong."
		return c.JSON(context)
	}

	context["statusText"] = "Ok."
	context["msg"] = "Record deleted successfully."

	c.Status(200)
	return c.JSON(context)

}
