package grpc

import (
	"github.com/marioscordia/rocket-science/inventory/internal/dto"
	inventoryv1 "github.com/marioscordia/rocket-science/shared/pkg/proto/inventory/v1"
)

// ToProtoPart converts a dto.Part to inventoryv1.Part
func ToProtoPart(d *dto.Part) *inventoryv1.Part {
	if d == nil {
		return nil
	}

	return &inventoryv1.Part{
		Uuid:          d.UUID,
		Name:          d.Name,
		Description:   d.Description,
		Price:         d.Price,
		StockQuantity: d.StockQuantity,
		Category:      categoryToProto(d.Category),
		Dimensions:    dimensionsToProto(d.Dimensions),
		Manufacturer:  manufacturerToProto(d.Manufacturer),
		Tags:          d.Tags,
		Metadata:      metadataToProto(d.Metadata),
		CreatedAt:     nil, // Convert time.Time to timestamppb.Timestamp if needed
		UpdatedAt:     nil, // Convert time.Time to timestamppb.Timestamp if needed
	}
}

// ToProtoParts converts a slice of dto.Part to a slice of inventoryv1.Part
func ToProtoParts(parts []*dto.Part) []*inventoryv1.Part {
	if parts == nil {
		return nil
	}

	result := make([]*inventoryv1.Part, len(parts))
	for i, part := range parts {
		result[i] = ToProtoPart(part)
	}
	return result
}

func categoryToProto(c dto.Category) inventoryv1.Category {
	return inventoryv1.Category(c)
}

func dimensionsToProto(d dto.Dimensions) *inventoryv1.Dimensions {
	return &inventoryv1.Dimensions{
		Length: d.Length,
		Width:  d.Width,
		Height: d.Height,
		Weight: d.Weight,
	}
}

func manufacturerToProto(m dto.Manufacturer) *inventoryv1.Manufacturer {
	return &inventoryv1.Manufacturer{
		Name:    m.Name,
		Country: m.Country,
		Website: m.Website,
	}
}

func metadataToProto(m map[string]*dto.Value) map[string]*inventoryv1.Value {
	if m == nil {
		return nil
	}

	result := make(map[string]*inventoryv1.Value, len(m))
	for k, v := range m {
		result[k] = valueToProto(v)
	}
	return result
}

func valueToProto(v *dto.Value) *inventoryv1.Value {
	if v == nil {
		return &inventoryv1.Value{}
	}

	if v.StringValue != nil {
		return &inventoryv1.Value{
			Value: &inventoryv1.Value_StringValue{StringValue: *v.StringValue},
		}
	}
	if v.Int64Value != nil {
		return &inventoryv1.Value{
			Value: &inventoryv1.Value_Int64Value{Int64Value: *v.Int64Value},
		}
	}
	if v.DoubleValue != nil {
		return &inventoryv1.Value{
			Value: &inventoryv1.Value_DoubleValue{DoubleValue: *v.DoubleValue},
		}
	}
	if v.BoolValue != nil {
		return &inventoryv1.Value{
			Value: &inventoryv1.Value_BoolValue{BoolValue: *v.BoolValue},
		}
	}
	return &inventoryv1.Value{}
}
