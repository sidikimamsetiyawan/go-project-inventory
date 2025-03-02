package controller

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/sidikimamsetiyawan/go-project-inventory/database"
	"github.com/sidikimamsetiyawan/go-project-inventory/model"
)

// Category List
func CategoryList(c *fiber.Ctx) error {

	context := fiber.Map{
		"statusText": "Ok",
		"msg":        "Category List",
	}

	db := database.DBConn

	var records []model.Categories

	db.Find(&records)

	fmt.Println("Category List Data : ")
	fmt.Println(records)

	context["category_records"] = records

	c.Status(200)
	return c.JSON(context)

}

// Add a blog into database
func CategoryCreate(c *fiber.Ctx) error {
	context := fiber.Map{
		"statusText": "Ok",
		"msg":        "Add a category data into database",
	}

	var categories []model.Categories

	// Parse the incoming JSON body into the slice of Category structs
	if err := c.BodyParser(&categories); err != nil {
		log.Println("Error in parsing request.")
		context["statusText"] = ""
		context["msg"] = "Something went wrong."
	}

	// Print the parsed data to the console
	printDataJSON(categories)

	// Insert categories into the database
	result := database.DBConn.Create(categories)

	if result.Error != nil {
		log.Println("error in saving data.")
		context["statusText"] = ""
		context["msg"] = "Something went wrong"
	}

	context["msg"] = "Record is saved successfully."
	context["data"] = categories

	c.Status(201)
	return c.JSON(context)

}

// Update a blog
func CategoryUpdate(c *fiber.Ctx) error {
	context := fiber.Map{
		"statusText": "Ok",
		"msg":        "Update Blog",
	}

	id := c.Params("id")

	var record []model.Categories

	database.DBConn.First(&record, id)

	if record[0].CategoryID == 0 {
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
func CategoryDelete(c *fiber.Ctx) error {

	c.Status(400)

	context := fiber.Map{
		"statusText": "",
		"msg":        "",
	}

	id := c.Params("id")

	var record []model.Categories

	database.DBConn.First(&record, id)

	if record[0].CategoryID == 0 {
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

func printDataJSON(categories []model.Categories) {
	fmt.Println("Received JSON Data:")
	for i, category := range categories {
		// Marshal the category struct into a JSON string for pretty printing
		categoryJSON, _ := json.MarshalIndent(category, "", "  ")
		fmt.Printf("Category %d:\n%s\n", i+1, categoryJSON)
	}
}
