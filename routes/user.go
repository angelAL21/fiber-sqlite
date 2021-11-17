package routes

import (
	"errors"

	"github.com/angelAL21/fiber/database"
	"github.com/angelAL21/fiber/models"
	"github.com/gofiber/fiber/v2"
)

type User struct {
	//this is not the model user. This is a serializer
	ID        uint64 `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func CreateResponseUser(userModel models.User) User {
	return User{ID: userModel.ID, FirstName: userModel.FirstName, LastName: userModel.LastName}
}

//endpoint
//createuser creates a new user with postman and json format
func CreateUser(c *fiber.Ctx) error {
	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	database.Database.Db.Create(&user)
	responseUser := CreateResponseUser(user)
	return c.Status(200).JSON(responseUser)
}

//getusers will get all users
func GetUsers(c *fiber.Ctx) error {
	users := []models.User{}

	database.Database.Db.Find(&users) //slice of users.
	responserUsers := []User{}
	for _, user := range users {
		responseUser := CreateResponseUser(user)
		responserUsers = append(responserUsers, responseUser)
	}
	return c.Status(200).JSON(responserUsers)
}

//GetUser will get only one user
func GetUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var user models.User
	if err != nil {
		return c.Status(400).JSON("the user does not exist")
	}
	if err := findUser(id, &user); err != nil {
		return c.Status(400).JSON("the user does not exist")
	}
	responseUser := CreateResponseUser(user)
	return c.Status(200).JSON(responseUser)
}

//helper to GetUser (getuserbyid)
func findUser(id int, user *models.User) error {
	database.Database.Db.Find(&user, "id = ?", id)
	if user.ID == 0 {
		return errors.New("user does not exist")
	}
	return nil
}

//updateUser will update the user with name and last_name
func UpdateUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var user models.User
	if err != nil {
		return c.Status(400).JSON("the user does not exist")
	}
	if err := findUser(id, &user); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	//what to update
	type UpdateUser struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
	}

	var UpdateData UpdateUser

	if err := c.BodyParser(&UpdateData); err != nil {
		return c.Status(500).JSON(err.Error())
	}

	user.FirstName = UpdateData.FirstName
	user.LastName = UpdateData.LastName

	database.Database.Db.Save(&user)

	responseUser := CreateResponseUser(user)
	return c.Status(200).JSON(responseUser)
}

//Deleteuser will delete the user
func DeleteUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var user models.User

	if err != nil {
		return c.Status(400).JSON("the user does not exist")
	}
	if err := findUser(id, &user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	if err := database.Database.Db.Delete(&user).Error; err != nil {
		return c.Status(404).JSON(err.Error())
	}
	return c.Status(200).SendString("user deleted successfully")
}
