package controller

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/sidikimamsetiyawan/go-project-personal-blog/database"
	"github.com/sidikimamsetiyawan/go-project-personal-blog/model"
)

// Blog List
func BlogList(c *fiber.Ctx) error {

	context := fiber.Map{
		"statusText": "Ok",
		"msg":        "Blog List",
	}

	db := database.DBConn

	var records []model.Blog

	db.Find(&records)

	context["blog_records"] = records

	c.Status(200)
	return c.JSON(context)

}

// Add a blog into database
func BlogCreate(c *fiber.Ctx) error {
	context := fiber.Map{
		"statusText": "Ok",
		"msg":        "Add a blog into database",
	}

	record := new(model.Blog)

	if err := c.BodyParser(&record); err != nil {
		log.Println("Error in parsing request.")
		context["statusText"] = ""
		context["msg"] = "Something went wrong."
	}

	result := database.DBConn.Create(record)

	if result.Error != nil {
		log.Println("error in saving data.")
		context["statusText"] = ""
		context["msg"] = "Something went wrong"
	}

	context["msg"] = "Record is saved successfully."
	context["data"] = record

	c.Status(201)
	return c.JSON(context)

}

// Update a blog
func BlogUpdate(c *fiber.Ctx) error {
	context := fiber.Map{
		"statusText": "Ok",
		"msg":        "Update Blog",
	}

	id := c.Params("id")

	var record model.Blog

	database.DBConn.First(&record, id)

	if record.ID == 0 {
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
func BlogDelete(c *fiber.Ctx) error {

	c.Status(400)

	context := fiber.Map{
		"statusText": "",
		"msg":        "",
	}

	id := c.Params("id")

	var record model.Blog

	database.DBConn.First(&record, id)

	if record.ID == 0 {
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
