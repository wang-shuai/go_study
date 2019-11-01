package main

import "fmt"

func test(a []int, c chan int) {
	for _, v := range a {
		c <- v
	}
	close(c)
}

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 11, 12, 13, 14, 15, 16}
	c := make(chan int, 2)

	go test(a[:len(a)/2], c)
	go test(a[len(a)/2:], c)
	//go test(a,c)
	// for i := 0; i<len(a);i++{
	// 	v := <-c
	// 	fmt.Println(v)
	// }

	for v := range c {
		fmt.Println(v)
	}
}
