package routes

import (
	"errors"

	"github.com/angelAL21/fiber/database"
	"github.com/angelAL21/fiber/models"
	"github.com/gofiber/fiber/v2"
)

type Product struct {
	ID          uint64 `json:"id"`
	Name        string `json:"name"`
	SerialNumer string `json:"serial_number"`
}

func CreateResponseProduct(productModel models.Product) Product {
	return Product{ID: productModel.ID, Name: productModel.Name, SerialNumer: productModel.SerialNumer}
}

//endpoint createproduct
func CreateProduct(c *fiber.Ctx) error {
	var product models.Product

	if err := c.BodyParser(&product); err != nil {
		return c.Status(400).JSON((err.Error()))
	}
	database.Database.Db.Create(&product)
	responseProdcut := CreateResponseProduct(product)

	return c.Status(200).JSON(responseProdcut)
}

//GetProducts will return all products
func GetProducts(c *fiber.Ctx) error {
	products := []models.Product{}
	database.Database.Db.Find(&products)
	responseProducts := []Product{}
	for _, product := range products {
		responseProduct := CreateResponseProduct(product)
		responseProducts = append(responseProducts, responseProduct)
	}
	return c.Status(200).JSON(responseProducts)
}

//will return one product by its ID
func GetProduct(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var product models.Product
	if err != nil {
		return c.Status(400).JSON("the product does not exist")
	}
	if err := findProduct(id, &product); err != nil {
		return c.Status(400).JSON("the product does not exist")
	}
	responseProduct := CreateResponseProduct(product)
	return c.Status(200).JSON(responseProduct)
}

//helper to GetProduct (getuserbyid)
func findProduct(id int, product *models.Product) error {
	database.Database.Db.Find(&product, "id = ?", id)
	if product.ID == 0 {
		return errors.New("product does not exist")
	}
	return nil
}

//updateProduct will update the product with name and serialnumber
func UpdateProduct(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var product models.Product
	if err != nil {
		return c.Status(400).JSON("the product does not exist")
	}
	if err := findProduct(id, &product); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	//what to update
	type UpdateProduct struct {
		Name        string `json:"name"`
		SerialNumer string `json:"serial_number"`
	}

	var UpdateData UpdateProduct

	if err := c.BodyParser(&UpdateData); err != nil {
		return c.Status(500).JSON(err.Error())
	}

	product.Name = UpdateData.Name
	product.SerialNumer = UpdateData.SerialNumer

	database.Database.Db.Save(&product)

	responseProduct := CreateResponseProduct(product)
	return c.Status(200).JSON(responseProduct)
}

//DeleteProuct will delete the product
func DeleteProduct(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var product models.Product

	if err != nil {
		return c.Status(400).JSON("the product does not exist")
	}
	if err := findProduct(id, &product); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	if err := database.Database.Db.Delete(&product).Error; err != nil {
		return c.Status(404).JSON(err.Error())
	}
	return c.Status(200).SendString("product deleted successfully")
}
