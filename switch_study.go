package main
import (
"fmt"
"strconv"
)

type Element interface{}
type List []Element

type Person struct{
	name string
	age int
}

func(p Person) String() string {
	return "(name:" + p.name + " is " +strconv.Itoa(p.age) + " years"
}

func main() {
	list := make(List,3)
	list[0] = 1
	list[1] = "hello"
	list[2] = Person{"shine",31}

	for index,element := range list{
		if _,ok := element.(int);ok{
			fmt.Printf("list[%d] is int, value is %d \n", index,list[index])
		}else if _,ok := element.(string);ok{
			fmt.Printf("list[%d] is string, value is %s \n",index,list[index])
		}else if value,ok := element.(Person);ok {
			fmt.Printf("list[%d] is Person, value is %s \n" , index,value)
		}else{
			fmt.Printf("list[%d] is undefined \n" , index)
		}
	}

	for index,element := range list{
		switch value := element.(type){
		case int :
			fmt.Printf("list[%d] is int, value is %d \n", index,value)
		case string :
			fmt.Printf("list[%d] is string, value is %s \n", index,value)
		case Person:
			fmt.Printf("list[%d] is Person, value is %s \n", index,list[index])
		default:
			fmt.Printf("list[%d] is undefined \n" ,index)
		}
	}
}