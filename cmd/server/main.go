package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	fmt.Println("Starting Peril server...")
	const rabbitmqURL = "amqp://guest:guest@localhost:5672/"

	conn, err := amqp.Dial(rabbitmqURL)
	if err != nil {
		fmt.Printf("Failed to connect to RabbitMQ: %v\n", err)
		return
	}
	defer conn.Close()
	fmt.Println("Connected to RabbitMQ")

	sigs := make(chan os.Signal, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	done := make(chan bool, 1)

	go func() {
		sig := <-sigs
		fmt.Printf("\nReceived signal: %v. Shutting down Peril server...\n", sig)
		done <- true
	}()

	fmt.Println("Peril server is running. Press Ctrl+C to stop.")
	<-done
	fmt.Println("Peril server stopped.")
}
