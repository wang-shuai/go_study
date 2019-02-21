package main

//import "fmt"

import (
	"yixin_redis_cacheclear/common"

	"strconv"
)

func main() {
	//fmt.Println(config.RdsConfig)

	common.Test()
	var keys []string
	for i := 0; i < 10; i++ {
		keys = append(keys, "test"+strconv.Itoa(i))
	}
	common.Clear(keys)
}
