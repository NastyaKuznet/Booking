package grpcclient

import "mainservice/product/booking"

type Room struct {
	Id          int64
	RoomNumber  string
	Description string
	Available   bool
	Price       float64
}

func fromProtoModel(rooms []*booking.Room) []Room {
	result := make([]Room, len(rooms))
	for i, room := range rooms {
		result[i] = Room{
			Id:          room.Id,
			RoomNumber:  room.Name,
			Description: room.Description,
			Available:   room.Available,
			Price:       0,
		}
	}
	return result
}
