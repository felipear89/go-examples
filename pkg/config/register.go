package config

import "go.uber.org/fx"

func Register() fx.Option {
	return fx.Provide(NewConfig, NewServer, NewDatabase)
}
