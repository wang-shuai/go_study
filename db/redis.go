package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/garyburd/redigo/redis"
)

var (
	Pool *redis.Pool
)

func init() {
	redisHost := ":10001"
	Pool = newPool(redisHost)
	close()
}

func newPool(server string) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,

		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", server)
			if err != nil {
				return nil, err
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}

func close() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)
	signal.Notify(c, syscall.SIGKILL)

	go func() {
		<-c
		Pool.Close()
		os.Exit(0)
	}()
}

type Human struct {
	Name    string   `json:"name"`
	Age     int      `json:"age"`
	Sex     int      `json:"sex"`
	Hobbies []string `json:"hobbies"`
	Desc    string   `json:"description"`
}

func Set(name string, man Human) error {
	conn := Pool.Get()
	defer conn.Close()

	strM, _ := json.Marshal(man)

	_, err := conn.Do("SET", name, strM)

	fmt.Println("set success:", string(strM))
	return err
}

func GetBytes(key string) ([]byte, error) {
	conn := Pool.Get()
	defer conn.Close()

	data, err := redis.Bytes(conn.Do("GET", key))
	if err != nil {
		return data, fmt.Errorf("error get key %s : %v", key, err)
	}
	return data, err
}
func GetString(key string) (string, error) {
	conn := Pool.Get()
	defer conn.Close()

	data, err := redis.String(conn.Do("GET", key))
	if err != nil {
		return data, fmt.Errorf("error get key %s : %v", key, err)
	}
	return data, err
}

func GetMutiple(keys ...interface{}) (interface{}, error) {
	conn := Pool.Get()
	defer conn.Close()
	// 只能取到一条数据。。
	//for _,key := range keys{
	//	conn.Send("get",key)
	//}
	//conn.Flush()
	//data,err := conn.Receive()
	//var h Human
	//fmt.Println("字符串结果",string(data.([]byte)))
	//if bt,ok := data.([]byte);ok {
	//	json.Unmarshal(bt,&h)
	//	fmt.Println("用户",h)
	//}

	fmt.Println(conn.Do("mget",keys...))
	replys, err := redis.Values(conn.Do("MGET", keys...))
	fmt.Println("replys",replys)

	var h Human
	for i,r := range replys{
		if bt,ok := r.([]byte);ok {
			json.Unmarshal(bt,&h)
			fmt.Println(i,"用户",h)
		}
	}

	return replys, err
}

func main() {
	// man := Human{"shine", 18, 1, []string{"badminton", "football"}, "des"}

	// Set("shine", man)
	// test, err := GetBytes("shine")
	// fmt.Println(test, err)
	// teststr, err1 := GetString("shine")
	// //fmt.Println(string(test), err)
	// fmt.Println(teststr, err1)

	// var m Human
	// err = json.Unmarshal(test, &m)
	// if err != nil {
	// 	fmt.Println("unmarshal err :", err)
	// } else {
	// 	fmt.Println(m)
	// }

	// err = json.Unmarshal([]byte(teststr), &m)
	// if err != nil {
	// 	fmt.Println("unmarshal err :", err)
	// } else {
	// 	fmt.Println(m)
	// }

	fmt.Println(".......................................")

	man1 := Human{"shine1", 18, 1, []string{"badminton", "football"}, "des"}
	man2 := Human{"shine2", 18, 1, []string{"badminton", "football"}, "des"}
	Set("shine1", man1)
	Set("shine2", man2)

	result,err:=GetMutiple("shine2","shine1")
	fmt.Println("result:",result,err)
}
