package main

import (
	"fmt"
	"time"
)

type Message struct {
	id  int
	Msg string
}

func main() {
	c1 := make(chan Message)
	c2 := make(chan Message)

	// RabbitMQ
	go func() {
		time.Sleep(time.Second * 2)
		msg := Message{1, "Hello from RABBIT mq"}
		c1 <- msg
	}()

	// Kafka
	go func() {
		time.Sleep(time.Second * 1)
		msg2 := Message{2, "Hello from KAFKA"}
		c2 <- msg2
	}()

	// This select will choose the thread or channel faster
	for {
		select {
		case msg1 := <-c1: // rabbit mq
			fmt.Printf("received FROM RABBITMQ %d :%s \n", msg1.id, msg1.Msg)

		case msg2 := <-c2: // kafka
			fmt.Printf("received FROM KAFKA %d :%s \n", msg2.id, msg2.Msg)

		case <-time.After(time.Second * 3):
			println("TIME OUT !!!")
		}
	}

}
