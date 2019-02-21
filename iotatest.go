package main

import (
	"fmt"
)

const (
	i = iota
	j = 3.14
	k = iota
	l
)

const (
	a = iota
	b = 3.14
	c
	d
)

func main() {
	fmt.Println(i, j, k, l) //0 3.14 2 3
	fmt.Println(a, b, c, d) //0 3.14 3.14 3.14
}
