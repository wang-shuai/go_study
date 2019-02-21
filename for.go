package main

import (
	"fmt"
)

type student struct {
	name string
	age  int
}

func main() {

	stus := []student{
		student{"wang", 23},
		student{"chen", 22},
	}
	m := make(map[string]student)

	for _, s := range stus {
		s.age = s.age + 3
		m[s.name] = s
	}

	for k, v := range m {
		fmt.Println(k, ":", v)
	}

	fmt.Printf("%05d \n", 22)

	m1 := map[string]int{"one": 1, "two": 2, "three": 3, "four": 4}
	for k, v := range m1 {
		fmt.Println(k, v)
	}
}
