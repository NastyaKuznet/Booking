package main

import (
	"context"
	"fmt"
	"log/slog"
	"mainservice/internal/core/usecase"
	"mainservice/internal/delivery/httpengine"
	"mainservice/internal/delivery/httpengine/handler"
	"mainservice/internal/lib/grpcclient"
	"mainservice/internal/lib/rabbitclient"
	"mainservice/internal/repository/datasources/grpc"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v3"
)

type Config struct {
	BookingService grpcclient.Config   `yaml:"booking_service"`
	Router         httpengine.Config   `yaml:"router"`
	Rabbit         rabbitclient.Config `yaml:"rabbit"`
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

	cfg, err := loadConfig("config.yaml")
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	bookingClient := grpcclient.NewClient(cfg.BookingService)

	repo := grpc.NewBookingRepository(bookingClient)

	notificationClient := rabbitclient.NewRabbirClient(cfg.Rabbit)

	uc := usecase.NewUsecase(repo, notificationClient)

	bookingHandler := handler.NewHandler(uc)

	router := httpengine.InitRouter(context.Background(), cfg.Router, bookingHandler)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", cfg.Router.Port), router); err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
}
