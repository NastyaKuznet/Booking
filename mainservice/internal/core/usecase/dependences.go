package usecase

import (
	"context"
	"mainservice/internal/lib/booking"
)

type BookingClient interface {
	GetAllRooms(ctx context.Context) ([]booking.Room, error)
	GetAvailableRooms(ctx context.Context) ([]booking.Room, error)
	BookRoom(ctx context.Context, bookingRoom booking.BookingRoom) (booking.BookingRoomState, error)
	CancelBooking(ctx context.Context, bookingId int64) (booking.CancelingBookingState, error)
}

type NotificationClient interface {
	Publish(ctx context.Context, message string) error
}

type AuthClient interface {
	Login(ctx context.Context, login string, hashPassword string) (string, error)
	Register(ctx context.Context, login string, hashPassword string) (string, error)
}
