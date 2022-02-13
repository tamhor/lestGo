package app

import (
	"log"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/tamhor/lestGo/database"
	"github.com/tamhor/lestGo/helper"
)

type ReqGetUser struct {
	Username string `validate:"required" json:"username"`
}

type ReqAddUser struct {
	Username string `validate:"required,min=6,max=32" json:"username"`
	Password string `validate:"required,min=8,max=32" json:"password"`
	Email    string `validate:"required,email,min=6,max=32" json:"email"`
}

type ReqUpdateUser struct {
	Uuid     uuid.UUID `validate:"required" json:"uuid"`
	Username string    `json:"username"`
	Password string    `json:"password"`
	Email    string    `json:"email"`
}

type ReqDeleteUser struct {
	Uuid uuid.UUID `validate:"required" json:"uuid"`
}

func GetUsers(c *fiber.Ctx) error {
	var user []User
	res := database.DB.Find(&user)
	if res.RowsAffected == 0 {
		return c.JSON(fiber.Map{
			"status": true,
			"data":   nil,
			"error":  "User not found",
		})
	}
	if res.Error != nil {
		log.Println(res.Error)
	}
	return c.JSON(fiber.Map{
		"status": true,
		"data":   user,
		"error":  nil,
	})
}

func GetUser(c *fiber.Ctx) error {
	var user User
	request := new(ReqGetUser)
	if errors := helper.ValidateStruct(c, request); errors != nil {
		return c.JSON(errors)
	}

	res := database.DB.Where("username = ?", request.Username).First(&user)

	if res.RowsAffected == 0 {
		return c.JSON(fiber.Map{
			"status": true,
			"data":   nil,
			"error":  "User not found",
		})
	}

	if res.Error != nil {
		log.Println(res.Error)
	}

	return c.JSON(fiber.Map{
		"status": true,
		"data":   user,
		"error":  nil,
	})
}

func AddUser(c *fiber.Ctx) error {
	var user User
	request := new(ReqAddUser)
	if errors := helper.ValidateStruct(c, request); errors != nil {
		return c.JSON(errors)
	}

	if database.DB.Where("username = ?", strings.ToLower(request.Username)).First(&user).RowsAffected == 0 {
		newUser := User{
			Uuid:     uuid.New(),
			Username: strings.ToLower(request.Username),
			Password: request.Password,
			Email:    request.Email,
		}

		if database.DB.Create(&newUser).RowsAffected == 0 {
			return c.JSON(fiber.Map{
				"status": false,
				"data":   nil,
				"error":  "Failed on Add User",
			})
		}

		return c.JSON(fiber.Map{
			"status": true,
			"data":   &newUser,
			"error":  nil,
		})
	} else {
		return c.JSON(fiber.Map{
			"status": false,
			"data":   nil,
			"error":  user.Username + " is already exist",
		})
	}
}

func UpdateUser(c *fiber.Ctx) error {
	var user User
	request := new(ReqUpdateUser)
	if errors := helper.ValidateStruct(c, request); errors != nil {
		return c.JSON(errors)
	}

	if database.DB.Where("uuid = ?", request.Uuid).First(&user).RowsAffected == 1 {
		updateUser := User{
			Username: strings.ToLower(request.Username),
			Password: request.Password,
			Email:    request.Email,
		}
		res := database.DB.Where("uuid = ?", request.Uuid).Updates(&updateUser)
		if res.RowsAffected == 0 {
			return c.JSON(fiber.Map{
				"status": false,
				"data":   nil,
				"error":  "Failed to Update User",
			})
		}

		if res.Error != nil {
			log.Println(res.Error)
		}

	} else {
		return c.JSON(fiber.Map{
			"status": false,
			"data":   nil,
			"error":  "User not Found",
		})
	}
	return c.JSON(fiber.Map{
		"status": true,
		"data":   "Success update user",
		"error":  nil,
	})
}

func DeleteUser(c *fiber.Ctx) error {
	var user User
	request := new(ReqDeleteUser)
	if errors := helper.ValidateStruct(c, request); errors != nil {
		return c.JSON(errors)
	}

	if database.DB.Where("uuid = ?", request.Uuid).First(&user).RowsAffected == 1 {
		res := database.DB.Where("uuid = ?", request.Uuid).Delete(&user)
		if res.RowsAffected == 0 {
			return c.JSON(fiber.Map{
				"status": false,
				"data":   nil,
				"error":  "Failed to Delete User",
			})
		}

		if res.Error != nil {
			log.Println(res.Error)
		}

	} else {
		return c.JSON(fiber.Map{
			"status": false,
			"data":   nil,
			"error":  "User not Found",
		})
	}
	return c.JSON(fiber.Map{
		"status": true,
		"data":   "Success Delete user",
		"error":  nil,
	})
}
