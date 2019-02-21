package queue

import (
	"amqptest/config"
	"fmt"
	"github.com/streadway/amqp"
	// "log"
)

func Pub(msg string) {

	amqpconfig := config.InitAmqpConfig()

	conn, err := amqp.Dial(amqpconfig.Address)
	FailOnError(err, "链接队列失败")
	defer conn.Close()

	ch, err := conn.Channel()
	FailOnError(err, "创建channel失败")
	defer ch.Close()

	err = ch.ExchangeDeclare(exchange_fanout, "fanout", true, false, false, false, nil)
	FailOnError(err, "声明exchange失败")

	err = ch.Publish(exchange_fanout, "", false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(msg),
	})
	FailOnError(err, "发布消息失败")

	fmt.Println("发布消息：" + msg)
}
