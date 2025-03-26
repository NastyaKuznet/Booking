package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/streadway/amqp"
)

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

func webHandler(msgs <-chan amqp.Delivery) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		timeout := time.After(5 * time.Second) // Установите таймаут в 5 секунд
		messageReceived := false               // Флаг для проверки, были ли сообщения

		// Создаем селектор для чтения сообщений или истечения таймаута
		select {
		case d := <-msgs:
			fmt.Fprintf(w, "Получено сообщение: %s\n", d.Body)
			messageReceived = true
		case <-timeout:
			if !messageReceived {
				fmt.Fprintln(w, "Нет новых сообщений в очереди.")
			}
		}

		// Завершаем ответ
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "Обработка завершена.")
	}
}

func startWebServer() {
	conn, channel, msgs, err := initRabbitMQForWeb()
	if err != nil {
		log.Fatalf("Ошибка при инициализации RabbitMQ для веб-сервера: %v", err)
	}
	defer conn.Close()
	defer channel.Close()

	http.HandleFunc("/queue", webHandler(msgs))

	log.Println("Запуск веб-сервера на :8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
