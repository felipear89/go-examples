package router

import (
	"github.com/gofiber/fiber/v2"
)

type Router struct {
	Api         fiber.Router
	Books       fiber.Router
	GoogleBooks fiber.Router
	Publishers  fiber.Router
}

func NewRouteAPI(app *fiber.App) *Router {
	api := app.Group("/api")
	return &Router{
		Api:         api,
		Books:       api.Group("/books"),
		GoogleBooks: api.Group("/googlebooks"),
		Publishers:  api.Group("/publishers"),
	}
}
