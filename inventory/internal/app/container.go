package app

import (
	"context"
	"fmt"

	"github.com/marioscordia/rocket-science/inventory/internal/config"
	"github.com/marioscordia/rocket-science/inventory/internal/handler"
	"github.com/marioscordia/rocket-science/inventory/internal/handler/grpc"
	"github.com/marioscordia/rocket-science/inventory/internal/repository"
	"github.com/marioscordia/rocket-science/inventory/internal/usecase"
	"github.com/marioscordia/rocket-science/platform/pkg/closer"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	mongoStore "github.com/marioscordia/rocket-science/inventory/internal/repository/mongo"
)

type container struct {
	handler *grpc.Handler
	usecase handler.UseCase
	repo    usecase.Repo

	mongoStore        repository.Store
	mongoDBClient     *mongo.Client
	mongoDBHandle     *mongo.Database
	mongoDBCollection *mongo.Collection
}

func NewDiContainer() *container {
	return &container{}
}

func (c *container) Handler(ctx context.Context) *grpc.Handler {
	if c.handler == nil {
		c.handler = grpc.NewHandler(c.Usecase(ctx))
	}

	return c.handler
}

func (c *container) Usecase(ctx context.Context) handler.UseCase {
	if c.usecase == nil {
		c.usecase = usecase.NewUseCase(c.Repository(ctx))
	}

	return c.usecase
}

func (c *container) Repository(ctx context.Context) usecase.Repo {
	if c.repo == nil {
		c.repo = repository.NewRepository(c.MongoStore(ctx))
	}

	return c.repo
}

func (d *container) MongoStore(ctx context.Context) repository.Store {
	if d.mongoStore == nil {
		d.mongoStore = mongoStore.NewMongo(d.MongoDBCollection(ctx))
	}

	return d.mongoStore
}

func (d *container) MongoDBClient(ctx context.Context) *mongo.Client {
	if d.mongoDBClient == nil {
		client, err := mongo.Connect(ctx, options.Client().ApplyURI(config.AppConfig().Mongo.GetURI()))
		if err != nil {
			panic(fmt.Sprintf("failed to connect to MongoDB: %s\n", err.Error()))
		}

		err = client.Ping(ctx, readpref.Primary())
		if err != nil {
			panic(fmt.Sprintf("failed to ping MongoDB: %v\n", err))
		}

		closer.AddNamed("MongoDB client", func(ctx context.Context) error {
			return client.Disconnect(ctx)
		})

		d.mongoDBClient = client
	}

	return d.mongoDBClient
}

func (d *container) MongoDBHandle(ctx context.Context) *mongo.Database {
	if d.mongoDBHandle == nil {
		d.mongoDBHandle = d.MongoDBClient(ctx).Database(config.AppConfig().Mongo.GetDatabase())
	}

	return d.mongoDBHandle
}

func (d *container) MongoDBCollection(ctx context.Context) *mongo.Collection {
	if d.mongoDBCollection == nil {
		d.mongoDBCollection = d.MongoDBHandle(ctx).Collection(config.AppConfig().Mongo.GetCollection())
	}

	return d.mongoDBCollection
}
