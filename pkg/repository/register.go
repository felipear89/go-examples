package repository

import (
	"github.com/felipear89/go-examples/pkg/repository/books"
	"github.com/felipear89/go-examples/pkg/repository/publishers"
	"go.uber.org/fx"
)

func Register() fx.Option {
	return fx.Provide(books.NewGetAll, books.NewCreate, publishers.NewCreate, publishers.NewGetAll)
}
