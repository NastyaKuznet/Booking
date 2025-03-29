package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	pb "notificationservice/api/proto" // Импорт сгенерированного пакета

	amqp "github.com/rabbitmq/amqp091-go"
	"google.golang.org/grpc" // Библиотека для работы с gRPC
)

// Реализация сервиса уведомлений
type NotificationServer struct {
	pb.UnimplementedNotificationServiceServer               // Автоматически созданная заготовка для gRPC-сервиса
	rabbitChannel                             *amqp.Channel // Канал RabbitMQ
}

// Метод для отправки уведомления о бронировании через gRPC
func (s *NotificationServer) SendBookingNotification(ctx context.Context, req *pb.BookingNotificationRequest) (*pb.BookingNotificationResponse, error) {
	log.Printf("Отправка уведомления: %+v\n", req)

	// Формирование сообщения
	message := fmt.Sprintf("Номер %s забронирован гостем %s на даты: %s - %s",
		req.RoomId, req.GuestName, req.CheckInDate, req.CheckOutDate)

	// Отправляем сообщение в очередь RabbitMQ
	err := s.rabbitChannel.Publish(
		"",        // Default exchange
		"booking", // Routing key (имя очереди)
		false,     // Mandatory
		false,     // Immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		},
	)
	if err != nil {
		return nil, fmt.Errorf("не удалось отправить сообщение в RabbitMQ: %v", err)
	}

	return &pb.BookingNotificationResponse{
		Success: true,
		Message: "Уведомление успешно отправлено!",
	}, nil
}

// Инициализация соединения с RabbitMQ
func initRabbitMQ() (*amqp.Channel, error) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		return nil, fmt.Errorf("не удалось подключиться к RabbitMQ: %v", err)
	}
	channel, err := conn.Channel()
	if err != nil {
		return nil, fmt.Errorf("не удалось открыть канал RabbitMQ: %v", err)
	}

	// Создаем очередь "booking"
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

// Подписываемся на очередь RabbitMQ и выводим сообщения в консоль
func consumeFromRabbitMQ(channel *amqp.Channel) {
	msgs, err := channel.Consume(
		"booking", // имя очереди
		"",        // consumer tag
		true,      // auto-ack
		false,     // exclusive
		false,     // no-local
		false,     // no-wait
		nil,       // args
	)
	if err != nil {
		log.Fatalf("Не удалось подписаться на очередь: %v", err)
	}

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		for msg := range msgs {
			log.Printf("Получено сообщение: %s", string(msg.Body))
		}
	}()

	log.Printf(" [*] Ожидание сообщений. Для выхода нажмите CTRL+C")
	<-sig
	log.Println("Завершение работы...")
}

func main() {
	port := ":50051" // Запускаем gRPC-сервер на порту 50051
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Не удалось прослушать порт %s: %v", port, err)
	}

	s := grpc.NewServer()

	// Инициализируем соединение с RabbitMQ
	rabbitChannel, err := initRabbitMQ()
	if err != nil {
		log.Fatalf("Ошибка при инициализации RabbitMQ: %v", err)
	}

	// Регистрируем gRPC-сервис NotificationService
	pb.RegisterNotificationServiceServer(s, &NotificationServer{rabbitChannel: rabbitChannel})

	// Запускаем подписчик RabbitMQ в отдельной горутине
	go consumeFromRabbitMQ(rabbitChannel)

	log.Printf("Запуск gRPC-сервера уведомлений на %s...\n", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Не удалось запустить сервер: %v", err)
	}
}
