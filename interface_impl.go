package main

import (
	"fmt"
)

type People interface {
	Speak(string) string
}

type Stduent struct{}

func (stu *Stduent) Speak(think string) (talk string) {
	if think == "bitch" {
		talk = "You are a good boy"
	} else {
		talk = "hi"
	}
	return
}

func main() {
	
	think := "bitch"
	//// Stduent does not implement People (Speak method has pointer receiver)
	// var peo People = Stduent{}
	// fmt.Println(peo.Speak(think))

	var pp People = new(Stduent)
	fmt.Println(pp.Speak(think))
}