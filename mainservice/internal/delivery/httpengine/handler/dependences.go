package handler

import (
	"context"
	"mainservice/internal/core/entity"
)

type BookingClient interface {
	GetAllRooms(ctx context.Context, token string) ([]entity.Room, error)
	GetAvailableRooms(ctx context.Context, startDate string, endDate string, token string) ([]entity.Room, error)
	BookRoom(ctx context.Context, bookingRoom entity.BookingRoom, token string) (entity.BookingRoomState, error)
	CancelBooking(ctx context.Context, bookingId int64, token string) (entity.CancelingBookingState, error)
	GetAllBookings(ctx context.Context, token string) ([]entity.Booking, error)
}

type AuthClient interface {
	Register(ctx context.Context, login string, hashPassword string) (string, error)
	Login(ctx context.Context, login string, hashPassword string) (string, error)
}
