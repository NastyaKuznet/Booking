package handler

import (
	"context"
	"mainservice/internal/core/entity"
)

type BookingClient interface {
	GetAllRooms(ctx context.Context) ([]entity.Room, error)
	GetAvailableRooms(ctx context.Context) ([]entity.Room, error)
	BookRoom(ctx context.Context, bookingRoom entity.BookingRoom) (entity.BookingRoomState, error)
	CancelBooking(ctx context.Context, bookingId int64) (entity.CancelingBookingState, error)
}

type AuthClient interface {
	Register(ctx context.Context, login string, hashPassword string) (string, error)
	Login(ctx context.Context, login string, hashPassword string) (string, error)
}
