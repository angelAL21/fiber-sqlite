package main

import (
	"log"

	"github.com/angelAL21/fiber/database"
	"github.com/angelAL21/fiber/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	database.ConnectDb()
	app := fiber.New() //starting.

	setUpRoutes(app)
	log.Fatal(app.Listen(":3000"))
}

func welcome(c *fiber.Ctx) error {
	return c.SendString("Welcome to api with fiber")
}

//user endpoints
func setUpRoutes(app *fiber.App) {
	//welcome
	app.Get("/api", welcome)
	//User create users
	app.Post("/api/users", routes.CreateUser)
	//getallusers
	app.Get("/api/users", routes.GetUsers)
	//getuserbyid
	app.Get("/api/users/:id", routes.GetUser)
	//updateuser
	app.Put("/api/users/:id", routes.UpdateUser)
	//deleteUser
	app.Delete("/api/users/:id", routes.DeleteUser)

	//product endpoints
	//create product
	app.Post("/api/products", routes.CreateProduct)
	//get list of GetProducts
	app.Get("/api/products", routes.GetProducts)
	//getproductbyid
	app.Get("/api/products/:id", routes.GetProduct)
	//updateproduct
	app.Put("/api/products/:id", routes.UpdateProduct)
	//deleteproduct
	app.Delete("/api/products/:id", routes.DeleteProduct)

	//order endpoints
	app.Post("/api/orders", routes.CreateOrder)
	//getallorder
	app.Get("/api/orders", routes.GetOrders)
	//getorderbyid
	app.Get("/api/orders/:id", routes.GetOrder)

}
