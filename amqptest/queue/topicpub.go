package queue

import (
	"amqptest/config"
	"github.com/streadway/amqp"
)

func TopicProducer(msg string, topickeys []string) {
	conf := config.InitAmqpConfig()

	conn, err := amqp.Dial(conf.Address)
	FailOnError(err, "链接队列失败")
	defer conn.Close()

	ch, err := conn.Channel()
	FailOnError(err, "打开通道失败")
	defer ch.Close()

	err = ch.ExchangeDeclare(exchange_topic, "topic", true, false, false, false, nil)
	FailOnError(err, "声明exchange失败")

	for _, key := range topickeys {
		err = ch.Publish(
			exchange_topic,
			key,
			false,
			false,
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(msg),
			})
		FailOnError(err, "发布消息失败")
	}

}
