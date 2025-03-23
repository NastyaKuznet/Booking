package handler

import "context"

type BookingClient interface {
	GetRooms(ctx context.Context) ([]Room, error)
}
