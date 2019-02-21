package queue

import "log"
import "github.com/streadway/amqp"
import (
	"amqptest/config"
	"bytes"
	"time"
)

// import "fmt"

func Receive() {
	amqpConfig := config.InitAmqpConfig()
	// fmt.Println(amqpConfig)
	conn, err := amqp.Dial(amqpConfig.Address)
	FailOnError(err, "链接队列失败")
	defer conn.Close()

	ch, err := conn.Channel()
	FailOnError(err, "打开通道失败")
	defer ch.Close()

	_, err = ch.QueueDeclare(
		amqpConfig.QueueName,
		amqpConfig.Durable,
		amqpConfig.AutoDelete,
		false,
		false,
		nil,
	)
	FailOnError(err, "声明队列失败")

	msgs, err := ch.Consume(
		amqpConfig.QueueName,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	FailOnError(err, "注册消费者失败")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("接收到消息：%s", d.Body)
		}
	}()

	log.Printf(" [*]消息等待中，停止请按Ctrl + c")
	<-forever
}

//Using message acknowledgments and prefetch count you can set up a work queue.
//The durability options let the tasks survive even if RabbitMQ is restarted.

func Receive_Worker() {
	amqpConfig := config.InitAmqpConfig()
	// fmt.Println(amqpConfig)
	conn, err := amqp.Dial(amqpConfig.Address)
	FailOnError(err, "链接队列失败")
	defer conn.Close()

	ch, err := conn.Channel()
	FailOnError(err, "打开通道失败")
	defer ch.Close()

	_, err = ch.QueueDeclare(
		amqpConfig.QueueName,
		amqpConfig.Durable,
		amqpConfig.AutoDelete,
		false,
		false,
		nil,
	)
	FailOnError(err, "声明队列失败")

	err = ch.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)

	msgs, err := ch.Consume(
		amqpConfig.QueueName,
		"",
		false, //设置为不是自动反馈消息成功
		false,
		false,
		false,
		nil,
	)
	FailOnError(err, "注册消费者失败")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("消费者接收到消息：%s", d.Body)
			dot_cnt := bytes.Count(d.Body, []byte("."))
			t := time.Duration(dot_cnt)
			time.Sleep(t * time.Second)
			log.Printf("done")
			d.Ack(false) // 这里消费成功之后 反馈已经成功-true 其他未消费记录 继续发到这个队列，false 其他队列一起处理未消费的队列
		}
	}()

	log.Printf(" [*]消息等待中，停止请按Ctrl + c")
	<-forever
}
