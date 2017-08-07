package rabbitmq

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

var (
	qUser string = "**"
	qPass string = "***"
	qHost string = "***"
	qPort int    = 5672
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func InitMq() (ch *amqp.Channel, conn *amqp.Connection, err error) {
	rabbitmq := fmt.Sprintf("amqp://%s:%s@%s:%d/", qUser, qPass, qHost, qPort)
	conn, err = amqp.Dial(rabbitmq)
	//fmt.Println(conn, err)
	//defer conn.Close()
	ch, err = conn.Channel()
	//failOnError(err, "Failed to open a channel")
	//defer ch.Close()
	return
}

func SendText(body string) error {
	ch, conn, err := InitMq()
	failOnError(err, "Failed to open a channel")
	defer conn.Close()
	defer ch.Close()
	//body = "hello"
	err = ch.Publish(
		"appinfo",  // exchange
		"consumer", // routing key
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	log.Printf(" [x] Sent %s", body)
	failOnError(err, "Failed to publish a message")
	return err
}

func ReceiveText() {
	ch, conn, err := InitMq()
	failOnError(err, "Failed to open a channel")
	defer conn.Close()
	defer ch.Close()
	msgs, err := ch.Consume(
		"appConsumer", // queue
		"",            // consumer
		true,          // auto-ack
		false,         // exclusive
		false,         // no-local
		false,         // no-wait
		nil,           // args
	)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)
	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever

}
