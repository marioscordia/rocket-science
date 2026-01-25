package mongo

import (
	"context"
	"errors"
	"fmt"

	"github.com/marioscordia/rocket-science/inventory/internal/dto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type DB struct {
	conn *mongo.Collection
}

func NewMongo(collection *mongo.Collection) *DB {
	return &DB{
		conn: collection,
	}
}

func (db *DB) GetPartByID(ctx context.Context, partID string) (*dto.Part, error) {
	if partID == "" {
		return nil, errors.New("part ID cannot be empty")
	}

	var part Part
	filter := bson.M{"_id": partID}

	err := db.conn.FindOne(ctx, filter).Decode(&part)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, fmt.Errorf("part with ID %s not found", partID)
		}
		return nil, fmt.Errorf("failed to get part: %w", err)
	}

	return part.ToDTO(), nil
}

func (db *DB) ListParts(ctx context.Context, filter *dto.PartsFilter) ([]*dto.Part, error) {
	mongoFilter := buildMongoFilter(filter)

	cursor, err := db.conn.Find(ctx, mongoFilter)
	if err != nil {
		return nil, fmt.Errorf("failed to query parts: %w", err)
	}
	defer cursor.Close(ctx)

	var parts []*Part
	if err := cursor.All(ctx, &parts); err != nil {
		return nil, fmt.Errorf("failed to decode parts: %w", err)
	}

	dtoParts := make([]*dto.Part, len(parts))
	for i, part := range parts {
		dtoParts[i] = part.ToDTO()
	}

	return dtoParts, nil
}

func buildMongoFilter(filter *dto.PartsFilter) bson.M {
	if filter == nil {
		return bson.M{}
	}

	mongoFilter := bson.M{}

	// Filter by UUIDs
	if len(filter.UUIDs) > 0 {
		mongoFilter["_id"] = bson.M{"$in": filter.UUIDs}
	}

	// Filter by names
	if len(filter.Names) > 0 {
		mongoFilter["name"] = bson.M{"$in": filter.Names}
	}

	// Filter by categories
	if len(filter.Categories) > 0 {
		mongoFilter["category"] = bson.M{"$in": filter.Categories}
	}

	// Filter by manufacturer countries
	if len(filter.ManufacturerCountries) > 0 {
		mongoFilter["manufacturer.country"] = bson.M{"$in": filter.ManufacturerCountries}
	}

	// Filter by tags (part must have at least one of the specified tags)
	if len(filter.Tags) > 0 {
		mongoFilter["tags"] = bson.M{"$in": filter.Tags}
	}

	return mongoFilter
}
