package  main
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
	fmt.Printf("hi,i am %s,you can call me on %s \n",h.name,h.phone)
}
func(e *Employee) SayHi() {
	fmt.Printf("hi,i am %s, i am working in %s \n",e.name,e.company)
}

func main() {
	mark := Student{Human{"mark",23,"123456"},"6th high school"}
	lily := Employee{Human{"lily",32,"98765"},"yixincapital"}

	mark.SayHi();
	lily.SayHi();
}