package grpc

import (
	"context"

	"github.com/marioscordia/rocket-science/inventory/internal/dto"
	"github.com/marioscordia/rocket-science/inventory/internal/handler"
	inventoryv1 "github.com/marioscordia/rocket-science/shared/pkg/proto/inventory/v1"
)

type Handler struct {
	inventoryv1.UnimplementedInventoryServiceServer
	usecase handler.UseCase
}

func (h *Handler) GetPart(ctx context.Context, req *inventoryv1.GetPartRequest) (*inventoryv1.GetPartResponse, error) {
	part, err := h.usecase.GetPartByID(ctx, req.GetUuid())
	if err != nil {
		// Handle error appropriately (e.g., return gRPC error codes)
		return nil, err
	}

	return &inventoryv1.GetPartResponse{
		Part: ToProtoPart(part),
	}, nil
}

func (h *Handler) ListParts(ctx context.Context, req *inventoryv1.ListPartsRequest) (*inventoryv1.ListPartsResponse, error) {
	categories := make([]string, len(req.Filter.GetCategories()))
	for i, category := range req.Filter.GetCategories() {
		categories[i] = category.String()
	}

	filters := &dto.PartsFilter{
		UUIDs:                 req.Filter.GetUuids(),
		Names:                 req.Filter.GetNames(),
		Categories:            categories,
		ManufacturerCountries: req.Filter.GetManufacturerCountries(),
		Tags:                  req.Filter.GetTags(),
	}

	parts, err := h.usecase.ListParts(ctx, filters)
	if err != nil {
		return nil, err
	}

	return &inventoryv1.ListPartsResponse{
		Parts: ToProtoParts(parts),
	}, nil
}
