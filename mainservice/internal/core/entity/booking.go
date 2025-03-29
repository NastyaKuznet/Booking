package entity

import (
	"mainservice/internal/lib/booking"
)

type Room struct {
	Id          int64
	RoomNumber  string
	Description string
	Available   bool
	Price       int64
}

type BookingRoom struct {
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

func FromProtoModelRooms(rooms []*booking.Room) []Room {
	result := make([]Room, len(rooms))
	for i, room := range rooms {
		result[i] = Room{
			Id:          room.Id,
			RoomNumber:  room.RoomNumber,
			Description: room.Description,
			Available:   room.Available,
			Price:       room.Price,
		}
	}
	return result
}
