package httpengine

import (
	"context"
	"mainservice/internal/delivery/httpengine/handler"

	"github.com/gin-gonic/gin"
)

func InitRouter(ctx context.Context, config Config, handler *handler.Handler) *gin.Engine {
	router := gin.Default()
	router.GET("/rooms", handler.GetAllRooms)
	router.GET("/rooms/avalaible", handler.GetAvailableRooms)
	router.POST("/booking", handler.BookRoom)
	router.DELETE("/booking", handler.CancelBooking)
	return router
}
