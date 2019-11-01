package main 
import (
	"fmt"
	"runtime"
	// "time"
)

func say(s string) {
	for i:=0;i<3;i++ {
		runtime.Gosched()
		fmt.Println(s)
	}
}

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

func main() {

	a := 1
	b := 2
	defer calc("1", a, calc("10", a, b))
	a = 0
	defer calc("2", a, calc("20", a, b))
	b = 1


	// runtime.GOMAXPROCS(1)

	go say("world")
	say("hello")

	// time.Sleep(100 * time.Millisecond)



}