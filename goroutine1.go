package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)

	go func() {
		for m := range ch {
			fmt.Println("processed:", m)
		}
	}()

	ch <- "cmd.1"
	ch <- "cmd.2" //不一定能输出 如果没有后边的操作

	inch := make(chan int)
	outch := make(chan int)

	go func() {
		var in <-chan int = inch
		var out chan<- int
		var val int
		for {
			select {
			case out <- val:
				out = nil
				in = inch
			case val = <-in:
				out = outch
				in = inch
			}
		}
	}()

	go func() {
		for r := range outch {
			fmt.Println("result:", r)
		}
	}()

	time.Sleep(0)
	inch <- 1
	inch <- 2
	time.Sleep(3 * time.Second)
}
