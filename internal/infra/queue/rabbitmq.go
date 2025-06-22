package queue

import (
	"context"
	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQPublisher struct {
	conn    *amqp.Connection
	channel *amqp.Channel
}

func NewRabbitMQPublisher(amqpURL string) (*RabbitMQPublisher, error) {
	conn, err := amqp.Dial(amqpURL)
	if err != nil {
		return nil, err
	}

	channel, err := conn.Channel()
	if err != nil {
		err := conn.Close()
		if err != nil {
			return nil, err
		}
		return nil, err
	}

	return &RabbitMQPublisher{
		conn:    conn,
		channel: channel,
	}, nil
}

func (p *RabbitMQPublisher) Publish(ctx context.Context, topic string, message []byte) error {
	q, err := p.channel.QueueDeclare(
		topic, // name
		true,  // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		return err
	}

	return p.channel.PublishWithContext(
		ctx,
		"", // exchange
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType:  "application/json",
			DeliveryMode: amqp.Persistent,
			Body:         message,
		})
}
