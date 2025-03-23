package grpc

import (
	"context"
	"mainservice/internal/lib/grpcclient"
)

type BookingClient interface {
	GetAllRooms(ctx context.Context) ([]grpcclient.Room, error)
}
