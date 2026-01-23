package app

import (
	"context"
	"fmt"
	"net/http"

	"github.com/marioscordia/rocket-science/order/internal/config"
	"github.com/marioscordia/rocket-science/platform/pkg/closer"
	"github.com/marioscordia/rocket-science/platform/pkg/logger"
)

type App struct {
	container  *container
	httpServer *http.Server
}

func NewApp(ctx context.Context) (*App, error) {
	app := &App{}

	if err := app.initDeps(ctx); err != nil {
		return nil, err
	}

	return app, nil
}

func (a *App) initDeps(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.initContainer,
		a.initLogger,
		a.initCloser,
		a.initHttpServer,
	}

	for _, f := range inits {
		err := f(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *App) initContainer(_ context.Context) error {
	a.container = NewContainer()

	return nil
}

func (a *App) initLogger(_ context.Context) error {
	return logger.Init(
		config.AppConfig().Logger.GetLevel(),
		config.AppConfig().Logger.AsJSON(),
	)
}

func (a *App) initCloser(_ context.Context) error {
	closer.SetLogger(logger.Logger())
	return nil
}

func (a *App) initHttpServer(ctx context.Context) error {
	a.httpServer = &http.Server{
		Addr:    config.AppConfig().GRPC.GetAddress(),
		Handler: a.container.OpenAPIServer(ctx),
	}

	return nil
}

func (a *App) Run(ctx context.Context) error {
	logger.Info(ctx, fmt.Sprintf("🚀 gRPC InventoryService server listening on %s", config.AppConfig().GRPC.GetAddress()))

	return a.httpServer.ListenAndServe()
}
