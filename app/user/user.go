package app

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/tamhor/lestGo/database"
	"github.com/tamhor/lestGo/helper"
)

type ReqUser struct {
	Uuid uuid.UUID `validate:"required" json:"uuid"`
}

func GetUsers(c *fiber.Ctx) error {
	var user []User
	res := database.DB.Find(&user)
	if res.Error != nil {
		log.Println(res.Error)
	}
	return c.JSON(user)
}

func GetUser(c *fiber.Ctx) error {
	var user User
	request := new(ReqUser)
	if errors := helper.ValidateStruct(c, request); errors != nil {
		return c.JSON(errors)
	}

	res := database.DB.Where("uuid").First(&user)

	if res.RowsAffected == 0 {
		return c.JSON(fiber.Map{
			"status": true,
			"data":   nil,
		})
	}

	if res.Error != nil {
		log.Println(res.Error)
	}

	return c.JSON(user)
}
