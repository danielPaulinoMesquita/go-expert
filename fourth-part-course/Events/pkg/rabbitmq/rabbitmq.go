package rabbitmq

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

func OpenChannel() (*amqp.Channel, error) {
	// this channel is different of channel from MultiTreading
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		panic(err)
	}

	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}

	return ch, nil
}

func Consume(channel *amqp.Channel, out chan<- amqp.Delivery) error {
	// minhafila is the name of our queue created in RabbitMq
	msgs, err := channel.Consume(
		"minhafila",
		"go-consumer",
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		return err
	}

	for msg := range msgs {
		out <- msg
	}

	return nil

}

func Publish(channel *amqp.Channel, body string) error {
	// amq.drect is the name of exchange, it is one of the default names in the rabbitmq, you must add the queue with feature of binding
	err := channel.Publish(
		"amq.direct",
		"",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		},
	)

	if err != nil {
		return err
	}

	return nil
}
