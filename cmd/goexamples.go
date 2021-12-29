package main

import (
	"context"
	"github.com/felipear89/go-examples/pkg/api"
	"github.com/felipear89/go-examples/pkg/config"
	"github.com/felipear89/go-examples/pkg/repository"
	"github.com/felipear89/go-examples/pkg/service"
	log "github.com/sirupsen/logrus"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
)

func register() *fx.App {
	app := fx.New(
		config.Register(),
		repository.Register(),
		api.Register(),
		service.Register(),
		fx.WithLogger(fxNoLogger()))
	return app
}

func fxNoLogger() func() fxevent.Logger {
	return func() fxevent.Logger {
		return fxevent.NopLogger
	}
}

func main() {
	app := register()
	ctx := context.Background()
	err := app.Start(ctx)
	if err != nil {
		log.Fatal(err)
	}

	<-app.Done()
	err = app.Stop(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
