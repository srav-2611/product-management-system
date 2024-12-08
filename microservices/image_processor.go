package main

import (
	"context"
	"encoding/json"
	"log"
	"os" // Ensure this is imported
	"strings"
	"product-management-system/config"

	"github.com/streadway/amqp"
)

// ImageTask represents the structure of the message from RabbitMQ
type ImageTask struct {
	ProductID int    `json:"product_id"`
	ImageURL  string `json:"image_url"`
}

func main() {
	// Connect to the database
	config.ConnectDatabase()
	defer config.CloseDatabase()

	// Connect to RabbitMQ
	rabbitMQURL := os.Getenv("RABBITMQ_URL")
	if rabbitMQURL == "" {
		rabbitMQURL = "amqp://guest:guest@rabbitmq:5672/"
	}
	conn, err := amqp.Dial(rabbitMQURL)
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
	}
	defer ch.Close()

	// Declare the queue
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
	}

	// Start consuming messages
	msgs, err := ch.Consume(
		queueName,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to register a consumer: %v", err)
	}

	log.Printf("Waiting for messages in queue: %s", queueName)

	// Process messages
	forever := make(chan bool)
	go func() {
		for d := range msgs {
			var task ImageTask
			err := json.Unmarshal(d.Body, &task)
			if err != nil {
				log.Printf("Failed to unmarshal message: %v", err)
				continue
			}

			// Simulate image compression
			compressedImage := simulateImageCompression(task.ImageURL)

			// Update the database with the compressed image
			err = updateCompressedImages(task.ProductID, compressedImage)
			if err != nil {
				log.Printf("Failed to update compressed images for product %d: %v", task.ProductID, err)
			} else {
				log.Printf("Successfully processed image for product %d: %s", task.ProductID, compressedImage)
			}
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}

// simulateImageCompression simulates compressing an image and returns the new URL
func simulateImageCompression(imageURL string) string {
	// This simulates adding a "-compressed" suffix to the image URL
	return strings.Replace(imageURL, ".jpg", "-compressed.jpg", 1)
}

// updateCompressedImages updates the compressed images in the database
func updateCompressedImages(productID int, compressedImage string) error {
	query := `
		UPDATE products
		SET compressed_product_images = array_append(compressed_product_images, $1)
		WHERE id = $2
	`
	_, err := config.DB.Exec(context.Background(), query, compressedImage, productID)
	return err
}
