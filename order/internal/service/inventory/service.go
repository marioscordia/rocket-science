package inventory

import (
	"context"

	"github.com/marioscordia/rocket-science/order/internal/dto"
	inventoryv1 "github.com/marioscordia/rocket-science/shared/pkg/proto/inventory/v1"
	"google.golang.org/grpc"
)

type Service struct {
	client inventoryv1.InventoryServiceClient
}

func NewService(conn grpc.ClientConnInterface) *Service {
	client := inventoryv1.NewInventoryServiceClient(conn)

	return &Service{
		client: client,
	}
}

func (s *Service) GetParts(ctx context.Context, partIDs []string) (map[string]*dto.Part, error) {
	req := &inventoryv1.ListPartsRequest{
		Filter: &inventoryv1.PartsFilter{
			Uuids: partIDs,
		},
	}

	resp, err := s.client.ListParts(ctx, req)
	if err != nil {
		return nil, err
	}

	parts := make(map[string]*dto.Part, len(resp.Parts))

	for _, part := range resp.Parts {
		parts[part.Uuid] = &dto.Part{
			ID:       part.Uuid,
			Price:    part.Price,
			Quantity: int32(part.StockQuantity),
		}
	}

	return parts, nil
}
