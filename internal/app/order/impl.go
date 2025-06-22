package order

import (
	"github.com/peterprospl12/logi-flow/internal/domain/order"
	"github.com/peterprospl12/logi-flow/internal/infra/queue"
)

type svc struct {
	repo order.Repository
	mq   queue.Publisher
}

func NewService(r order.Repository, mq queue.Publisher) Service {
	return &svc{repo: r, mq: mq}
}

func (s *svc) AcceptBid(cmd AcceptBidDTO) error {
	o, err := s.repo.GetByID(cmd.OrderID)
	if err != nil {
		return err
	}

}

func (s *svc) CreateOrder(cmd CreateOrderDTO) error {
	o, err := order.New(cmd.ShipperID, cmd.From, cmd.To, cmd.Price)
	if err != nil {
		return err
	}

	return nil
}
