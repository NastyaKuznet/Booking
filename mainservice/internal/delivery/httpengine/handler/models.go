package handler

import "mainservice/internal/core/entity"

type Token struct {
	Token string `json:"token"`
}

type Room struct {
	Id          int64  `json:"id"`
	RoomNumber  string `json:"roomnumber"`
	Description string `json:"description"`
	Price       int64  `json:"price"`
}

type BookingRoom struct {
	Token     string `json:"token"`
	RoomId    int64  `json:"roomId"`
	UserId    int64  `json:"userId"`
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
}

type Booking struct {
	Id        int64  `json:"id"`
	RoomId    int64  `json:"roomId"`
	UserId    int64  `json:"userId"`
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
}

type AvailableRoom struct {
	Token     string `json:"token"`
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
}

type BookingRoomState struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type BookingId struct {
	Token     string `json:"token"`
	BookingId int64  `json:"bookingId"`
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
			Price:       room.Price,
		}
	}
	return result
}

func fromEntityBookings(bookings []entity.Booking) []Booking {
	result := make([]Booking, len(bookings))
	for i, booking := range bookings {
		result[i] = Booking{
			Id:        booking.Id,
			UserId:    booking.UserId,
			RoomId:    booking.RoomId,
			StartDate: booking.StartDate,
			EndDate:   booking.EndDate,
		}
	}
	return result
}
