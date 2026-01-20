package mongo

import "github.com/marioscordia/rocket-science/inventory/internal/dto"

// ToDTO converts a mongo Part to a dto.Part
func (p *Part) ToDTO() *dto.Part {
	if p == nil {
		return nil
	}

	return &dto.Part{
		UUID:          p.UUID,
		Name:          p.Name,
		Description:   p.Description,
		Price:         p.Price,
		StockQuantity: p.StockQuantity,
		Category:      categoryFromString(p.Category),
		Dimensions:    convertDimensions(p.Dimensions),
		Manufacturer:  convertManufacturer(p.Manufacturer),
		Tags:          p.Tags,
		Metadata:      convertMetadata(p.Metadata),
		CreatedAt:     p.CreatedAt,
		UpdatedAt:     p.UpdatedAt,
	}
}

// FromDTO converts a dto.Part to a mongo Part
func FromDTO(d *dto.Part) *Part {
	if d == nil {
		return nil
	}

	return &Part{
		UUID:          d.UUID,
		Name:          d.Name,
		Description:   d.Description,
		Price:         d.Price,
		StockQuantity: d.StockQuantity,
		Category:      categoryToString(d.Category),
		Dimensions:    convertDimensionsFromDTO(d.Dimensions),
		Manufacturer:  convertManufacturerFromDTO(d.Manufacturer),
		Tags:          d.Tags,
		Metadata:      convertMetadataFromDTO(d.Metadata),
		CreatedAt:     d.CreatedAt,
		UpdatedAt:     d.UpdatedAt,
	}
}

func convertDimensions(d *Dimensions) dto.Dimensions {
	if d == nil {
		return dto.Dimensions{}
	}
	return dto.Dimensions{
		Length: d.Length,
		Width:  d.Width,
		Height: d.Height,
		Weight: d.Weight,
	}
}

func convertDimensionsFromDTO(d dto.Dimensions) *Dimensions {
	return &Dimensions{
		Length: d.Length,
		Width:  d.Width,
		Height: d.Height,
		Weight: d.Weight,
	}
}

func convertManufacturer(m *Manufacturer) dto.Manufacturer {
	if m == nil {
		return dto.Manufacturer{}
	}
	return dto.Manufacturer{
		Name:    m.Name,
		Country: m.Country,
		Website: m.Website,
	}
}

func convertManufacturerFromDTO(m dto.Manufacturer) *Manufacturer {
	return &Manufacturer{
		Name:    m.Name,
		Country: m.Country,
		Website: m.Website,
	}
}

func convertMetadata(m map[string]interface{}) map[string]*dto.Value {
	if m == nil {
		return nil
	}

	result := make(map[string]*dto.Value, len(m))
	for k, v := range m {
		result[k] = convertValue(v)
	}
	return result
}

func convertValue(v interface{}) *dto.Value {
	if v == nil {
		return nil
	}

	value := &dto.Value{}
	switch val := v.(type) {
	case string:
		value.StringValue = &val
	case int:
		i64 := int64(val)
		value.Int64Value = &i64
	case int32:
		i64 := int64(val)
		value.Int64Value = &i64
	case int64:
		value.Int64Value = &val
	case float32:
		f64 := float64(val)
		value.DoubleValue = &f64
	case float64:
		value.DoubleValue = &val
	case bool:
		value.BoolValue = &val
	}
	return value
}

func convertMetadataFromDTO(m map[string]*dto.Value) map[string]interface{} {
	if m == nil {
		return nil
	}

	result := make(map[string]interface{}, len(m))
	for k, v := range m {
		if v == nil {
			continue
		}
		if v.StringValue != nil {
			result[k] = *v.StringValue
		} else if v.Int64Value != nil {
			result[k] = *v.Int64Value
		} else if v.DoubleValue != nil {
			result[k] = *v.DoubleValue
		} else if v.BoolValue != nil {
			result[k] = *v.BoolValue
		}
	}
	return result
}

func categoryFromString(s string) dto.Category {
	switch s {
	case "engine":
		return dto.CategoryEngine
	case "fuel":
		return dto.CategoryFuel
	case "porthole":
		return dto.CategoryPorthole
	case "wing":
		return dto.CategoryWing
	default:
		return dto.CategoryUnknown
	}
}

func categoryToString(c dto.Category) string {
	switch c {
	case dto.CategoryEngine:
		return "engine"
	case dto.CategoryFuel:
		return "fuel"
	case dto.CategoryPorthole:
		return "porthole"
	case dto.CategoryWing:
		return "wing"
	default:
		return "unknown"
	}
}
