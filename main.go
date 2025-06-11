package main

import (
	"github.com/gofiber/fiber/v2"
	"product-api/config"
	"product-api/models"
	"product-api/routes"
)

func main() {
	app := fiber.New()

	config.ConnectDB()
	config.DB.AutoMigrate(&models.Product{})

	routes.RegisterProductRoutes(app)

	app.Listen(":3000")
}
