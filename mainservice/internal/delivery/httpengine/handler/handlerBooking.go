package handler

import (
	"mainservice/internal/core/entity"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HandlerBooking struct {
	bookingClient BookingClient
}

func NewHandlerBooking(client BookingClient) *HandlerBooking {
	return &HandlerBooking{bookingClient: client}
}

func (h *HandlerBooking) GetAllRooms(c *gin.Context) {
	var req Token

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	rooms, err := h.bookingClient.GetAllRooms(c.Request.Context(), req.Token)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, fromEntityRooms(rooms))
}

func (h *HandlerBooking) GetAvailableRooms(c *gin.Context) {
	var req AvailableRoom

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	rooms, err := h.bookingClient.GetAvailableRooms(c.Request.Context(), req.StartDate, req.EndDate, req.Token)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, fromEntityRooms(rooms))
}

func (h *HandlerBooking) BookRoom(c *gin.Context) {
	var req BookingRoom

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	bookingRoomState, err := h.bookingClient.BookRoom(c.Request.Context(), entity.BookingRoom{
		RoomId:    req.RoomId,
		UserId:    req.UserId,
		StartDate: req.StartDate,
		EndDate:   req.EndDate,
	}, req.Token)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, BookingRoomState{
		Success: bookingRoomState.Success,
		Message: bookingRoomState.Message,
	})
}

func (h *HandlerBooking) CancelBooking(c *gin.Context) {
	var req BookingId

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	cancelBookingState, err := h.bookingClient.CancelBooking(c.Request.Context(), req.BookingId, req.Token)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, CancelingBookingState{
		Success: cancelBookingState.Success,
		Message: cancelBookingState.Message,
	})
}

func (h *HandlerBooking) GetAllBookings(c *gin.Context) {
	var req Token

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	bookings, err := h.bookingClient.GetAllBookings(c.Request.Context(), req.Token)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, fromEntityBookings(bookings))
}
