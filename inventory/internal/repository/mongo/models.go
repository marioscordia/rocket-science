package mongo

import (
	"time"
)

type Part struct {
	UUID          string                 `bson:"_id"`
	Name          string                 `bson:"name"`
	Description   string                 `bson:"description"`
	Price         float64                `bson:"price"`
	StockQuantity int64                  `bson:"stock_quantity"`
	Category      string                 `bson:"category"`
	Dimensions    *Dimensions            `bson:"dimensions,omitempty"`
	Manufacturer  *Manufacturer          `bson:"manufacturer,omitempty"`
	Tags          []string               `bson:"tags,omitempty"`
	Metadata      map[string]interface{} `bson:"metadata,omitempty"`
	CreatedAt     time.Time              `bson:"created_at"`
	UpdatedAt     time.Time              `bson:"updated_at"`
}

type Dimensions struct {
	Length float64 `bson:"length"`
	Width  float64 `bson:"width"`
	Height float64 `bson:"height"`
	Weight float64 `bson:"weight"`
}

type Manufacturer struct {
	Name    string `bson:"name"`
	Country string `bson:"country"`
	Website string `bson:"website,omitempty"`
}

type PartsFilter struct {
	UUIDs                 []string
	Names                 []string
	Categories            []string
	ManufacturerCountries []string
	Tags                  []string
}
