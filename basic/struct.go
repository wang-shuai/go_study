package main

import "fmt"

type person struct{
	name string
	age int
	phone string
}

func Older(p1,p2 person)(person,int){
	if(p1.age>p2.age){
		return p1,p1.age-p2.age
	}
	return p2,p2.age-p1.age
}

type Employee struct{
	person
	speciality string
	phone string
}

func main(){
	var tom person
	tom.age,tom.name=25,"tom"

	bob := person{"bob",18,"11111"}

	paul:= new (person);	 //person{name:"paul",age:41}
	paul.name,paul.age="paul",41;

	tb_older,tb_diff := Older(tom,bob)
	tp_older,tp_diff := Older(tom,*paul)
	bp_older,bp_diff := Older(bob,*paul)

	fmt.Printf("Of %s and %s , %s is older by %d years \n",tom.name,bob.name,tb_older.name,tb_diff)
	fmt.Printf("Of %s and %s , %s is older by %d years \n",tom.name,paul.name,tp_older.name,tp_diff)
	fmt.Printf("Of %s and %s , %s is older by %d years \n",paul.name,bob.name,bp_older.name,bp_diff)

	jack := Employee{person{"jack",28,"12345"},"computer","67890"}

	fmt.Println("jack's work phone is:", jack.phone)
	// 如果我们要访问Human的phone字段
	fmt.Println("jack's personal phone is:", jack.person.phone)
}