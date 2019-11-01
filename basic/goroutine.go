package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		runtime.Gosched()
		fmt.Println(s, i)
	}
}
func say2(s string, c chan int) {
	for i := 0; i < 5; i++ {
		fmt.Println(s, i)
	}
	c<-0
	close(c)
}

func main() {
	// c := make(chan int)
	// go say2("world ",c)
	// say("hello main")
	// fmt.Println(<-c)

	var wg sync.WaitGroup
	done := make(chan struct{})
	workerCount := 2

	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			doit(i, done)
		}(i)
	}

	close(done)
	wg.Wait()
	fmt.Println("all done!")

	// for i := 0; i < workerCount; i++ {
	// 	go func(i int) {
	// 		do(i)
	// 	}(i)
	// }
	// fmt.Println("all done11111!")

}

func doit(workerId int, done <-chan struct{}) {
	fmt.Printf("[%v] is running\n", workerId)
	fmt.Println(<-done)
	fmt.Printf("[%v] is done\n", workerId)
}

func do(workerId int) {
	fmt.Printf("[%v] is running\n", workerId)
	time.Sleep(3 * time.Second)
	fmt.Printf("[%v] is done\n", workerId)
}
