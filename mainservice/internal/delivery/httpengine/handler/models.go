package handler

import "mainservice/internal/core/entity"

type Room struct {
	Id          int64  `json:"id"`
	RoomNumber  string `json:"roomnumber"`
	Description string `json:"description"`
	Available   bool   `json:"available"`
	Price       int64  `json:"price"`
}

type BookingRoom struct {
	RoomId    int64  `json:"roomId"`
	UserId    int64  `json:"userId"`
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
}

type BookingRoomState struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type BookingId struct {
	BookingId int64 `json:"bookingId"`
}

type CancelingBookingState struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func fromEntityRooms(rooms []entity.Room) []Room {
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
