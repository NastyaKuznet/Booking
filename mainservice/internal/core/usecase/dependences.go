package usecase

import (
	"context"
	"mainservice/internal/lib/grpcclient"
)

type BookingClient interface {
	GetAllRooms(ctx context.Context) ([]grpcclient.Room, error)
	GetAvailableRooms(ctx context.Context) ([]grpcclient.Room, error)
	BookRoom(ctx context.Context, bookingRoom grpcclient.BookingRoom) (grpcclient.BookingRoomState, error)
	CancelBooking(ctx context.Context, bookingId int64) (grpcclient.CancelingBookingState, error)
}

type NotificationClient interface {
	Publish(ctx context.Context, message string) error
}

type AuthClient interface {
	GetUser(ctx context.Context, login string, hashPassword string) (string, error)
	Register(ctx context.Context, login string, hashPassword string) (string, error)
}
