package queue

import "context"

type Publisher interface {
	Publish(ctx context.Context, topic string, message []byte) error
	Close() error
}
