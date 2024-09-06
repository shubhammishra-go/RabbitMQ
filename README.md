# RabbitMQ

RabbitMQ is an open-source `message-broker` software (sometimes called `message-oriented middleware`) that originally implemented the `Advanced Message Queuing Protocol (AMQP)` and has since been extended with a plug-in architecture to support `Streaming Text Oriented Messaging Protocol (STOMP)`, `MQ Telemetry Transport (MQTT)`, and other protocols.

![alt text](image.png)


It is written in `Erlang` Language, the RabbitMQ server is built on the `Open Telecom Platform framework` for clustering and failover. Client libraries to interface with the broker are available for all major programming languages. The source code is released under the `Mozilla Public License`.


Since November 2020, there are commercial offerings available of RabbitMQ, for support and enterprise features: "VMware RabbitMQ OVA", "VMware RabbitMQ" and "VMware RabbitMQ for Kubernetes" (different feature levels) Open-Source RabbitMQ is also packaged by Bitnami and commercially for VMware's Tanzu Application Service. 

To know more about Message Queue

https://github.com/shubhammishra-go/Message-Queue


## History

Originally developed by `Rabbit Technologies Ltd`. which started as a `joint venture between LShift and CohesiveFT in 2007`, RabbitMQ was acquired in April 2010 by `SpringSource`, a division of `VMware`. 

The project became part of Pivotal Software in May 2013. Which then got acquired back by VMWare in December 2019.

The project consists of: 

- The RabbitMQ exchange server
    
- Gateways for AMQP, HTTP, STOMP, and MQTT protocols

- AMQP client libraries for Java, .NET Framework and Erlang. (AMQP clients for other languages are available from other vendors.)

- A plug-in platform for extensibility, with a predefined collection of supported plug-ins, including:
        
- - A "Shovel" plug-in that takes care of moving or copying (replicating) messages from one broker to another.

- - A "Federation" plug-in that enables efficient sharing of messages between brokers (at the exchange level).

- - A "Management" plug-in that enables monitoring and control of brokers and clusters of brokers.


## Which protocols does RabbitMQ support?

RabbitMQ supports several messaging protocols, directly and through the use of plugins. 

- `AMQP 0-9-1 and extensions`

RabbitMQ was originally developed to support AMQP 0-9-1. As such this protocol is the "core" protocol supported by the broker. All of these variants are fairly similar to each other, with later versions tidying up unclear or unhelpful parts of earlier versions. We have extended AMQP 0-9-1 in various ways.

AMQP 0-9-1 is a binary protocol, and defines quite strong messaging semantics. For clients it's a reasonably easy protocol to implement, and as such there are a large number of client libraries available for many different programming languages and environments.

AMQP 0-9-1 is the protocol used by `RabbitMQ tutorials`.

- `Streaming Text Oriented Messaging Protocol (STOMP)`

