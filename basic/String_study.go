package main 
import (
	"fmt"
	"strconv"
)

type Human struct{
	name string
	age int
	phone string
}

func (h Human)String() string{
	return "<"+h.name+"-" + strconv.Itoa(h.age) + "-" + h.phone + ">"
}

func main() {
	human := Human{"shine",31,"18611662004"}
	fmt.Println("the man is ",human);
}