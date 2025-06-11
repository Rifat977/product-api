package controllers

import (
	"product-api/config"
	"product-api/models"

	"strconv"

	"github.com/gofiber/fiber/v2"
)

func CreateProduct(c *fiber.Ctx) error {
	var product models.Product
	if err := c.BodyParser(&product); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	config.DB.Create(&product)
	return c.JSON(product)
}

func GetProducts(c *fiber.Ctx) error {
	var products []models.Product
	query := config.DB.Model(&models.Product{})

	if cat := c.Query("category"); cat != "" {
		query = query.Where("category = ?", cat)
	}
	if min := c.Query("min_price"); min != "" {
		query = query.Where("price >= ?", min)
	}
	if max := c.Query("max_price"); max != "" {
		query = query.Where("price <= ?", max)
	}

	limit := c.QueryInt("limit", 50)
	offset := c.QueryInt("offset", 0)

	query = query.Limit(limit).Offset(offset)

	if err := query.Find(&products).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(products)
}

func GetProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	var product models.Product
	result := config.DB.First(&product, id)
	if result.Error != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Product not found"})
	}
	return c.JSON(product)
}

func UpdateProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	var product models.Product
	if err := config.DB.First(&product, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Product not found"})
	}

	var update models.Product
	if err := c.BodyParser(&update); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	config.DB.Model(&product).Updates(update)
	return c.JSON(product)
}

func DeleteProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	result := config.DB.Delete(&models.Product{}, id)
	if result.RowsAffected == 0 {
		return c.Status(404).JSON(fiber.Map{"error": "Product not found"})
	}
	return c.SendStatus(204)
}

func BulkInsertProducts(c *fiber.Ctx) error {
	products := make([]models.Product, 0, 100000)

	for i := 1; i <= 100000; i++ {
		p := models.Product{
			Name:        "Product " + strconv.Itoa(i),
			Description: "Description for product " + strconv.Itoa(i),
			Price:       float64(i) * 1.5,
			Category:    "Category " + strconv.Itoa(i%10), // 10 categories
		}
		products = append(products, p)
	}

	// Batch insert with chunk size 1000
	if err := config.DB.CreateInBatches(products, 1000).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "Inserted 100000 products successfully"})
}
