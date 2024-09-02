package main

import (
	"context"
	"log"
	"time"

	"github.com/rabbitmq/amqp091-go"
	"github.com/shubhammishra-go/internal"
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

	// Creating Q1 as durable queue

	err = client.CreateQueue("customers_created", true, false)

	if err != nil {
		panic(err)
	}

	// Creating Q2 as not durable queue, so for every restart of rabbimq server it will be deleted

	err = client.CreateQueue("customers_test", false, false)

	if err != nil {
		panic(err)
	}

	// Creating Bind

	err = client.CreateBinding("customers_created", "customers.created.*", "customer_events")

	if err != nil {
		panic(err)
	}

	// Creating Another Binding to "customers_test" Queue

	err = client.CreateBinding("customers_test", "customers.*", "customer_events")

	if err != nil {
		panic(err)
	}

	//context

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	// Sending a Persistant Message

	err = client.Send(ctx, "customer_events", "customers.created.us", amqp091.Publishing{
		ContentType:  "text/plain",
		DeliveryMode: amqp091.Persistent,
		Body:         []byte(`An cool message between services`),
	})

	if err != nil {
		panic(err)
	}

	// Sending a Transient Message

	err = client.Send(ctx, "customer_events", "customers.test", amqp091.Publishing{
		ContentType:  "text/plain",
		DeliveryMode: amqp091.Transient,
		Body:         []byte(`An Uncool message between services`),
	})

	if err != nil {
		panic(err)
	}

	time.Sleep(10 * time.Second)

	log.Println(client)

}
