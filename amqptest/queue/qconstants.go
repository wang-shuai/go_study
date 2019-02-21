package queue

import (
	"log"
	"fmt"
)

var (
	exchange_fanout = "logs_test"
	exchange_topic  = "topic_test"
)

func FailOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s:%s", msg, err)
		panic(fmt.Sprintf("%s:%s", msg, err))
	}
}
