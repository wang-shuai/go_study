package queue

import (
	"amqptest/config"
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

func Sub() {
	conf := config.InitAmqpConfig()

	conn, err := amqp.Dial(conf.Address)
	FailOnError(err, "链接队列失败")
	defer conn.Close()

	ch, err := conn.Channel()
	FailOnError(err, "创建channel失败")
	defer ch.Close()

	err = ch.ExchangeDeclare(exchange_fanout, "fanout", true, false, false, false, nil)
	FailOnError(err, "exchange声明失败")

	q, err := ch.QueueDeclare("", false, false, false, false, nil) // name为空，随机的临时队列
	FailOnError(err, "声明队列失败")

	err = ch.QueueBind(q.Name, "", exchange_fanout, false, nil)
	FailOnError(err, "绑定exchange失败")

	msgs, err := ch.Consume(q.Name, "", true, false, false, false, nil)
	FailOnError(err, "注册消费者程序失败")

	forever := make(chan bool)

	go func() {
		for msg := range msgs {
			log.Printf("收到消息%s", msg.Body)
		}
	}()

	fmt.Println("消息订阅成功")
	<-forever
}
