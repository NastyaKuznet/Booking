package rabbitclient

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"sync"

	"github.com/rabbitmq/amqp091-go"
)

type RabbitClient struct {
	pool    *sync.Pool
	channel string
}

func NewRabbirClient(config Config) RabbitClient {
	pool := sync.Pool{
		New: func() interface{} {
			conn, err := amqp091.Dial(
				fmt.Sprintf("amqp://%s:%s@%s:%s",
					config.Login,
					config.Password,
					config.Host,
					config.Port),
			)
			if err != nil {
				slog.Error(err.Error())
			}

			return conn
		},
	}

	return RabbitClient{pool: &pool, channel: config.Channel}
}

func (r RabbitClient) Publish(ctx context.Context, message string) error {
	conn := r.pool.Get().(*amqp091.Connection)
	defer r.pool.Put(conn)

	ch, err := conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	res, err := json.Marshal(message)
	if err != nil {
		return err
	}

	err = ch.Publish(
		"",
		r.channel,
		false,
		false,
		amqp091.Publishing{ContentType: "application/json", Body: res})

	return err
}
