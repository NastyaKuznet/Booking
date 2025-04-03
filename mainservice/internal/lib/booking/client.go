package booking

import (
	"context"
	"fmt"
	"log/slog"
	"mainservice/product/booking"
	"sync"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	conns *sync.Pool
}

func NewClient(config Config) Client {
	pool := &sync.Pool{
		New: func() interface{} {
			conn, err := grpc.NewClient(
				fmt.Sprintf("%s:%d", config.Host, config.Port), grpc.WithTransportCredentials(insecure.NewCredentials()))
			if err != nil {
				slog.Error(err.Error())
			}
			return conn
		},
	}
	return Client{conns: pool}
}

func (c Client) GetAllRooms(ctx context.Context) ([]Room, error) {
	client := c.conns.Get().(*grpc.ClientConn)
	defer c.conns.Put(client)

	request := booking.GetAllRoomsRequest{}

	resp, err := booking.NewBookingServiceClient(client).GetAllRooms(context.Background(), &request)
	if err != nil {
		slog.Error(err.Error())
		return nil, err
	}
	return fromProtoModelRoom(resp.Rooms), nil
}

func (c Client) GetAvailableRooms(ctx context.Context, startDate string, endDate string) ([]Room, error) {
	client := c.conns.Get().(*grpc.ClientConn)
	defer c.conns.Put(client)

	request := booking.GetAvailableRoomsRequest{
		StartDate: startDate,
		EndDate:   endDate,
	}

	resp, err := booking.NewBookingServiceClient(client).GetAvailableRooms(context.Background(), &request)

	if err != nil {
		slog.Error(err.Error())
		return nil, err
	}
	return fromProtoModelRoom(resp.Rooms), nil
}

func (c Client) BookRoom(ctx context.Context, bookingRoom BookingRoom) (BookingRoomState, error) {
	client := c.conns.Get().(*grpc.ClientConn)
	defer c.conns.Put(client)

	request := booking.BookRoomRequest{
		RoomId:    bookingRoom.RoomId,
		UserId:    bookingRoom.UserId,
		StartDate: bookingRoom.StartDate,
		EndDate:   bookingRoom.EndDate,
	}

	resp, err := booking.NewBookingServiceClient(client).BookRoom(context.Background(), &request)
	if err != nil {
		slog.Error(err.Error())
		return BookingRoomState{
			Success: false,
			Message: "Fail in BookRoom send",
		}, err
	}

	return BookingRoomState{
		Success: resp.Success,
		Message: resp.Message,
	}, nil
}

func (c Client) CancelBooking(ctx context.Context, bookingId int64) (CancelingBookingState, error) {
	client := c.conns.Get().(*grpc.ClientConn)
	defer c.conns.Put(client)

	request := booking.CancelBookingRequest{
		BookingId: bookingId,
	}

	resp, err := booking.NewBookingServiceClient(client).CancelBooking(context.Background(), &request)
	if err != nil {
		slog.Error(err.Error())
		return CancelingBookingState{
			Success: false,
		}, err
	}

	return CancelingBookingState{
		Success: resp.Success,
		Message: resp.Message,
	}, nil
}

func (c Client) GetAllBookings(ctx context.Context) ([]Booking, error) {
	client := c.conns.Get().(*grpc.ClientConn)
	defer c.conns.Put(client)

	request := booking.GetAllBookingsRequest{}

	resp, err := booking.NewBookingServiceClient(client).GetAllBookings(context.Background(), &request)
	if err != nil {
		slog.Error(err.Error())
		return nil, err
	}

	return fromProtoModelBooking(resp.Bookings), nil
}
