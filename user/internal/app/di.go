package app

import (
	"context"
	"fmt"

	redigo "github.com/gomodule/redigo/redis"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/marioscordia/rocket-science/platform/pkg/closer"
	"github.com/marioscordia/rocket-science/platform/pkg/logger"
	userv1 "github.com/marioscordia/rocket-science/shared/pkg/proto/user/v1"
	"github.com/marioscordia/rocket-science/user/internal/config"
	"github.com/marioscordia/rocket-science/user/internal/handler"
	"github.com/marioscordia/rocket-science/user/internal/handler/grpc"
	"github.com/marioscordia/rocket-science/user/internal/repository"
	"github.com/marioscordia/rocket-science/user/internal/repository/postgres"
	"github.com/marioscordia/rocket-science/user/internal/repository/redis"
	"github.com/marioscordia/rocket-science/user/internal/usecase"
)

type container struct {
	server userv1.UserServiceServer

	useCase handler.UseCase

	repo usecase.Repository

	store repository.Store
	cache repository.Cache

	redisPool *redigo.Pool
	pool      *pgxpool.Pool
}

func NewDiContainer() *container {
	return &container{}
}

func (c *container) GRPCServer(ctx context.Context) userv1.UserServiceServer {
	if c.server == nil {
		c.server = grpc.NewHandler(c.UseCase(ctx))
	}

	return c.server
}

func (c *container) UseCase(ctx context.Context) handler.UseCase {
	if c.useCase == nil {
		c.useCase = usecase.New(c.Repository(ctx))
	}

	return c.useCase
}

func (c *container) Repository(ctx context.Context) usecase.Repository {
	if c.repo == nil {
		c.repo = repository.NewRepository(c.Store(ctx), c.Cache(ctx))
	}

	return c.repo
}

func (c *container) Store(ctx context.Context) repository.Store {
	if c.store == nil {
		pool := c.Pool(ctx)

		c.store = postgres.NewStore(pool)
	}

	return c.store
}

func (c *container) Pool(ctx context.Context) *pgxpool.Pool {
	if c.pool == nil {
		pool, err := pgxpool.New(ctx, config.AppConfig().Postgres.GetURL())
		if err != nil {
			panic(fmt.Errorf("failed to create pgx pool: %w", err))
		}

		closer.AddNamed("Pgx pool", func(_ context.Context) error {
			pool.Close()
			return nil
		})

		c.pool = pool
	}

	return c.pool
}

func (d *container) Cache(ctx context.Context) repository.Cache {
	if d.cache == nil {
		d.cache = redis.NewClient(d.RedisPool(), logger.Logger(), config.AppConfig().Redis.ConnectionTimeout())
	}

	return d.cache
}

func (d *container) RedisPool() *redigo.Pool {
	if d.redisPool == nil {
		d.redisPool = &redigo.Pool{
			MaxIdle:     config.AppConfig().Redis.MaxIdle(),
			IdleTimeout: config.AppConfig().Redis.IdleTimeout(),
			DialContext: func(ctx context.Context) (redigo.Conn, error) {
				return redigo.DialContext(ctx, "tcp", config.AppConfig().Redis.Address())
			},
		}
	}

	return d.redisPool
}
