package authclient

import (
	"context"
	"fmt"
	"log/slog"
	order "mainservice/product/auth"
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

func (c Client) Login(ctx context.Context, login string, password string) (string, error) {
	client := c.conns.Get().(*grpc.ClientConn)
	defer c.conns.Put(client)

	request := order.LoginRequest{
		Login:    login,
		Password: password,
	}

	resp, err := order.NewAuthServiceClient(client).Login(context.Background(), &request)
	if err != nil {
		slog.Error(err.Error())
		return "", err
	}
	return resp.Token, nil
}

func (c Client) Register(ctx context.Context, login string, password string) (string, error) {
	client := c.conns.Get().(*grpc.ClientConn)
	defer c.conns.Put(client)

	request := order.RegisterRequest{
		Login:    login,
		Password: password,
	}

	resp, err := order.NewAuthServiceClient(client).Register(context.Background(), &request)
	if err != nil {
		slog.Error(err.Error())
		return "", err
	}
	return resp.Token, nil
}

func (c Client) ValidateToken(ctx context.Context, token string) (ValidateToken, error) {
	client := c.conns.Get().(*grpc.ClientConn)
	defer c.conns.Put(client)

	request := order.ValidateTokenRequest{
		Token: token,
	}

	resp, err := order.NewAuthServiceClient(client).ValidateToken(context.Background(), &request)
	if err != nil {
		slog.Error(err.Error())
		return ValidateToken{
			Valid: false,
			Error: "Error from auth service",
		}, err
	}
	return ValidateToken{
		Valid:  resp.Valid,
		Error:  resp.Error,
		Login:  resp.Login,
		UserId: resp.UserId,
	}, nil
}
