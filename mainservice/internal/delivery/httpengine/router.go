package httpengine

import (
	"context"
	"mainservice/internal/delivery/httpengine/handler"

	"github.com/gin-gonic/gin"
)

func InitRouter(
	ctx context.Context,
	config Config,
	handlerBooking *handler.HandlerBooking,
	handlerAuth *handler.HandlerAuth,
) *gin.Engine {
	router := gin.Default()
	router.GET("/rooms", handlerBooking.GetAllRooms)
	router.GET("/rooms/avalaible", handlerBooking.GetAvailableRooms)
	router.GET("/rooms/booking", handlerBooking.GetAllBookings)
	router.POST("/booking", handlerBooking.BookRoom)
	router.DELETE("/booking", handlerBooking.CancelBooking)
	router.POST("/regist", handlerAuth.RegisterUser)
	router.POST("/login", handlerAuth.LoginUser)
	return router
}
