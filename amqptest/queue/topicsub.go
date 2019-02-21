package queue

import (
	"amqptest/config"
	"fmt"
	"github.com/streadway/amqp"
)

func TopicConsumer(topickeys []string) {
	conf := config.InitAmqpConfig()

	conn, err := amqp.Dial(conf.Address)
	FailOnError(err, "链接队列失败")
	defer conn.Close()

	ch, err := conn.Channel()
	FailOnError(err, "打开通道失败")
	defer ch.Close()

	err = ch.ExchangeDeclare(exchange_topic, "topic", true, false, false, false, nil)
	FailOnError(err, "声明exchange失败")

	q, err := ch.QueueDeclare("", false, false, true, false, nil)
	FailOnError(err, "声明queue失败")

	for _, v := range topickeys {
		err = ch.QueueBind(q.Name, v, exchange_topic, false, nil)
		fmt.Println("监听主题：" + v)
	}

	msgs, err := ch.Consume(q.Name, "", true, false, false, false, nil)
	FailOnError(err, "注册消费者失败")

	forever := make(chan bool)

	go func() {
		for msg := range msgs {
			fmt.Println("接收到消息：" + string(msg.Body))
		}
	}()

	fmt.Printf("队列在监听%v类型的消息", topickeys)
	<-forever
}
