package httpengine

import (
	"context"
	"mainservice/internal/delivery/httpengine/handler"

	"github.com/gin-gonic/gin"
)

func InitRouter(ctx context.Context, config Config, handler *handler.Handler) *gin.Engine {
	router := gin.Default()
	router.GET("/rooms", handler.GetAllRooms)
	//тут дальше остальные эндпоинты
	return router
}
