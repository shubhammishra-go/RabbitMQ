package main

import (
	"context"
	"log"
	"time"

	"github.com/shubhammishra-go/internal"
	"golang.org/x/sync/errgroup"
)

func main() {

	conn, err := internal.ConnectRabbitMQ("learn", "learn", "localhost:5672", "customers")

	if err != nil {
		panic(err)
	}

	defer conn.Close()

	client, err := internal.NewRabbitMQClient(conn)

	if err != nil {
		panic(err)
	}

	defer client.Close()

	messageBus, err := client.Consume("customers_created", "email_service", false)

	if err != nil {
		panic(err)
	}

	// Way #1 of consuming messages

	// go func() {

	// 	for message := range messageBus {
	// 		log.Printf("New Message : %v\n", message)
	// 	}
	// }()

	// log.Println("Consuming messages, to close the program CTRL+C")

	// Way #2 of consuming messages

	var blocking chan struct{}

	// seting timeout for 15 sec. per tasks

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 15*time.Second)

	defer cancel()

	g, ctx := errgroup.WithContext(ctx)

	// errgroups allows us concurrent tasks

	g.SetLimit(10)

	go func() {

		for message := range messageBus {

			msg := message

			g.Go(
				func() error {

					log.Printf("New Message: %v", msg)
					time.Sleep(10 * time.Second)

					return nil

				},
			)
		}

	}()

	log.Println("Consuming messages, to close the program CTRL+C")

	<-blocking

}
