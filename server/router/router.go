package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sidikimamsetiyawan/go-project-inventory/controller"
)

// Setup routing information
func SetupRoutes(app *fiber.App) {
	// List =>
	// Add => Post
	// Update => Put
	// Delete => Delete

	app.Get("/", controller.CategoryList)
	app.Post("/", controller.CategoryCreate)
	app.Put("/:id", controller.CategoryUpdate)
	app.Delete("/:id", controller.CategoryDelete)

	app.Get("/stocks", controller.StockList)
	app.Post("/stocks", controller.StockCreate)
	app.Put("/stocks/:id", controller.StockUpdate)
	app.Delete("/stocks/:id", controller.StockDelete)

	app.Get("/products", controller.ProductList)
	app.Post("/products", controller.ProductCreate)
}
