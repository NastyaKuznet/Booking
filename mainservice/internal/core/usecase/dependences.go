package usecase

import (
	"context"
	"mainservice/internal/lib/authclient"
	"mainservice/internal/lib/booking"
)

type BookingClient interface {
	GetAllRooms(ctx context.Context) ([]booking.Room, error)
	GetAvailableRooms(ctx context.Context, startDate string, endDate string) ([]booking.Room, error)
	BookRoom(ctx context.Context, bookingRoom booking.BookingRoom) (booking.BookingRoomState, error)
	CancelBooking(ctx context.Context, bookingId int64) (booking.CancelingBookingState, error)
	GetAllBookings(ctx context.Context) ([]booking.Booking, error)
}

type NotificationClient interface {
	Publish(ctx context.Context, message string) error
}

type AuthClient interface {
	Login(ctx context.Context, login string, hashPassword string) (string, error)
	Register(ctx context.Context, login string, hashPassword string) (string, error)
	ValidateToken(ctx context.Context, token string) (authclient.ValidateToken, error)
}
