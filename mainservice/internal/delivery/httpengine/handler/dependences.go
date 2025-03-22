package handler

import "context"

type RoomClient interface {
	GetRooms(ctx context.Context) ([]Room, error)
}
