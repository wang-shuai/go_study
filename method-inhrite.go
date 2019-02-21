package main 

import "fmt"

type Human struct{
	name string
	age int
	phone string
}

type Student struct{
	Human
	school string
}

type Employee struct{
	Human
	company string
}

func(h *Human) SayHi(){
	fmt.Printf("hi,I am %s you can call me on %s \n",h.name,h.phone)
}

func main() {
	mark := Student{Human{"mark",20,"12345"},"school"}
	sam := Employee{Human{"sam",30,"67890"},"company"}

	mark.SayHi()
	sam.SayHi()
}
