package authclient

import (
	"context"
	"fmt"
	"log/slog"

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

func (c Client) GetUser(ctx context.Context, login string, hashPassword string) (string, error) {
	/*client := c.conns.Get().(*grpc.ClientConn)
	defer c.conns.Put(client)

	request := booking.GetAllRoomsRequest{}

	resp, err := booking.NewBookingServiceClient(client).GetAllRooms(context.Background(), &request)
	if err != nil {
		slog.Error(err.Error())
		return nil, err
	}*/
	return "login", nil
}

func (c Client) Register(ctx context.Context, login string, hashPassword string) (string, error) {
	//https://github.com/crutchM/crud/blob/main/internal/repository/postgres/authRepository.go
	/*client := c.conns.Get().(*grpc.ClientConn)
	defer c.conns.Put(client)

	request := booking.GetAllRoomsRequest{}

	resp, err := booking.NewBookingServiceClient(client).GetAllRooms(context.Background(), &request)
	if err != nil {
		slog.Error(err.Error())
		return nil, err
	}*/
	return "login", nil
}
