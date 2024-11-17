package main

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/google/uuid"
)

var validate = validator.New()

type ErrorResponse struct {
	Status      int32                 `json:"status"`
	ErrorDetail []ErrorResponseDetail `json:"errorDetail"`
}

type ErrorResponseDetail struct {
	FieldName   string `json:"fieldName"`
	Description string `json:"description"`
}

type userCreateRequest struct {
	FirstName string `json:"firstName" validate:"required,min=2"`
	LastName  string `json:"lastName"  validate:"required"`
	Email     string `json:"email" validate:"required"`
	Password  string `json:"password" validate:"required,min=8,max=16"`
	Age       int32  `json:"age" validate:"required"`
}

type CustomValidationError struct {
	HashError bool
	Field     string
	Tag       string
	Param     string
	Value     interface{}
}

func Validate(data interface{}) []CustomValidationError {
	var customValidationError []CustomValidationError
	if errors := validate.Struct(data); errors != nil {
		for _, fieldError := range errors.(validator.ValidationErrors) {
			var cve CustomValidationError
			cve.HashError = true
			cve.Field = fieldError.Field()
			cve.Tag = fieldError.Tag()
			cve.Param = fieldError.Param()
			cve.Value = fieldError.Value()
			customValidationError = append(customValidationError, cve)
		}
	}
	return customValidationError
}

func main() {

	//fiber framework http server
	app := fiber.New()

	validationErrorDescriptionMap := map[string]string{
		"min":       "Your value should be greater than ",
		"required":  "Your value is mandatory",
		"acceptAge": "Your value should be greater than 18",
	}

	//middleware

	app.Use(func(ctx *fiber.Ctx) error {
		fmt.Printf("Hello client, you're call my %s%s AND Method: %s\n", ctx.BaseURL(), ctx.Request().RequestURI(), ctx.Request().Header.Method())
		return ctx.Next()
	})

	app.Use("/user", func(ctx *fiber.Ctx) error {
		correlationId := ctx.Get("X-CorrelationId")
		if correlationId == "" {
			return ctx.Status(http.StatusBadRequest).JSON("CorrelationId is mendatory")
		}

		_, err := uuid.Parse(correlationId)

		if err != nil {
			return ctx.Status(http.StatusBadRequest).JSON("CorrelationId is must be guid")
		}
		ctx.Locals("correlationId", correlationId)

		return ctx.Next()

	})

	app.Use(recover.New())

	/*app.Use(func(ctx *fiber.Ctx) error {
	fmt.Printf("DON'T PANIC \n")
	panic("The app was crashing...")
	return ctx.Next()})*/

	//endpoint
	app.Get("/", func(ctx *fiber.Ctx) error {
		fmt.Println("Hello first get endpoint")
		ctx.Status(200)
		return ctx.SendString("Hello first get endpoint")
	})

	app.Get("/user/:userId", func(ctx *fiber.Ctx) error {
		userIdParam := ctx.Params("userId")
		fmt.Sprintf("userId: -> %s", userIdParam)
		return ctx.SendString(fmt.Sprintf("userId: -> %s", userIdParam))
	})

	app.Post("/user", func(ctx *fiber.Ctx) error {
		fmt.Printf("User Post Endpoint\n")
		var request userCreateRequest

		err := ctx.BodyParser(&request)

		if err != nil {
			fmt.Sprintf("There was an error while binding  json - ERROR: %v\n", err.Error())
			return err
		}

		if errors := Validate(request); len(errors) > 0 && errors[0].HashError {
			var errorResponse ErrorResponse
			var errorDetailList []ErrorResponseDetail

			for _, validationError := range errors {
				errorDetailList = append(errorDetailList, ErrorResponseDetail{
					FieldName:   validationError.Field,
					Description: fmt.Sprintf("%s field has en error because %s%s", validationError.Field, validationErrorDescriptionMap[validationError.Tag], validationError.Param),
				})
			}

			errorResponse.Status = http.StatusBadRequest
			errorResponse.ErrorDetail = errorDetailList

			return ctx.Status(http.StatusBadRequest).JSON(errorResponse)
		}

		requestMessage := fmt.Sprintf(`Created succesfully: %s`, request.LastName)
		return ctx.Status(http.StatusOK).JSON(requestMessage)
	})

	app.Listen(":3000")

	//validation -> validator
}
