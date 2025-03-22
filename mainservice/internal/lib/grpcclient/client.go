package grpcclient

import (
	"context"
	"fmt"
	"log/slog"
	"mainservice/product/booking"
	"sync"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	conns *sync.Pool
}

func NewClient(config Config) Client {
	pool := &sync.Pool{
		New: func() interface{} {
			conn, err := grpc.NewClient(
				fmt.Sprintf("%s:%d", config.Host, config.Port), grpc.WithTransportCredentials(insecure.NewCredentials()))
			if err != nil {
				slog.Error(err.Error())
			}
			return conn
		},
	}
	return Client{conns: pool}
}

func (c Client) GetAllRooms(ctx context.Context) ([]Room, error) {
	client := c.conns.Get().(*grpc.ClientConn)
	defer c.conns.Put(client)

	request := booking.GetAllRoomsRequest{}

	resp, err := booking.NewBookingServiceClient(client).GetAllRooms(context.Background(), &request)
	if err != nil {
		slog.Error(err.Error())
		return nil, err
	}
	return fromProtoModel(resp.Rooms), nil
}
