package dto

import "time"

// Category enum
type Category int32

const (
	CategoryUnknown  Category = 0
	CategoryEngine   Category = 1
	CategoryFuel     Category = 2
	CategoryPorthole Category = 3
	CategoryWing     Category = 4
)

// Part - main entity for inventory items
type Part struct {
	UUID          string
	Name          string
	Description   string
	Price         float64
	StockQuantity int64
	Category      Category
	Dimensions    Dimensions
	Manufacturer  Manufacturer
	Tags          []string
	Metadata      map[string]*Value
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

// Dimensions - physical dimensions of a part
type Dimensions struct {
	Length float64 // in cm
	Width  float64 // in cm
	Height float64 // in cm
	Weight float64 // in kg
}

// Manufacturer - information about part manufacturer
type Manufacturer struct {
	Name    string
	Country string
	Website string
}

// Value - flexible metadata value (oneof in protobuf)
type Value struct {
	// Only one of these should be set
	StringValue *string
	Int64Value  *int64
	DoubleValue *float64
	BoolValue   *bool
}

type PartsFilter struct {
	UUIDs                 []string
	Names                 []string
	Categories            []string
	ManufacturerCountries []string
	Tags                  []string
}
