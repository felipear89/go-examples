package service

import (
	"github.com/felipear89/go-examples/pkg/service/googlebooks"
	"go.uber.org/fx"
)

func Register() fx.Option {
	return fx.Options(googlebooks.Register())
}
