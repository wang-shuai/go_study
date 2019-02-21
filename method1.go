package main

import "fmt"

type Rectangle struct{
	width,height float64
}

func area(r Rectangle) float64{
	return r.width * r.height
}

func main() {
	rec1 := Rectangle{12,3}
	rec2 := Rectangle{8,4}
	fmt.Println("area of rec1 is : ",area(rec1))
	fmt.Println("area of rec2 is : ",area(rec2))
}