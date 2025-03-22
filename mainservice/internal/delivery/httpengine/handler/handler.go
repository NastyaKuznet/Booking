package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	roomClient RoomClient
}

func NewHandler(client RoomClient) *Handler {
	return &Handler{roomClient: client}
}

type Room struct {
	Id          int64   `json:"id"`
	RoomNumber  string  `json:"roomnumber"`
	Description string  `json:"description"`
	Available   bool    `json:"available"`
	Price       float64 `json:"price"`
}

func (h *Handler) GetAllRooms(c *gin.Context) {
	rooms, err := h.roomClient.GetRooms(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, rooms)
}
