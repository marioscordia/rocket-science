package app

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/marioscordia/rocket-science/order/internal/config"
	"github.com/marioscordia/rocket-science/order/internal/handler"
	"github.com/marioscordia/rocket-science/order/internal/handler/http"
	"github.com/marioscordia/rocket-science/order/internal/migrator"
	"github.com/marioscordia/rocket-science/order/internal/repository"
	"github.com/marioscordia/rocket-science/order/internal/repository/postgres"
	"github.com/marioscordia/rocket-science/order/internal/service/inventory"
	"github.com/marioscordia/rocket-science/order/internal/service/payment"
	"github.com/marioscordia/rocket-science/order/internal/usecase"
	order_v1 "github.com/marioscordia/rocket-science/shared/pkg/openapi/order/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type container struct {
	server *order_v1.Server

	handler *http.Handler

	usecase handler.UseCase

	repo usecase.Repo

	pool *pgxpool.Pool

	paymentSvc   usecase.PaymentService
	inventorySvc usecase.InventoryService
}

const migrationsDir = "./migrations"

func NewContainer() *container {
	c := &container{}
	return c
}

func (c *container) OpenAPIServer(ctx context.Context) *order_v1.Server {
	if c.server == nil {
		srv, err := order_v1.NewServer(c.Handler(ctx))
		if err != nil {
			panic(fmt.Errorf("failed to create openapi server: %w", err))
		}

		c.server = srv
	}

	return c.server
}

func (c *container) Handler(ctx context.Context) *http.Handler {
	if c.handler == nil {
		c.handler = http.NewHandler(c.UseCase(ctx))
	}
	return c.handler
}

func (c *container) UseCase(ctx context.Context) handler.UseCase {
	if c.usecase == nil {
		c.usecase = usecase.NewOrderUseCase(
			c.Repository(ctx),
			c.PaymentService(ctx),
			c.InventoryService(ctx),
		)
	}
	return c.usecase
}

func (c *container) Repository(ctx context.Context) usecase.Repo {
	if c.repo == nil {
		pool := c.GetPool(ctx)

		if err := migrator.Up(pool, migrationsDir); err != nil {
			panic(fmt.Errorf("failed to run migrations: %w", err))
		}

		store, err := postgres.NewDB(pool)
		if err != nil {
			panic(fmt.Errorf("failed to connect to postgre: %w", err))
		}

		c.repo = repository.NewRepo(store)
	}

	return c.repo
}

func (c *container) GetPool(ctx context.Context) *pgxpool.Pool {
	if c.pool == nil {
		pool, err := pgxpool.New(ctx, config.AppConfig().Postgre.GetURL())
		if err != nil {
			panic(fmt.Errorf("failed to create pgx pool: %w", err))
		}

		c.pool = pool
	}

	return c.pool
}

func (c *container) PaymentService(ctx context.Context) usecase.PaymentService {
	if c.paymentSvc == nil {
		conn := GetGRPCConnection(ctx, config.AppConfig().PaymentSvc.GetAddress(), "payment service")

		c.paymentSvc = payment.NewService(conn)
	}

	return c.paymentSvc
}

func (c *container) InventoryService(ctx context.Context) usecase.InventoryService {
	if c.inventorySvc == nil {
		conn := GetGRPCConnection(ctx, config.AppConfig().InventorySvc.GetAddress(), "inventory service")

		c.inventorySvc = inventory.NewService(conn)
	}

	return c.inventorySvc
}

func GetGRPCConnection(ctx context.Context, target, svcName string) grpc.ClientConnInterface {
	conn, err := grpc.NewClient(
		target,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		panic(fmt.Errorf("failed to create grpc connection to %s: %w", svcName, err))
	}

	return conn
}
