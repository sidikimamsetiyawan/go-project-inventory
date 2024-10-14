package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sidikimamsetiyawan/go-project-personal-blog/controller"
)

// Setup routing information
func SetupRoutes(app *fiber.App) {
	// List =>
	// Add => Post
	// Update => Put
	// Delete => Delete

	app.Get("/", controller.BlogList)
	app.Post("/", controller.BlogCreate)
	app.Put("/:id", controller.BlogUpdate)
	app.Delete("/:id", controller.BlogDelete)
}
