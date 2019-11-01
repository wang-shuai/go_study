package main
import "fmt"
import "strconv"

type Human struct{
	name string
	age int
	phone string
}

type Student struct{
	Human
	school string 
	loan float32
}

type Employee struct{
	Human
	company string
	money float32
}

func (h Human) sayhi(){
	fmt.Printf("hi,i am %s,%d years old,you can call me %s \n",h.name,h.age,h.phone)
}

func(h Human) sing(lyrics string){
	fmt.Println("lalalalalala~",lyrics)
}

func(h Human) String() string{
	return "{" + h.name + " - " + strconv.Itoa(h.age) + " years -  ✆ " + h.phone + "}"
}

func(e Employee) sayhi(){
	fmt.Printf("hi,i am %s,%d years old,work at %s,you can call me %s \n",e.name,e.age,e.company,e.phone)
}

type Men interface{
	sayhi()
	sing(lyrics string)
}

func main(){
	mike := Student{Human{"mike",20,"123123"},"university",1000}
	paul := Student{Human{"paul",16,"321321"},"high school",500}
	sam := Employee{Human{"sam",26,"5678"},"tttCompany",2000}
	tom := Employee{Human{"tom",29,"9876"},"gggCompany",3000}

	var i Men

	i=mike
	fmt.Println("this is mike,a student:")
	i.sayhi()
	i.sing("october eleventh")

	i = tom
	fmt.Println("this is tom,a employee:")
	i.sayhi()
	i.sing("november twelfth")

	x := make([]Men,3)
	x[0],x[1],x[2]=paul,sam,tom

	for _,m := range x{
		m.sayhi();
	}

	fmt.Println("------------interface parameter--------------")
	bob := Human{"bob",21,"12390"}
	fmt.Println("this human :", bob)
}