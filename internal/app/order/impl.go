package order

import (
	"github.com/peterprospl12/logi-flow/internal/domain/order"
)

type svc struct {
	repo order.Repository
	mq   queue.Publisher
}

func NewService(r order.Repository, mq queue.Publisher) Service {
	return &svc{repo: r, mq: mq}
}

func (s *svc) AcceptBid(cmd AcceptBidDTO) error {
	o, err := s.repo.Get(cmd.OrderID)
	if err != nil {
		return err
	}
}
