package grpc

import (
	"context"
	"fmt"
	"mainservice/internal/lib/grpcclient"
)

type bookingRepository struct {
	client BookingClient
}

func NewBookingRepository(client BookingClient) *bookingRepository {
	return &bookingRepository{client: client}
}

func (b *bookingRepository) GetAllRooms(ctx context.Context) ([]grpcclient.Room, error) {
	rooms, err := b.client.GetAllRooms(ctx)

	if err != nil {
		if len(rooms) == 0 {
			return nil, fmt.Errorf("rooms check warning: %w", ErrRoomsNotFound)
		}
		return nil, fmt.Errorf("%w: %w", ErrInternal, err)
	}
	return rooms, err
}
