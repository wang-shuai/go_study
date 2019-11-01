package main

import "fmt"

type Rectangle struct{
	width int
	height int
	color string
}

func(rect Rectangle) area() int{
	return rect.width * rect.height
}

func(rect Rectangle) setColor(color string) {
	rect.color = color
}

func (rect *Rectangle) setColorWithRefer(color string) {
	rect.color = color
}

func main() {
	rect := Rectangle{6,8,"red"}
	fmt.Printf("rect area is : %d \n" , rect.area())
	rect.setColor("blue")
	fmt.Println("rect color is :" + rect.color)
	rect.setColorWithRefer("green")
	fmt.Println("rect color is :" + rect.color)// green,说明非指针也能以引用调用方法
	
	rect1 := new(Rectangle)
	rect1.width = 10
	rect1.height = 10
	rect1.color = "white"
	fmt.Printf("rect1 area is : %d \n" , rect1.area())
	rect1.setColor("blue")
	fmt.Println("rect1 color is :" + rect1.color) //white,说明指针调用非引用方法时 还是非引用形式
	rect1.setColorWithRefer("green")
	fmt.Println("rect1 color is :" + rect1.color)
}



