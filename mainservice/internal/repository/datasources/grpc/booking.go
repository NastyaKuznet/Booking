package grpc

import (
	"context"
	"fmt"
	"mainservice/internal/lib/booking"
)

type bookingRepository struct {
	client BookingClient
}

func NewBookingRepository(client BookingClient) *bookingRepository {
	return &bookingRepository{client: client}
}

func (b *bookingRepository) GetAllRooms(ctx context.Context) ([]booking.Room, error) {
	rooms, err := b.client.GetAllRooms(ctx)

	if err != nil {
		if len(rooms) == 0 {
			return nil, fmt.Errorf("rooms check warning: %w", ErrRoomsNotFound)
		}
		return nil, fmt.Errorf("%w: %w", ErrInternal, err)
	}
	return rooms, err
}

func (b *bookingRepository) GetAvailableRooms(ctx context.Context, startDate string, endDate string) ([]booking.Room, error) {
	rooms, err := b.client.GetAvailableRooms(ctx, startDate, endDate)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrInternal, err)
	}
	return rooms, err
}

func (b *bookingRepository) BookRoom(ctx context.Context, bookingRoom booking.BookingRoom) (
	booking.BookingRoomState, error) {
	bookingRoomState, err := b.client.BookRoom(ctx, bookingRoom)
	if err != nil {
		return booking.BookingRoomState{
			Success: false,
			Message: "Internal err",
		}, fmt.Errorf("%w: %w", ErrInternal, err)
	}
	return bookingRoomState, err
}

func (b *bookingRepository) CancelBooking(ctx context.Context, bookingId int64) (
	booking.CancelingBookingState, error) {
	cancelingBookingState, err := b.client.CancelBooking(ctx, bookingId)
	if err != nil {
		return booking.CancelingBookingState{
			Success: false,
			Message: "Internal err",
		}, fmt.Errorf("%w: %w", ErrInternal, err)
	}
	return cancelingBookingState, err
}

func (b *bookingRepository) GetAllBookings(ctx context.Context) ([]booking.Booking, error) {
	bookings, err := b.client.GetAllBookings(ctx)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrInternal, err)
	}
	return bookings, err
}
