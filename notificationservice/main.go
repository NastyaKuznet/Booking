package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "notificationservice/api/proto" // Импорт сгенерированного пакета

	"google.golang.org/grpc"
)

// Реализация сервиса уведомлений
type NotificationServer struct { //Создается структура NotificationServer, которая будет обрабатывать запросы к gRPC-сервису
	pb.UnimplementedNotificationServiceServer //Это заготовка, автоматически созданная при генерации кода из .proto
}

// Метод для отправки уведомления о бронировании
func (s *NotificationServer) SendBookingNotification(ctx context.Context, req *pb.BookingNotificationRequest) (*pb.BookingNotificationResponse, error) {
	log.Printf("Отправка уведомления: %+v\n", req) //Записывает данные запроса (req) в логи для отладки.

	// Формирование сообщения
	message := fmt.Sprintf("Номер %s забронирован гостем %s на даты: %s - %s",
		req.RoomId, req.GuestName, req.CheckInDate, req.CheckOutDate)

	fmt.Println(message) //Выводит сообщение в консоль

	return &pb.BookingNotificationResponse{ //Указывает, что всё прошло успешно
		Success: true,
		Message: "Уведомление успешно отправлено!",
	}, nil
}

func main() {
	// Создаем gRPC-сервер
	port := ":50051" // Порт для gRPC
	lis, err := net.Listen("tcp", port) //создает TCP-соединение для прослушивания входящих запросов.
	if err != nil {
		log.Fatalf("Не удалось прослушать порт %s: %v", port, err)
	}

	s := grpc.NewServer()
	pb.RegisterNotificationServiceServer(s, &NotificationServer{}) //Регистрируется сервис NotificationService с помощью метода RegisterNotificationServiceServer
	//Сервер теперь знает, как обрабатывать запросы к методу SendBookingNotification

	log.Printf("Запуск gRPC-сервера уведомлений на %s...\n", port)
	if err := s.Serve(lis); err != nil { //Сервер начинает прослушивать входящие запросы через вызов s.Serve(lis)
		log.Fatalf("Не удалось запустить сервер: %v", err)
	}
}
