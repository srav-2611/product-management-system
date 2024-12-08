package queue

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

// PublishImageProcessingTask sends an image processing task to RabbitMQ
func PublishImageProcessingTask(productID int, imageURL string) error {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
		return err
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
		return err
	}
	defer ch.Close()

	queueName := "image_processing"
	_, err = ch.QueueDeclare(
		queueName,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to declare a queue: %v", err)
		return err
	}

	body := fmt.Sprintf(`{"product_id": %d, "image_url": "%s"}`, productID, imageURL)
	err = ch.Publish(
		"",
		queueName,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        []byte(body),
		},
	)
	if err != nil {
		log.Fatalf("Failed to publish a message: %v", err)
		return err
	}

	log.Printf("Published image processing task: %s", body)
	return nil
}
