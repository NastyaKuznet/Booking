package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"

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

// Эндпоинт для приема данных через HTTP
func bookingHandler(rabbitChannel *amqp.Channel) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
			return
		}

		// Читаем данные из формы
		r.ParseForm()
		roomID := r.FormValue("room_id")
		guestName := r.FormValue("guest_name")
		checkInDate := r.FormValue("check_in_date")
		checkOutDate := r.FormValue("check_out_date")

		// Проверяем, что все поля заполнены
		if roomID == "" || guestName == "" || checkInDate == "" || checkOutDate == "" {
			http.Error(w, "All fields are required", http.StatusBadRequest)
			return
		}

		// Формирование сообщения
		message := fmt.Sprintf("Номер %s забронирован гостем %s на даты: %s - %s",
			roomID, guestName, checkInDate, checkOutDate)

		// Отправка сообщения в RabbitMQ
		err := rabbitChannel.Publish(
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
			http.Error(w, "Failed to send message to RabbitMQ", http.StatusInternalServerError)
			return
		}

		// Отправка ответа клиенту
		w.WriteHeader(http.StatusCreated)
		fmt.Fprintf(w, "Бронь успешно добавлена: %s\n", message)
	}
}

func startWebServer(rabbitChannel *amqp.Channel) {
	http.HandleFunc("/queue", func(w http.ResponseWriter, r *http.Request) {
		conn, channel, _, err := initRabbitMQForWeb() // Принимаем все 4 значения
		if err != nil {
			http.Error(w, "Ошибка при инициализации RabbitMQ", http.StatusInternalServerError)
			return
		}
		defer conn.Close()
		defer channel.Close()

		// Создаем срез для хранения сообщений
		var messages []string

		// Читаем все доступные сообщения из очереди
		for i := 0; i < 10; i++ { // Ограничиваем количество попыток чтения (например, до 10)
			delivery, ok, err := channel.Get("booking", false) // Извлекаем одно сообщение из очереди
			if err != nil || !ok {
				break // Если сообщений больше нет, выходим из цикла
			}

			// Подтверждаем получение сообщения
			err = delivery.Ack(false)
			if err != nil {
				log.Printf("Ошибка подтверждения сообщения: %v", err)
			}

			// Добавляем тело сообщения в список
			messages = append(messages, string(delivery.Body))
		}

		// Если сообщений нет, показываем уведомление
		if len(messages) == 0 {
			fmt.Fprintln(w, "Нет сообщений в очереди.")
		} else {
			// Выводим список сообщений
			for _, msg := range messages {
				fmt.Fprintf(w, "Получено сообщение: %s\n", msg)
			}
		}

		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "Конец списка сообщений.")
	})

	http.HandleFunc("/book", bookingHandler(rabbitChannel))

	// Страница с формой для ввода данных
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `
            <!DOCTYPE html>
            <html>
            <head>
                <title>Бронирование номеров</title>
            </head>
            <body>
                <h1>Бронирование номеров</h1>
                <form action="/book" method="post">
                    <label for="room_id">ID номера:</label><br>
                    <input type="text" id="room_id" name="room_id" required><br>
                    
                    <label for="guest_name">Имя гостя:</label><br>
                    <input type="text" id="guest_name" name="guest_name" required><br>
                    
                    <label for="check_in_date">Дата заезда:</label><br>
                    <input type="date" id="check_in_date" name="check_in_date" required><br>
                    
                    <label for="check_out_date">Дата выезда:</label><br>
                    <input type="date" id="check_out_date" name="check_out_date" required><br>
                    
                    <input type="submit" value="Забронировать">
                </form>
            </body>
            </html>
        `)
	})

	log.Println("Запуск веб-сервера на :8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func initRabbitMQForWeb() (*amqp.Connection, *amqp.Channel, <-chan amqp.Delivery, error) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		return nil, nil, nil, fmt.Errorf("не удалось подключиться к RabbitMQ: %v", err)
	}
	channel, err := conn.Channel()
	if err != nil {
		return nil, nil, nil, fmt.Errorf("не удалось открыть канал RabbitMQ: %v", err)
	}

	// Объявление очереди
	queue, err := channel.QueueDeclare(
		"booking", // имя очереди
		true,      // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("не удалось объявить очередь: %v", err)
	}

	// Подписка на очередь
	msgs, err := channel.Consume(
		queue.Name, // очередь
		"",         // consumer
		true,       // auto-ack
		false,      // exclusive
		false,      // no-local
		false,      // no-wait
		nil,        // args
	)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("не удалось подписаться на очередь: %v", err)
	}

	return conn, channel, msgs, nil
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

	// Запуск веб-сервера
	go startWebServer(rabbitChannel)

	log.Printf("Запуск gRPC-сервера уведомлений на %s...\n", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Не удалось запустить сервер: %v", err)
	}
}
