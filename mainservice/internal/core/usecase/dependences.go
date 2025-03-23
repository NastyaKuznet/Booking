package usecase

import (
	"context"
	"mainservice/internal/lib/grpcclient"
)

type BookingClient interface {
	GetAllRooms(ctx context.Context) ([]grpcclient.Room, error)
}

type NotificationClient interface {
	Publish(ctx context.Context, message string) error
}
