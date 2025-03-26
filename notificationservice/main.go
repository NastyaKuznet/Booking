package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "notificationservice/api/proto" // Импорт сгенерированного пакета

	"github.com/streadway/amqp" // Добавляем RabbitMQ
	"google.golang.org/grpc"
)

// Реализация сервиса уведомлений
type NotificationServer struct {
	pb.UnimplementedNotificationServiceServer
	rabbitChannel *amqp.Channel // Канал RabbitMQ
}

// Метод для отправки уведомления о бронировании
func (s *NotificationServer) SendBookingNotification(ctx context.Context, req *pb.BookingNotificationRequest) (*pb.BookingNotificationResponse, error) {
	log.Printf("Отправка уведомления: %+v\n", req)

	// Формирование сообщения
	message := fmt.Sprintf("Номер %s забронирован гостем %s на даты: %s - %s",
		req.RoomId, req.GuestName, req.CheckInDate, req.CheckOutDate)

	// Отправка сообщения в RabbitMQ
	err := s.rabbitChannel.Publish(
		"",        // exchange
		"booking", // routing key (имя очереди)
		false,     // mandatory
		false,     // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		},
	)
	if err != nil {
		return nil, fmt.Errorf("не удалось отправить сообщение в RabbitMQ: %v", err)
	}

	fmt.Println(message)

	return &pb.BookingNotificationResponse{
		Success: true,
		Message: "Уведомление успешно отправлено!",
	}, nil
}

func initRabbitMQ() (*amqp.Channel, error) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		return nil, fmt.Errorf("не удалось подключиться к RabbitMQ: %v", err)
	}
	channel, err := conn.Channel()
	if err != nil {
		return nil, fmt.Errorf("не удалось открыть канал RabbitMQ: %v", err)
	}

	// Создание очереди
	_, err = channel.QueueDeclare(
		"booking", // имя очереди
		true,      // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		return nil, fmt.Errorf("не удалось объявить очередь: %v", err)
	}

	return channel, nil
}

func main() {
	port := ":50051"
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Не удалось прослушать порт %s: %v", port, err)
	}

	s := grpc.NewServer()

	// Инициализация RabbitMQ
	rabbitChannel, err := initRabbitMQ()
	if err != nil {
		log.Fatalf("Ошибка при инициализации RabbitMQ: %v", err)
	}

	pb.RegisterNotificationServiceServer(s, &NotificationServer{rabbitChannel: rabbitChannel})

	go startWebServer()

	log.Printf("Запуск gRPC-сервера уведомлений на %s...\n", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Не удалось запустить сервер: %v", err)
	}
}
