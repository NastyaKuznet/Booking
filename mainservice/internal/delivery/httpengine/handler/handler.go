package handler

import (
	"mainservice/internal/core/entity"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	bookingClient BookingClient
}

func NewHandler(client BookingClient) *Handler {
	return &Handler{bookingClient: client}
}

type Room struct {
	Id          int64  `json:"id"`
	RoomNumber  string `json:"roomnumber"`
	Description string `json:"description"`
	Available   bool   `json:"available"`
	Price       int64  `json:"price"`
}

func fromEntity(rooms []entity.Room) []Room {
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

func (h *Handler) GetAllRooms(c *gin.Context) {
	rooms, err := h.bookingClient.GetAllRooms(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, fromEntity(rooms))
}
