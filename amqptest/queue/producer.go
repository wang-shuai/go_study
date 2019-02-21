package queue

import "log"
import "github.com/streadway/amqp"
import "amqptest/config"
import "fmt"

func SendToMq(msg string) {
	amqpConfig := config.InitAmqpConfig()
	fmt.Println(amqpConfig.Address)
	conn, err := amqp.Dial(amqpConfig.Address)
	FailOnError(err, "链接队列失败")
	defer conn.Close()

	ch, err := conn.Channel()
	FailOnError(err, "打开通道失败")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		amqpConfig.QueueName,
		amqpConfig.Durable,
		amqpConfig.AutoDelete,
		false,
		false,
		nil,
	)
	FailOnError(err, "声明队列失败")

	err = ch.Publish(
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(msg),
		},
	)
	log.Printf(" [x] sent %s", msg)
	FailOnError(err, "发布消息失败")
}

func SendToMq_Worker(msg string) {
	amqpConfig := config.InitAmqpConfig()
	fmt.Println(amqpConfig.Address)
	conn, err := amqp.Dial(amqpConfig.Address)
	FailOnError(err, "链接队列失败")
	defer conn.Close()

	ch, err := conn.Channel()
	FailOnError(err, "打开通道失败")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		amqpConfig.QueueName,
		amqpConfig.Durable,
		amqpConfig.AutoDelete,
		false,
		false,
		nil,
	)
	FailOnError(err, "声明队列失败")

	err = ch.Publish(
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType:  "text/plain",
			Body:         []byte(msg),
			DeliveryMode: amqp.Persistent, // worker queue
		},
	)
	log.Printf(" [x] sent %s", msg)
	FailOnError(err, "发布消息失败")
}
