package publishers

import (
	"github.com/felipear89/go-examples/pkg/api/router"
	"github.com/felipear89/go-examples/pkg/model"
	"github.com/felipear89/go-examples/pkg/repository/publishers"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/dig"
)

type Params struct {
	dig.In
	Create publishers.Create
	GetAll publishers.GetAll
}

func Register(r *router.Router, p Params) {
	r.Publishers.Get("/", getPublishers(p.GetAll))
	r.Publishers.Post("/", postPublisher(p.Create))
}

func postPublisher(create publishers.Create) fiber.Handler {
	return func(c *fiber.Ctx) error {
		publisher := new(model.Publisher)
		if err := c.BodyParser(publisher); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&router.ErrorResponse{Error: err.Error()})
		}
		if err := ValidatePublisher(publisher); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&router.ErrorResponse{Validations: err, Error: "Validation error"})
		}
		if err := create(publisher); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(&router.ErrorResponse{Error: err.Error()})
		}
		return c.Status(fiber.StatusCreated).JSON(publisher)
	}
}

func getPublishers(getAll publishers.GetAll) fiber.Handler {
	return func(c *fiber.Ctx) error {
		p, err := getAll()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
		return c.JSON(p)
	}
}

func ValidatePublisher(p *model.Publisher) []*router.BodyValidationResponse {
	var errors []*router.BodyValidationResponse
	validate := validator.New()
	err := validate.Struct(p)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element router.BodyValidationResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}
