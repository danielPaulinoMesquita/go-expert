package main

import (
	"fmt"
	"github.com/daniel/fcutils/pkg/rabbitmq"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	msgs := make(chan amqp.Delivery)

	go rabbitmq.Consume(ch, msgs)

	for msg := range msgs {
		fmt.Println(string(msg.Body))
		err := msg.Ack(false)
		if err != nil {
			return
		} // this work for no put in queue again
	}

}
