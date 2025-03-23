package usecase

import (
	"context"
	"log/slog"
	"mainservice/internal/core/entity"
	"mainservice/internal/lib/grpcclient"
)

type Usecase struct {
	bookingClient      BookingClient
	notificationClient NotificationClient
}

func NewUsecase(bookingClient BookingClient, notificationClient NotificationClient) Usecase {
	return Usecase{
		bookingClient:      bookingClient,
		notificationClient: notificationClient,
	}
}

func (uc Usecase) GetAllRooms(ctx context.Context) ([]entity.Room, error) {
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

	roomPointers := make([]*grpcclient.Room, len(rooms))
	for i := range rooms {
		roomPointers[i] = &rooms[i] // Берем адрес каждого элемента
	}

	return entity.FromProtoModel(roomPointers), nil
}
