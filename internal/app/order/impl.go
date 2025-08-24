package order

import (
	"context"
	"github.com/google/uuid"
	domain "github.com/peterprospl12/logi-flow/internal/domain/order"
	//"github.com/peterprospl12/logi-flow/internal/infra/queue"
)

type service struct {
	repo Repository
	//publ queue.Publisher
}

func NewService(r Repository) Service {
	return &service{repo: r}
}

func (s *service) CreateOrder(ctx context.Context, cmd CreateOrderCommand) (*domain.Order, error) {
	o := &domain.Order{
		ID:           uuid.New(),
		ShipperID:    cmd.ShipperID,
		Origin:       cmd.ShipperID,
		Destination:  cmd.Destination,
		InitialPrice: cmd.InitialPrice,
		Status:       domain.Created,
	}

	if err := s.repo.Save(ctx, o); err != nil {
		return nil, err
	}

	return o, nil
}

func (s *service) GetOrders(ctx context.Context) ([]*domain.Order, error) {
	return s.repo.List(ctx)
}
