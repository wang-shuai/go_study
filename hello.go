package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Printf("hello world\n")

	fmt.Println(SliceRandList(13, 30))
}

func SliceRandList(min, max int) []int {
	if max < min {
		min, max = max, min
	}
	length := max - min + 1
	t0 := time.Now()
	rand.Seed(int64(t0.Nanosecond()))
	list := rand.Perm(length)
	fmt.Println(list)
	for index := range list {
		list[index] += min
	}
	return list
}
