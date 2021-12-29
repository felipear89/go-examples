package googlebooks

import (
	"github.com/felipear89/go-examples/pkg/api/router"
	"github.com/felipear89/go-examples/pkg/service/googlebooks"

	"github.com/gofiber/fiber/v2"
)

func Register(r *router.Router, searchBooks googlebooks.SearchBooks) {
	r.GoogleBooks.Get("/", getBooks(searchBooks))
}

func getBooks(searchBooks googlebooks.SearchBooks) func(ctx *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {

		query := c.Query("q")
		if query == "" {
			return c.Status(fiber.StatusBadRequest).JSON(&router.ErrorResponse{
				Error: "query is required",
			})
		}
		books, err := searchBooks(query)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(&router.ErrorResponse{Error: err.Error()})
		}
		return c.JSON(books)

	}
}
