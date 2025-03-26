package main

import (
	"context"
	"log"
	"time"

	pb "notificationservice/api/proto" // Импорт сгенерированного пакета

	"google.golang.org/grpc"                      //Основная библиотека для создания gRPC-клиента
	"google.golang.org/grpc/credentials/insecure" //Библиотека для подключения к серверу без TLS-шифрования
)

func main() {
	// Создается подключение к gRPC-серверу
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials())) //Указывает, что подключение будет не защищенным
	if err != nil {
		log.Fatalf("Не удалось подключиться к серверу: %v", err)
	}
	defer conn.Close() //После завершения работы программы, подключение к серверу закрывается автоматически

	client := pb.NewNotificationServiceClient(conn) //Создается клиент для взаимодействия с gRPC-сервисом NotificationService

	ctx, cancel := context.WithTimeout(context.Background(), time.Second) //Создается базовый контекст для запроса, если сервер не ответит в течение 1 секунды, запрос завершится с ошибкой.

	defer cancel() //После завершения запроса освобождается контекст

	// Данные для отправки
	req := &pb.BookingNotificationRequest{
		RoomId:       "101",
		GuestName:    "Иван Петров",
		CheckInDate:  "2025-12-01", //заезд
		CheckOutDate: "2025-12-05", //выезд
	}

	// Отправка запроса
	resp, err := client.SendBookingNotification(ctx, req) //Через созданного клиента вызывается метод SendBookingNotification gRPC-сервера
	if err != nil {
		log.Fatalf("Ошибка при вызове метода: %v", err)
	}

	// Вывод ответа
	log.Printf("Ответ сервера: %+v\n", resp)
}
