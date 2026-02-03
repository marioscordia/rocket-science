package main

import (
	"context"
	"fmt"
	"os/signal"
	"syscall"
	"time"

	"github.com/marioscordia/rocket-science/inventory/internal/app"
	"github.com/marioscordia/rocket-science/inventory/internal/config"
	"github.com/marioscordia/rocket-science/platform/pkg/closer"
	"github.com/marioscordia/rocket-science/platform/pkg/logger"
	"go.uber.org/zap"
)

const configPath = "../deploy/compose/inventory/.env"

func main() {
	if err := config.Load(configPath); err != nil {
		panic(fmt.Sprintf("failed to load config: %v", err))
	}

	appCtx, appCancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer appCancel()
	defer gracefulShutdown()

	closer.Configure(syscall.SIGINT, syscall.SIGTERM)

	a, err := app.NewApp(appCtx)
	if err != nil {
		logger.Error(appCtx, "❌ Не удалось создать приложение", zap.Error(err))
		return
	}

	err = a.Run(appCtx)
	if err != nil {
		logger.Error(appCtx, "❌ Ошибка при работе приложения", zap.Error(err))
		return
	}
}

func gracefulShutdown() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := closer.CloseAll(ctx); err != nil {
		logger.Error(ctx, "❌ Ошибка при завершении работы", zap.Error(err))
	}
}
