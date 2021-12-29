package googlebooks

import "go.uber.org/fx"

func Register() fx.Option {
	return fx.Provide(NewClient, NewSearchBooks, NewRequestSearchBooks)
}
