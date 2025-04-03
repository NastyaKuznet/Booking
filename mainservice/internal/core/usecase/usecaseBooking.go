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
	authClient         AuthClient
}

func NewUsecaseBooking(bookingClient BookingClient, notificationClient NotificationClient, authClient AuthClient) UsecaseBooking {
	return UsecaseBooking{
		bookingClient:      bookingClient,
		notificationClient: notificationClient,
		authClient:         authClient,
	}
}

func (uc UsecaseBooking) GetAllRooms(ctx context.Context, token string) ([]entity.Room, error) {
	validateToken, err := uc.authClient.ValidateToken(ctx, token)
	if err != nil || !validateToken.Valid {
		slog.Error("Valid token error", "err", err)
		return nil, err
	}

	rooms, err := uc.bookingClient.GetAllRooms(ctx)
	if err != nil {
		slog.Error("get rooms error", "err", err)
		return nil, err
	}
	err = uc.notificationClient.Publish(ctx, fmt.Sprintf("All rooms have been requested by user %d %s", validateToken.UserId, validateToken.Login))

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

func (uc UsecaseBooking) GetAvailableRooms(ctx context.Context, startDate string, endDate string, token string) ([]entity.Room, error) {
	validateToken, err := uc.authClient.ValidateToken(ctx, token)
	if err != nil {
		slog.Error("Valid token error", "err", err)
		return nil, err
	}

	rooms, err := uc.bookingClient.GetAvailableRooms(ctx, startDate, endDate)
	if err != nil {
		slog.Error("get available rooms error", "err", err)
		return nil, err
	}

	err = uc.notificationClient.Publish(ctx, fmt.Sprintf("All available rooms have been requested by user %d %s", validateToken.UserId, validateToken.Login))

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

func (uc UsecaseBooking) BookRoom(ctx context.Context, bookingRoom entity.BookingRoom, token string) (entity.BookingRoomState, error) {
	validateToken, err := uc.authClient.ValidateToken(ctx, token)
	if err != nil {
		slog.Error("Valid token error", "err", err)
		return entity.BookingRoomState{
			Success: false,
			Message: "Valid token error",
		}, err
	}

	bookingRoomState, err := uc.bookingClient.BookRoom(ctx, booking.BookingRoom{
		RoomId:    bookingRoom.RoomId,
		UserId:    validateToken.UserId,
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
		"Room %d have been requested by user %d %s", bookingRoom.RoomId, validateToken.UserId, validateToken.Login))

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

func (uc UsecaseBooking) CancelBooking(ctx context.Context, bookingId int64, token string) (entity.CancelingBookingState, error) {
	validateToken, err := uc.authClient.ValidateToken(ctx, token)
	if err != nil {
		slog.Error("Valid token error", "err", err)
		return entity.CancelingBookingState{
			Success: false,
			Message: "Valid token error",
		}, err
	}

	cancelingBookingState, err := uc.bookingClient.CancelBooking(ctx, bookingId)
	if err != nil {
		slog.Error("book room error", "err", err)
		return entity.CancelingBookingState{
			Success: false,
			Message: "book room error",
		}, err
	}

	err = uc.notificationClient.Publish(ctx, fmt.Sprintf(
		"Canceling booking %d have been requested by user %d %s", bookingId, validateToken.UserId, validateToken.Login))

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

func (uc UsecaseBooking) GetAllBookings(ctx context.Context, token string) ([]entity.Booking, error) {
	validateToken, err := uc.authClient.ValidateToken(ctx, token)
	if err != nil {
		slog.Error("Valid token error", "err", err)
		return nil, err
	}

	bookings, err := uc.bookingClient.GetAllBookings(ctx)
	if err != nil {
		slog.Error("book room error", "err", err)
		return nil, err
	}

	err = uc.notificationClient.Publish(ctx, fmt.Sprintf(
		"Get all bookings by user %d %s", validateToken.UserId, validateToken.Login))

	if err != nil {
		slog.Error("send notification error", "err", err)

		return nil, err
	}

	bookingsPointers := make([]*booking.Booking, len(bookings))
	for i := range bookings {
		bookingsPointers[i] = &bookings[i]
	}

	return entity.FromProtoModelBooking(bookingsPointers), nil
}
