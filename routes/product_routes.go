package routes

import (
	"product-api/controllers"

	"github.com/gofiber/fiber/v2"
)

func RegisterProductRoutes(app *fiber.App) {
	api := app.Group("/api/products")
	api.Post("/", controllers.CreateProduct)
	api.Post("/bulk-insert", controllers.BulkInsertProducts)
	api.Get("/", controllers.GetProducts)
	api.Get("/:id", controllers.GetProduct)
	api.Put("/:id", controllers.UpdateProduct)
	api.Delete("/:id", controllers.DeleteProduct)
}
