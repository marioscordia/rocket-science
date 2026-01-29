package converter

import (
	"github.com/marioscordia/rocket-science/inventory/internal/dto"
	"github.com/marioscordia/rocket-science/inventory/internal/repository/model"
)

// ToDTO converts a repository model Part to a dto.Part
func ToDTO(m *model.Part) *dto.Part {
	if m == nil {
		return nil
	}

	return &dto.Part{
		UUID:          m.UUID,
		Name:          m.Name,
		Description:   m.Description,
		Price:         m.Price,
		StockQuantity: m.StockQuantity,
		Category:      categoryFromString(m.Category),
		Dimensions:    convertDimensionsToDTO(m.Dimensions),
		Manufacturer:  convertManufacturerToDTO(m.Manufacturer),
		Tags:          m.Tags,
		Metadata:      convertMetadataToDTO(m.Metadata),
		CreatedAt:     m.CreatedAt,
		UpdatedAt:     m.UpdatedAt,
	}
}

// FromDTO converts a dto.Part to a repository model Part
func FromDTO(d *dto.Part) *model.Part {
	if d == nil {
		return nil
	}

	return &model.Part{
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

// ToDTOList converts a slice of model.Part to a slice of dto.Part
func ToDTOList(parts []*model.Part) []*dto.Part {
	if parts == nil {
		return nil
	}

	result := make([]*dto.Part, len(parts))
	for i, p := range parts {
		result[i] = ToDTO(p)
	}
	return result
}

// FilterToModel converts a dto.PartsFilter to a model.PartsFilter
func FilterToModel(f *dto.PartsFilter) *model.PartsFilter {
	if f == nil {
		return nil
	}

	return &model.PartsFilter{
		UUIDs:                 f.UUIDs,
		Names:                 f.Names,
		Categories:            f.Categories,
		ManufacturerCountries: f.ManufacturerCountries,
		Tags:                  f.Tags,
	}
}

func convertDimensionsToDTO(d *model.Dimensions) dto.Dimensions {
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

func convertDimensionsFromDTO(d dto.Dimensions) *model.Dimensions {
	return &model.Dimensions{
		Length: d.Length,
		Width:  d.Width,
		Height: d.Height,
		Weight: d.Weight,
	}
}

func convertManufacturerToDTO(m *model.Manufacturer) dto.Manufacturer {
	if m == nil {
		return dto.Manufacturer{}
	}
	return dto.Manufacturer{
		Name:    m.Name,
		Country: m.Country,
		Website: m.Website,
	}
}

func convertManufacturerFromDTO(m dto.Manufacturer) *model.Manufacturer {
	return &model.Manufacturer{
		Name:    m.Name,
		Country: m.Country,
		Website: m.Website,
	}
}

func convertMetadataToDTO(m map[string]interface{}) map[string]*dto.Value {
	if m == nil {
		return nil
	}

	result := make(map[string]*dto.Value, len(m))
	for k, v := range m {
		result[k] = convertValueToDTO(v)
	}
	return result
}

func convertValueToDTO(v interface{}) *dto.Value {
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
