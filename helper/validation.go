package helper

import (
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

type ErrorResponse struct {
	Status bool        `json:"status"`
	Error  interface{} `json:"error"`
}

func ValidateStruct(c *fiber.Ctx, request interface{}) []ErrorResponse {
	var errors []ErrorResponse
	var element ErrorResponse
	var errList []interface{}
	element.Status = false

	if err := c.BodyParser(request); err != nil {
		element.Error = err.Error()
		errors = append(errors, element)
		return errors
	}

	validate := validator.New()
	if err := validate.Struct(request); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			rowError := err.Field() + " is " + err.Tag()
			errList = append(errList, rowError)
		}
		element.Error = &errList
		errors = append(errors, element)
	}
	return errors
}
