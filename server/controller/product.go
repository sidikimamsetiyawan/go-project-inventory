package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/sidikimamsetiyawan/go-project-inventory/database"
	"github.com/sidikimamsetiyawan/go-project-inventory/model"
)

// Product List
func ProductList(c *fiber.Ctx) error {

	// Ambil informasi request
	method := c.Method()
	path := c.Path()

	// Waktu mulai request
	start := time.Now()

	context := fiber.Map{
		"statusText": "Ok",
		"msg":        "Product List",
	}

	db := database.DBConn

	var records []model.ListProducts

	db.Table("products").
		Select(`products.product_id as product_id
		, products.product_name as product_name
		, categories.category_name as category_name
		, products.product_serial_number as product_serial_number
		, products.product_img as product_img
		, products.additional_info as additional_info`).
		Joins("left join categories on categories.category_id = products.category_id").
		// Where("orders.amount > ?", 100).
		Find(&records)

	// Waktu selesai dan status response
	duration := time.Since(start)
	status := c.Response().StatusCode()

	// Log format
	logJSON, err := json.MarshalIndent(records, "", "  ")
	if err != nil {
		return err
	}

	logLine := fmt.Sprintf("[%s] %s %s %d %s\nResult: %s\n",
		start.Format(time.RFC3339), method, path, status, duration, string(logJSON))

	// Tampilkan di stdout (terminal)
	fmt.Print(logLine)

	// Tampilkan di stdout (terminal)
	fmt.Print(logLine)

	logFileName := fmt.Sprintf("log_%s.txt", start.Format("2006-01-02_15-04-05"))

	// Simpan ke file log
	f, err := os.OpenFile(logFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	if _, err := f.WriteString(logLine); err != nil {
		log.Fatal(err)
	}

	context["product_records"] = records

	c.Status(200)
	return c.JSON(context)

}

// Add a product into database
func ProductCreate(c *fiber.Ctx) error {
	context := fiber.Map{
		"statusText": "Ok",
		"msg":        "Add a product data into database",
	}

	var products []model.Products

	// Parse the incoming JSON body into the slice of Category structs
	if err := c.BodyParser(&products); err != nil {
		log.Println("Error in parsing request.")
		context["statusText"] = ""
		context["msg"] = "Something went wrong."
	}

	// Print the parsed data to the console
	// printDataJSON(products)

	// Insert products into the database
	result := database.DBConn.Create(products)

	if result.Error != nil {
		log.Println("error in saving data.")
		context["statusText"] = ""
		context["msg"] = "Something went wrong"
	}

	context["msg"] = "Record is saved successfully."
	context["data"] = products

	c.Status(201)
	return c.JSON(context)

}
