package main

import (
	"context"
	"fmt"
	"log/slog"
	"mainservice/internal/core/usecase"
	"mainservice/internal/delivery/httpengine"
	"mainservice/internal/delivery/httpengine/handler"
	"mainservice/internal/lib/authclient"
	"mainservice/internal/lib/booking"
	"mainservice/internal/lib/rabbitclient"
	"mainservice/internal/repository/datasources/grpc"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v3"
)

type Config struct {
	BookingService booking.Config      `yaml:"booking_service"`
	Router         httpengine.Config   `yaml:"router"`
	Rabbit         rabbitclient.Config `yaml:"rabbit"`
	AuthService    authclient.Config   `yaml:"auth_service"`
}

func loadConfig(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var config Config
	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		return nil, err
	}
	return &config, nil
}

func main() {
	gin.SetMode(gin.DebugMode)

	cfg, err := loadConfig("config.yml")
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	bookingClient := booking.NewClient(cfg.BookingService)
	authClient := authclient.NewClient(cfg.AuthService)

	bookingRepo := grpc.NewBookingRepository(bookingClient)
	authRepo := grpc.NewAuthRepository(authClient)

	notificationClient := rabbitclient.NewRabbirClient(cfg.Rabbit)

	bookingUc := usecase.NewUsecaseBooking(bookingRepo, notificationClient, authRepo)
	authUc := usecase.NewUsecaseAuth(authRepo, notificationClient)

	bookingHandler := handler.NewHandlerBooking(bookingUc)
	authHandler := handler.NewHandlerAuth(authUc)

	router := httpengine.InitRouter(context.Background(), cfg.Router, bookingHandler, authHandler)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", cfg.Router.Port), router); err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
}
