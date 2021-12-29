package books

import (
	"github.com/felipear89/go-examples/pkg/api/router"
	"github.com/felipear89/go-examples/pkg/model"
	booksRepo "github.com/felipear89/go-examples/pkg/repository/books"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/dig"
)

type Params struct {
	dig.In
	GetAll booksRepo.GetAll
	Create booksRepo.Create
}

func Register(r *router.Router, p Params) {
	r.Books.Get("/", getBooks(p.GetAll))
	r.Books.Post("/", postBook(p.Create))
}

func postBook(create booksRepo.Create) fiber.Handler {
	return func(c *fiber.Ctx) error {

		book := new(model.Book)

		if err := c.BodyParser(book); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&router.ErrorResponse{Error: err.Error()})
		}

		if err := ValidateBook(book); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&router.ErrorResponse{Validations: err})
		}

		if err := create(book); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(&router.ErrorResponse{Error: err.Error()})
		}

		return c.Status(fiber.StatusCreated).JSON(book)
	}
}

func getBooks(getAll booksRepo.GetAll) func(ctx *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {

		books, err := getAll()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(&router.ErrorResponse{Error: err.Error()})
		}

		return c.Status(fiber.StatusOK).JSON(&ListResponse{Items: books})
	}
}

func ValidateBook(book *model.Book) []*router.BodyValidationResponse {
	var errors []*router.BodyValidationResponse
	validate := validator.New()
	err := validate.Struct(book)
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
