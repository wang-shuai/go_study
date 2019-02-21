package common

import (
	"yixin_redis_cacheclear/config"
	"log"
	//"strings"
	"strconv"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"strings"
)

func Test() {

	conn := config.Pool.Get()
	defer conn.Close()

	//conn.Send("MULTI")
	//conn.Send("INCR", "foo")
	//conn.Send("INCR", "bar")
	//r, _ := conn.Do("EXEC")
	//fmt.Println(r) // prints [1, 1]

	for i := 0; i < 10; i++ {
		conn.Do("set", "test"+strconv.Itoa(i),i)
	}

	rs, err := redis.Strings(conn.Do("keys", "test*"))
	if err!=nil{
		fmt.Println(err)
	}
	for k,v := range rs {
		fmt.Println(k,v)
	}
}

func Clear(keysToClear []string) error {
	conn := config.Pool.Get()
	defer conn.Close()

	if len(keysToClear) > 0 {
		conn.Send("Multi")
		for _,key := range keysToClear{
		 	conn.Send("DEL",key)
		}

		n,err := conn.Do("exec")
		if  err != nil {
			log.Fatal(err)
			log.Println("del keys:" + strings.Join(keysToClear, " "))
			return err
		}
		log.Println(n)

		//var cnt int
		//for _,v := range keysToClear{
		//    cnt++
		//	if _,err :=	conn.Do("DEL",v);err != nil{
		//	   log.Fatal(err)
		//		cnt--
		//   }
		//}
		////fmt.Println(cnt)
		//log.Println(cnt)
	}
	return nil
}
