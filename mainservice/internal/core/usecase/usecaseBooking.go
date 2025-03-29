package usecase

import (
	"context"
	"fmt"
	"log/slog"
	"mainservice/internal/core/entity"
	"mainservice/internal/lib/booking"
)

type UsecaseBooking struct {
	bookingClient      BookingClient
	notificationClient NotificationClient
}

func NewUsecaseBooking(bookingClient BookingClient, notificationClient NotificationClient) UsecaseBooking {
	return UsecaseBooking{
		bookingClient:      bookingClient,
		notificationClient: notificationClient,
	}
}

func (uc UsecaseBooking) GetAllRooms(ctx context.Context) ([]entity.Room, error) {
	rooms, err := uc.bookingClient.GetAllRooms(ctx)
	if err != nil {
		slog.Error("get rooms error", "err", err)
		return nil, err
	}
	err = uc.notificationClient.Publish(ctx, "All rooms have been requested")

	if err != nil {
		slog.Error("send notification error", "err", err)

		return nil, err
	}

	roomPointers := make([]*booking.Room, len(rooms))
	for i := range rooms {
		roomPointers[i] = &rooms[i]
	}

	return entity.FromProtoModelRooms(roomPointers), nil
}

func (uc UsecaseBooking) GetAvailableRooms(ctx context.Context) ([]entity.Room, error) {
	rooms, err := uc.bookingClient.GetAvailableRooms(ctx)
	if err != nil {
		slog.Error("get available rooms error", "err", err)
		return nil, err
	}

	err = uc.notificationClient.Publish(ctx, "All available rooms have been requested")

	if err != nil {
		slog.Error("send notification error", "err", err)

		return nil, err
	}

	roomPointers := make([]*booking.Room, len(rooms))
	for i := range rooms {
		roomPointers[i] = &rooms[i]
	}

	return entity.FromProtoModelRooms(roomPointers), nil
}

func (uc UsecaseBooking) BookRoom(ctx context.Context, bookingRoom entity.BookingRoom) (entity.BookingRoomState, error) {
	bookingRoomState, err := uc.bookingClient.BookRoom(ctx, booking.BookingRoom{
		RoomId:    bookingRoom.RoomId,
		UserId:    bookingRoom.UserId,
		StartDate: bookingRoom.StartDate,
		EndDate:   bookingRoom.EndDate,
	})
	if err != nil {
		slog.Error("book room error", "err", err)
		return entity.BookingRoomState{
			Success: false,
			Message: "book room error",
		}, err
	}

	err = uc.notificationClient.Publish(ctx, fmt.Sprintf(
		"Room %d have been requested by user %d", bookingRoom.RoomId, bookingRoom.UserId))

	if err != nil {
		slog.Error("send notification error", "err", err)

		return entity.BookingRoomState{
			Success: false,
			Message: "book room error",
		}, err
	}

	return entity.BookingRoomState{
		Success: bookingRoomState.Success,
		Message: bookingRoomState.Message,
	}, nil
}

func (uc UsecaseBooking) CancelBooking(ctx context.Context, bookingId int64) (entity.CancelingBookingState, error) {
	cancelingBookingState, err := uc.bookingClient.CancelBooking(ctx, bookingId)
	if err != nil {
		slog.Error("book room error", "err", err)
		return entity.CancelingBookingState{
			Success: false,
		}, err
	}

	err = uc.notificationClient.Publish(ctx, fmt.Sprintf(
		"Canceling booking %d have been requested", bookingId))

	if err != nil {
		slog.Error("send notification error", "err", err)

		return entity.CancelingBookingState{
			Success: false,
		}, err
	}

	return entity.CancelingBookingState{
		Success: cancelingBookingState.Success,
		Message: cancelingBookingState.Message,
	}, nil
}