STOMP is a text-based messaging protocol emphasising (protocol) simplicity. It defines little in the way of messaging semantics, but is easy to implement and very easy to implement partially (it's the only protocol that can be used by hand over telnet).

RabbitMQ supports STOMP (all current versions) via a `plugin`.


- `Message Queue Telemetry Transport (MQTT)`

MQTT is a binary protocol emphasising lightweight publish / subscribe messaging, targeted towards clients in constrained devices. It has well defined messaging semantics for publish / subscribe, but not for other messaging idioms.

RabbitMQ supports MQTT versions 3.1, 3.1.1, and 5.0 via a `plugin`.

- `AMQP 1.0`

Despite the name, AMQP 1.0 is a radically different protocol from AMQP 0-9-1 / 0-9 / 0-8, `sharing essentially nothing at the wire level`. `AMQP 1.0 imposes far fewer semantic requirements`; it is therefore easier to add support for AMQP 1.0 to existing brokers. The protocol is substantially more complex than AMQP 0-9-1, and there are fewer client implementations.

RabbitMQ supports AMQP 1.0 via a `plugin`.

- `HTTP and WebSockets`

While HTTP is not really a messaging protocol, RabbitMQ can transmit messages over HTTP in three ways:


- - The Web STOMP plugin supports STOMP messaging to the browser using WebSockets.

- - The Web MQTT plugin supports MQTT messaging to the browser using WebSockets.

- - The management plugin supports a simple HTTP API to send and receive messages. This is primarily intended for diagnostic purposes but can be used for low volume messaging without reliable delivery.

- `RabbitMQ Streams`

The RabbitMQ Streams protocol allows communicating with streams. RabbitMQ supports the streams protocol via a `plugin`.


# Why RabbitMQ?

RabbitMQ is a reliable and mature messaging and streaming broker, which is easy to deploy on cloud environments, on-premises, and on your local machine. It is currently used by millions worldwide.

- `Interoperable` able to exchange and make use of information.

RabbitMQ supports several open standard protocols, including `AMQP 1.0` and `MQTT 5`. 

There are multiple client libraries available, which can be used with your programming language of choice, just pick one. No vendor lock-in!


- `Flexible` able to be easily modified to respond to altered circumstances.

RabbitMQ provides many options you can combine to define how your messages go from the publisher to one or many consumers. `Routing`, `filtering`, `streaming`, `federation`, and so on, you name it.


- `Reliable` consistently good in quality or performance; `able to be trusted`.

With the ability to acknowledge message delivery and to replicate messages across a cluster, you can ensure your messages are safe with RabbitMQ.


# Examples of Common use cases

Here are a few common use cases of RabbitMQ

- `Decoupling interconnected services`

You have a backend service that needs to send notifications to end users. There are two notification channels: emails and push notifications for the mobile application.

The backend publishes the notification to two queues, one for each channel. Programs that manage emails and push notification subscribe to the queue they are interested in and handle notifications as soon as they arrive.

![alt text](image-1.png)

// Benefits

--> RabbitMQ absorbs the load spike.
    
--> You can do some maintenance on the notification managers without interrupting the whole service.


- `Remote Procedure Call`

You own a concert hall. Tickets for the shows are sold on multiple websites and physical kiosks. Orders from all channels must go through a complex process to determine if a customer effectively get their tickets, depending on availability. The website or kiosk expect to get an answer for the order in minimal time.

Orders are published to a queue in RabbitMQ with a correlation ID. The caller that pushed the order then subscribes to another queue and waits for an answer with that same correlation ID.

To achieve low latency, a classic queue is a good fit here but it is at the expense of less safety — the caller can still retry. If the order cannot be lost, you might prefer to use a combination of acknowledgments and quorum queues to ensure a message is safe once confirmed.

This topology allows the processing of orders to be serialized to serve them in a first come first served order. This avoids the need for transactions.

![alt text](image-2.png)

// Benefits

--> A RabbitMQ client can be a publisher and a consumer at the same time.

--> RabbitMQ can be used to `dispatch RPC calls`.


- `Streaming`

You run a video platform. When a user uploads a new video, you have multiple tasks to complete when the video is safely stored: run some post-upload analysis, transcode lower quality copies of the video, notify other users subscribed to the author's creations, and so on.

The upload service appends “New video” events to a RabbitMQ stream. Multiple backend applications can subscribe to that stream and read new events independently of each other. Users must be notified right away, but the post-upload analysis can wait and run once a day.

![alt text](image-3.png)

// Benefits

--> Streams are very efficient and avoids the need to duplicate messages.
    
--> A consumers can go back and forth in the stream even if there are concurrent consumers.

- `IoT`

You provide package delivery services across the entire galaxy. You have a swarm of space drones that need to report their status on a regular basis to a server hosted on exoplanet Kepler-438 b. Unfortunately, the network connectivity is not great…

Each space drone runs a local RabbitMQ standalone node that buffers their report until a connection is possible with the upstream RabbitMQ.

When planets are aligned, the drone's RabbitMQ shovels all reports to the upstream RabbitMQ.

![alt text](image-4.png)


// Benefits

--> RabbitMQ deployments can be chained to cater for different needs in your service, using features such as shovels and federation.

--> MQTT is well suited for millions of concurrent connections.


# Installing RabbitMQ

The latest release of RabbitMQ is `3.13.7`.  can be installed & run using Docker

```bash 
docker run -it --rm --name rabbitmq -p 5672:5672 -p 15672:15672 rabbitmq:3.13-management

```
for full fledge installtion on specific OS, go through documentation

https://www.rabbitmq.com/docs/download

RabbitMQ server running on PORT = 5672 & Its Management UI running on PORT = 15672 !

Intially, when we run our server username=guest & password=guest is already created by RabbitMQ!

we can delete it, can create a new one!



# About RabbitMQ Jargons (terms used to describe Services)

RabbitMQ is a message broker: it accepts and forwards messages. You can think about it as a post office: when you put the mail that you want posting in a post box, you can be sure that the letter carrier will eventually deliver the mail to your recipient. In this analogy, RabbitMQ is a post box, a post office, and a letter carrier.


The major difference between RabbitMQ and the post office is that it doesn't deal with paper, instead it accepts, stores, and forwards binary blobs of data ‒ messages.


RabbitMQ, and messaging in general, uses some jargon.


`Producing` means nothing more than sending. A program that sends messages is a producer :

![alt text](image-5.png)

A `Queue` is the name for the post box in RabbitMQ. Although messages flow through RabbitMQ and your applications, they can only be stored inside a queue. A queue is only bound by the host's memory & disk limits, it's essentially a large message buffer.

Many producers can send messages that go to one queue, and many consumers can try to receive data from one queue.

This is how we represent a queue:

![alt text](image-6.png)


`Consuming` has a similar meaning to receiving. A consumer is a program that mostly waits to receive messages:

![alt text](image-7.png)

Note that the producer, consumer, and broker do not have to reside on the same host; indeed in most applications they don't. An application can be both a producer and consumer, too.


# Simple RabbitMQ Program in Go

In this part of the tutorial we'll write two small programs in Go; a producer that sends a single message, and a consumer that receives messages and prints them out. 

We'll gloss over some of the detail in the Go RabbitMQ API, concentrating on this very simple thing just to get started. It's the "Hello World" of messaging.

In the diagram below, "P" is our producer and "C" is our consumer. The box in the middle is a queue - a message buffer that RabbitMQ keeps on behalf of the consumer.

![alt text](image-8.png)

The Go RabbitMQ client library

RabbitMQ speaks multiple protocols. This tutorial uses AMQP 0-9-1, which is an open, general-purpose protocol for messaging. There are a number of clients for RabbitMQ in many different languages. We'll use the Go amqp client in this tutorial.

`Note :::` Before Jumping on Code make sure RabbitMQ server is running either on Operating System or on Docker!


- `Installation of Go based amqp library`


First, install amqp go library using `go get`:

```bash
go mod init <your-module-name>
go get github.com/rabbitmq/amqp091-go
```

Now we will create two files `send.go` and `receive.go`

- `Sending Messages to Queue`

We'll call our message publisher (sender) `send.go `. The publisher(sender) will connect to RabbitMQ, send a single message, then exit.

![alt text](image-9.png)


```go

package main

//we need to import the library first:

import (
  "context"
  "log"
  "time"

  amqp "github.com/rabbitmq/amqp091-go"
)

//We also need a helper function to check the return value for each amqp call: if error exits

func failOnError(err error, msg string) {
  if err != nil {
    log.Panicf("%s: %s", msg, err)
  }
}


// then create the main function and connect to RabbitMQ server

func main() {

conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
failOnError(err, "Failed to connect to RabbitMQ")
defer conn.Close()
  
  
/*The connection abstracts the socket connection, and takes care of protocol version negotiation and authentication and so on for us. Next we create a channel, which is where most of the API for getting things done resides:*/


ch, err := conn.Channel()
failOnError(err, "Failed to open a channel")
defer ch.Close()

//To send, we must declare a queue for us to send to; then we can publish a message to the queue:

// here we will create a queue named as "hello"

q, err := ch.QueueDeclare(
  "hello", // name
  false,   // durable
  false,   // delete when unused
  false,   // exclusive
  false,   // no-wait
  nil,     // arguments
)
failOnError(err, "Failed to declare a queue")

ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()

//now we are publishing some messages to queue (note q.Name is our "hello" queue on which we are publishing message (body, which is "Hello World"))

body := "Hello World!"
err = ch.PublishWithContext(ctx,
  "",     // exchange
  q.Name, // routing key
  false,  // mandatory
  false,  // immediate
  amqp.Publishing {
    ContentType: "text/plain",
    Body:        []byte(body),
  })
failOnError(err, "Failed to publish a message")
log.Printf(" [x] Sent %s\n", body)


/*Declaring a queue is "idempotent" - it will only be created if it doesn't exist already. The message content is a byte array, so you can encode whatever you like there.*/


}

```

- `Consuming Messages from Queue`

Our consumer listens for messages from RabbitMQ, so unlike the publisher which publishes a single message, we'll keep the consumer running to listen for messages and print them out.

![alt text](image-10.png)

Create a new directory for the consumer app, like `receiver/receive.go`, to avoid a duplicate declaration.

Setting up is the same as the publisher; we open a connection and a channel, and declare the queue from which we're going to consume. Note this matches up with the queue that send publishes to.


`Note` that we declare the queue here, as well. Because we might start the consumer before the publisher, we want to make sure the queue exists before we try to consume messages from it.


```go
package main

import (
  "log"

  amqp "github.com/rabbitmq/amqp091-go"
)

func failOnError(err error, msg string) {
  if err != nil {
    log.Panicf("%s: %s", msg, err)
  }
}


func main(){

conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
failOnError(err, "Failed to connect to RabbitMQ")
defer conn.Close()

ch, err := conn.Channel()
failOnError(err, "Failed to open a channel")
defer ch.Close()

q, err := ch.QueueDeclare(
  "hello", // name
  false,   // durable
  false,   // delete when unused
  false,   // exclusive
  false,   // no-wait
  nil,     // arguments
)
failOnError(err, "Failed to declare a queue")


/* We're about to tell the server to deliver us the messages from the queue. Since it will push us messages asynchronously, we will read the messages from a channel (returned by amqp::Consume) in a goroutine. */


msgs, err := ch.Consume(
  q.Name, // queue
  "",     // consumer
  true,   // auto-ack
  false,  // exclusive
  false,  // no-local
  false,  // no-wait
  nil,    // args
)
failOnError(err, "Failed to register a consumer")

var forever chan struct{}

go func() {
  for d := range msgs {
    log.Printf("Received a message: %s", d.Body)
  }
}()

log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
<-forever


}

```


- `To Run Both Scripts`

Now we can run both scripts. In a terminal, run the publisher:

```bash 
go run send.go
```

then, run the consumer:

```bash
go run receiver/receive.go
```

The consumer will print the message it gets from the publisher via RabbitMQ. The consumer will keep running, waiting for messages (Use Ctrl-C to stop it), so try running the publisher from another terminal.


If you want to check on the queue, try using `sudo docker rabbitmqctl list_queues`.

if you are running RabbitMQ server through Docker


# Queue Description

While decalaration of Queue there are some fields.

```go
 q, err := ch.QueueDeclare(
                "task_queue", // name
                true,         // durable
                false,        // delete when unused
                false,        // exclusive
                false,        // no-wait
                nil,          // arguments
        )
```

- `Queue name` :: Queues have names so that applications can reference them. 

- `Durable` :: the queue will survive a broker restart.

- `Auto-delete` :: queue that has had at least one consumer is deleted when last consumer unsubscribes.

- `Exclusive` :: used by only one connection and the queue will be deleted when that connection closes.

- `Arguments` :: (optional); used by plugins and broker-specific features such as message TTL, queue length limit, etc.

for more informations visit ::: https://www.rabbitmq.com/docs/queues


# Work Queues

In the above tutorial (`Simple RabbitMQ Program in Go`) we wrote programs to send and receive messages from a named queue. In this one we'll create a Work Queue that will be used to distribute time-consuming tasks among multiple workers.

The main idea behind Work Queues (aka: Task Queues) is to avoid doing a resource-intensive task immediately and having to wait for it to complete. Instead we schedule the task to be done later. 

We encapsulate a task as a message and send it to a queue. A worker process running in the background will pop the tasks and eventually execute the job. 

When you run many workers the tasks will be shared between them.

This concept is especially useful in web applications where it's impossible to handle a complex task during a short HTTP request window.


In the previous part of this tutorial we sent a message containing "Hello World!". Now we'll be sending strings that stand for complex tasks. We don't have a real-world task, like images to be resized or pdf files to be rendered, so let's fake it by just pretending we're busy - by using the `time.Sleep` function. We'll take the number of dots in the string as its complexity; 

every dot will account for one second of "work". For example, a fake task described by `Hello...` will take three seconds.

We will slightly modify the `send.go` code from our previous example, to allow arbitrary messages to be sent from the command line. This program will schedule tasks to our work queue, so let's name it `new_task.go`:


```go
package main

import (
  "context"
  "log"
  "time"
  "os"

  amqp "github.com/rabbitmq/amqp091-go"
)

func failOnError(err error, msg string) {
  if err != nil {
    log.Panicf("%s: %s", msg, err)
  }
}

func bodyFrom(args []string) string {
	var s string
	if (len(args) < 2) || os.Args[1] == "" {
		s = "hello"
	} else {
		s = strings.Join(args[1:], " ")
	}
	return s
}


func main() {

conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
failOnError(err, "Failed to connect to RabbitMQ")
defer conn.Close()


ch, err := conn.Channel()
failOnError(err, "Failed to open a channel")
defer ch.Close()


q, err := ch.QueueDeclare(
  "hello", // name
  false,   // durable
  false,   // delete when unused
  false,   // exclusive
  false,   // no-wait
  nil,     // arguments
)
failOnError(err, "Failed to declare a queue")


ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()

body := bodyFrom(os.Args)
err = ch.PublishWithContext(ctx,
  "",           // exchange
  q.Name,       // routing key
  false,        // mandatory
  false,
  amqp.Publishing {
    DeliveryMode: amqp.Persistent,
    ContentType:  "text/plain",
    Body:         []byte(body),
  })
failOnError(err, "Failed to publish a message")
log.Printf(" [x] Sent %s", body)


}
```

Our old `receive.go` script also requires some changes: it needs to fake a second of work for every dot in the message body. 

It will pop messages from the queue and perform the task, so let's call it `worker.go`:


```go
package main

import (
  "log"

  amqp "github.com/rabbitmq/amqp091-go"
)

func failOnError(err error, msg string) {
  if err != nil {
    log.Panicf("%s: %s", msg, err)
  }
}


func main(){

conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
failOnError(err, "Failed to connect to RabbitMQ")
defer conn.Close()

ch, err := conn.Channel()
failOnError(err, "Failed to open a channel")
defer ch.Close()

q, err := ch.QueueDeclare(
  "hello", // name
  false,   // durable
  false,   // delete when unused
  false,   // exclusive
  false,   // no-wait
  nil,     // arguments
)
failOnError(err, "Failed to declare a queue")


msgs, err := ch.Consume(
  q.Name, // queue
  "",     // consumer
  true,   // auto-ack
  false,  // exclusive
  false,  // no-local
  false,  // no-wait
  nil,    // args
)
failOnError(err, "Failed to register a consumer")

var forever chan struct{}

go func() {
  for d := range msgs {
    log.Printf("Received a message: %s", d.Body)
    dotCount := bytes.Count(d.Body, []byte("."))
    t := time.Duration(dotCount)
    time.Sleep(t * time.Second)
    log.Printf("Done")
  }
}()

log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
<-forever


}

```

`Note :` that our fake task simulates execution time.

Run them as in tutorial one:

```shell
# shell 1
go run worker.go
```

```shell
# shell 2
go run new_task.go
```

## Message acknowledgment

Doing a task can take a few seconds, you may wonder what happens if a consumer starts a long task and it terminates before it completes. With our current code, once RabbitMQ delivers a message to the consumer, it immediately marks it for deletion. In this case, if you terminate a worker, the message it was just processing is lost. The messages that were dispatched to this particular worker but were not yet handled are also lost.

But we don't want to lose any tasks. If a worker dies, we'd like the task to be delivered to another worker.

In order to make sure a message is never lost, RabbitMQ supports `message acknowledgments (https://www.rabbitmq.com/docs/confirms)`. An ack(nowledgement) is sent back by the consumer to tell RabbitMQ that a particular message has been received, processed and that RabbitMQ is free to delete it.

If a consumer dies (its channel is closed, connection is closed, or TCP connection is lost) without sending an ack, RabbitMQ will understand that a message wasn't processed fully and will re-queue it. If there are other consumers online at the same time, it will then quickly redeliver it to another consumer. That way you can be sure that no message is lost, even if the workers occasionally die.


A timeout (30 minutes by default) is enforced on consumer delivery acknowledgement. This helps detect buggy (stuck) consumers that never acknowledge deliveries. You can increase this timeout as described in `Delivery Acknowledgement Timeout (https://www.rabbitmq.com/docs/consumers#acknowledgement-timeout)`.


In this tutorial we will use manual message acknowledgements by passing a false for the `"auto-ack"` argument and then send a proper acknowledgment from the worker with `d.Ack(false)` (this acknowledges a single delivery), once we're done with a task.


```go
msgs, err := ch.Consume(
  q.Name, // queue
  "",     // consumer
  false,  // auto-ack
  false,  // exclusive
  false,  // no-local
  false,  // no-wait
  nil,    // args
)
failOnError(err, "Failed to register a consumer")

var forever chan struct{}

go func() {
  for d := range msgs {
    log.Printf("Received a message: %s", d.Body)
    dotCount := bytes.Count(d.Body, []byte("."))
    t := time.Duration(dotCount)
    time.Sleep(t * time.Second)
    log.Printf("Done")
    d.Ack(false)
  }
}()

log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
<-forever
```

Using this code, you can ensure that even if you terminate a worker using CTRL+C while it was processing a message, nothing is lost. Soon after the worker terminates, all unacknowledged messages are redelivered.


Acknowledgement must be sent on the same channel that received the delivery. Attempts to acknowledge using a different channel will result in a channel-level protocol exception. See the `doc guide on confirmations (https://www.rabbitmq.com/docs/confirms)` to learn more.


- `Forgotten acknowledgment`

It's a common mistake to miss the `ack`. It's an easy error, but the consequences are serious. 

Messages will be redelivered when your client quits (which may look like random redelivery), but RabbitMQ will eat more and more memory as it won't be able to release any unacked messages.

In order to debug this kind of mistake you can use `rabbitmqctl` to print the `messages_unacknowledged` field:


```bash
sudo rabbitmqctl list_queues name messages_ready messages_unacknowledged
```

## Message durability

We have learned how to make sure that even if the consumer dies, the task isn't lost. But our tasks will still be lost if RabbitMQ server stops.

When RabbitMQ quits or crashes it will forget the queues and messages unless you tell it not to. Two things are required to make sure that messages aren't lost: we need to mark both the queue and messages as durable.

First, we need to make sure that the queue will survive a RabbitMQ node restart. In order to do so, we need to declare it as `durable`:


```go
q, err := ch.QueueDeclare(
  "hello",      // name
  true,         // durable
  false,        // delete when unused
  false,        // exclusive
  false,        // no-wait
  nil,          // arguments
)
failOnError(err, "Failed to declare a queue")
```

Although this command is correct by itself, it won't work in our present setup. That's because we've already defined a queue called `hello` which is not durable in producer code (send.go). 

RabbitMQ doesn't allow you to redefine an existing queue with different parameters and will return an error to any program that tries to do that. But there is a quick workaround - let's declare a queue with different name, for example `task_queue`:

`This durable option change needs to be applied to both the producer and consumer code.`

At this point we're sure that the task_queue queue won't be lost even if RabbitMQ restarts. Now we need to mark our messages as persistent

by using the `amqp.Persistent` option amqp.Publishing takes.

```go
err = ch.PublishWithContext(ctx,
  "",           // exchange
  q.Name,       // routing key
  false,        // mandatory
  false,
  amqp.Publishing {
    DeliveryMode: amqp.Persistent,
    ContentType:  "text/plain",
    Body:         []byte(body),
})
```

`Note :::` Marking messages as persistent doesn't fully guarantee that a message won't be lost. Although it tells RabbitMQ to save the message to disk, there is still a short time window when RabbitMQ has accepted a message and hasn't saved it yet. Also, RabbitMQ doesn't do `fsync(2)` for every message -- it may be just saved to cache and not really written to the disk. 

The persistence guarantees aren't strong, but it's more than enough for our simple task queue. If you need a stronger guarantee then you can use `publisher confirms (https://www.rabbitmq.com/docs/confirms)`.


## Round-Robin dispatching

One of the advantages of using a `Task Queue (Work Queue)` is the ability to easily parallelise work.

If we are building up a backlog of work, we can just add more workers and that way, scale easily.

First, let's try to run two `worker.go (consumer script)` scripts at the same time. They will both get messages from the queue, but how exactly? Let's see.

You need three consoles open. Two will run the worker.go script. These consoles will be our two consumers - C1 and C2.

![alt text](image-11.png)

```shell
# shell 1
go run worker.go
# => [*] Waiting for messages. To exit press CTRL+C
```

```shell
# shell 2
go run worker.go
# => [*] Waiting for messages. To exit press CTRL+C
```

In the third one we'll publish new tasks. Once you've started the consumers you can publish a few messages:

```shell
# shell 3
go run new_task.go First message.
go run new_task.go Second message..
go run new_task.go Third message...
go run new_task.go Fourth message....
go run new_task.go Fifth message.....
```

Let's see what is delivered to our workers:


```shell
# shell 1
go run worker.go
# => [*] Waiting for messages. To exit press CTRL+C
# => [x] Received 'First message.'
# => [x] Received 'Third message...'
# => [x] Received 'Fifth message.....'
```

```shell
# shell 2
go run worker.go
# => [*] Waiting for messages. To exit press CTRL+C
# => [x] Received 'Second message..'
# => [x] Received 'Fourth message....'
```

By default, RabbitMQ will send each message to the next consumer, in sequence. On average every consumer will get the same number of messages. This way of distributing messages is called round-robin. Try this out with three or more workers.


## Fair dispatch

You might have noticed that the dispatching still doesn't work exactly as we want. For example in a situation with two workers, when all odd messages are heavy and even messages are light, one worker will be constantly busy and the other one will do hardly any work. Well, RabbitMQ doesn't know anything about that and will still dispatch messages evenly.

This happens because RabbitMQ just dispatches a message when the message enters the queue. It doesn't look at the number of unacknowledged messages for a consumer. It just blindly dispatches every n-th message to the n-th consumer.

![alt text](image-12.png)


In order to defeat that we can set the `prefetch count with the value of 1`. This tells RabbitMQ not to give more than one message to a worker at a time. Or, in other words, don't dispatch a new message to a worker until it has processed and acknowledged the previous one. Instead, it will dispatch it to the next worker that is not still busy.


```go
err = ch.Qos(
  1,     // prefetch count
  0,     // prefetch size
  false, // global
)
failOnError(err, "Failed to set QoS")
```



# Kafka vs RabbitMQ

...


