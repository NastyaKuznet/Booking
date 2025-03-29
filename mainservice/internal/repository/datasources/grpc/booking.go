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

func (b *bookingRepository) GetAvailableRooms(ctx context.Context) ([]grpcclient.Room, error) {
	rooms, err := b.client.GetAvailableRooms(ctx)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrInternal, err)
	}
	return rooms, err
}

func (b *bookingRepository) BookRoom(ctx context.Context, bookingRoom grpcclient.BookingRoom) (
	grpcclient.BookingRoomState, error) {
	bookingRoomState, err := b.client.BookRoom(ctx, bookingRoom)
	if err != nil {
		return grpcclient.BookingRoomState{
			Success: false,
			Message: "Internal err",
		}, fmt.Errorf("%w: %w", ErrInternal, err)
	}
	return bookingRoomState, err
}

func (b *bookingRepository) CancelBooking(ctx context.Context, bookingId int64) (
	grpcclient.CancelingBookingState, error) {
	cancelingBookingState, err := b.client.CancelBooking(ctx, bookingId)
	if err != nil {
		return grpcclient.CancelingBookingState{
			Success: false,
			Message: "Internal err",
		}, fmt.Errorf("%w: %w", ErrInternal, err)
	}
	return cancelingBookingState, err
}
