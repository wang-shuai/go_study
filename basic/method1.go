package main

import "fmt"

type Rectangle struct{
	width,height float64
}

func area(r Rectangle) float64{
	return r.width * r.height
}

func(rec *Rectangle) resize() {
	rec.width = 222;
	rec.height = 222;
}
func(rec Rectangle) resize1() {
	rec.width = 111;
	rec.height = 111;
}
func main() {
	rec1 := Rectangle{12,3}
	rec2 := Rectangle{8,4}
	fmt.Println("area of rec1 is : ",area(rec1))
	fmt.Println("area of rec2 is : ",area(rec2))

	rec1.resize()
	fmt.Println("area of rec1 is : ",area(rec1))
	rec2.resize1()
	fmt.Println("area of rec2 is : ",area(rec2))
}