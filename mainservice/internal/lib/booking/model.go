package booking

import "mainservice/product/booking"

type Room struct {
	Id          int64
	RoomNumber  string
	Description string
	Price       int64
}

type BookingRoom struct {
	RoomId    int64
	UserId    int64
	StartDate string
	EndDate   string
}

type Booking struct {
	Id        int64
	RoomId    int64
	UserId    int64
	StartDate string
	EndDate   string
}

type BookingRoomState struct {
	Success bool
	Message string
}

type CancelingBookingState struct {
	Success bool
	Message string
}

func fromProtoModelRoom(rooms []*booking.Room) []Room {
	result := make([]Room, len(rooms))
	for i, room := range rooms {
		result[i] = Room{
			Id:          room.Id,
			RoomNumber:  room.RoomNumber,
			Description: room.Description,
			Price:       room.Price,
		}
	}
	return result
}

func fromProtoModelBooking(bookings []*booking.Booking) []Booking {
	result := make([]Booking, len(bookings))
	for i, booking := range bookings {
		result[i] = Booking{
			Id:        booking.Id,
			RoomId:    booking.RoomId,
			UserId:    booking.UserId,
			StartDate: booking.StartDate,
			EndDate:   booking.EndDate,
		}
	}
	return result
}
