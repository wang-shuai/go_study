package main

import "fmt"

type Human struct {
	name  string
	age   int
	phone string
}

type Student struct {
	Human
	school string
}

type Employee struct {
	Human
	company string
}

type Person interface {
	Speak()
}

type Pps interface {
	Talk()
}

func (h Human) Speak() {
	fmt.Printf("%s is speaking \n", h.name)
}

func (h *Human) Talk() {
	fmt.Printf("%s is talking \n", h.name)
}

func (h *Human) SayHi() {
	fmt.Printf("hi,I am %s you can call me on %s \n", h.name, h.phone)
	h.SayHello()
}

func (h *Human) SayHello() {
	fmt.Printf("human,nice to meet you! \n")
}

func (s *Student) SayHi() {
	fmt.Printf("hi,I am %s,study at %s, you can call me on %s \n", s.name, s.school, s.phone)
}

func (e *Employee) SayHello() {
	fmt.Printf("employee,hello \n")
}

// func (e *Employee) SayHi() {
// 	fmt.Printf("hi,I am %s,work at %s, you can call me on %s \n", e.name, e.company, e.phone)
// }

func main() {
	mark := Student{Human{"mark", 20, "12345"}, "school"}
	sam := Employee{Human{"sam", 30, "67890"}, "company"}

	mark.SayHi()
	sam.SayHi()

	var p Person = Human{"test", 32, "212"}
	p.Speak()

	var t Pps = &Human{"tt", 32, "212"}
	t.Talk()
}
