package model

import "time"

// Part represents the database model for parts
type Part struct {
	UUID          string
	Name          string
	Description   string
	Price         float64
	StockQuantity int64
	Category      string
	Dimensions    *Dimensions
	Manufacturer  *Manufacturer
	Tags          []string
	Metadata      map[string]interface{}
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

// Dimensions represents physical dimensions of a part
type Dimensions struct {
	Length float64
	Width  float64
	Height float64
	Weight float64
}

// Manufacturer represents manufacturer information
type Manufacturer struct {
	Name    string
	Country string
	Website string
}

// PartsFilter represents filter criteria for querying parts
type PartsFilter struct {
	UUIDs                 []string
	Names                 []string
	Categories            []string
	ManufacturerCountries []string
	Tags                  []string
}
