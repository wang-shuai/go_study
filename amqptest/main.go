package main

import (
	"amqptest/queue"
	//"github.com/streadway/amqp"
	"fmt"
	"os"
)

func main() {
	tt := fmt.Sprintf("%s", "test")
	fmt.Println(tt)

	//_, err := amqp.Dial("amqp://yixincapital:123456@192.168.145.5:5672/yx_mq")
	//if err != nil{
	//	fmt.Println(err)
	//}

	//queue.SendToMq("first go mqueue message")
	//queue.Receive()

	// dots := "."
	// for i:=0;i<10;i++{
	// 	queue.SendToMq_Worker("msg" + dots)
	// 	dots += "."
	// }

	//queue.Receive_Worker()

	// queue.Sub()
	// for i := 0; i < 10; i++ {
	// 	queue.Pub("pub" + string(i))
	// }


	topics := getSubTopic()
	queue.TopicProducer("msg",topics)
	//queue.TopicConsumer(topics)

}

func getSubTopic() []string {
	if len(os.Args)<2 || os.Args[1]==""{
	return []string{"nothing"}
}

		return os.Args[1:]

}