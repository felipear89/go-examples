package api

import (
	"github.com/felipear89/go-examples/pkg/api/books"
	"github.com/felipear89/go-examples/pkg/api/googlebooks"
	"github.com/felipear89/go-examples/pkg/api/publishers"
	"github.com/felipear89/go-examples/pkg/api/router"

	"go.uber.org/fx"
)

func Register() fx.Option {
	return fx.Options(
		fx.Provide(router.NewRouteAPI),
		fx.Invoke(books.Register),
		fx.Invoke(publishers.Register),
		fx.Invoke(googlebooks.Register),
	)
}
