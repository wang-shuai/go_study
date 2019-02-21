package main
import "fmt"

type Human struct{
	name string
	age int
}

type Student struct{
	Human
	score float64
}
type Emplyee struct{
	*Human // 指针 实现*T的method，不带*则不能实现interface有*T的method
	salary float64
}

type person interface{
	do()
	eat()
} 

type pSpecial interface{
	growUp()
}

func (h Human) eat(){
	fmt.Println(h.name + " is eating")
}
func (h *Human) growUp(){
	fmt.Println(h.name + " grow up")
	h.age++
}
func (h Human) do() {
	fmt.Println(h.name + " is doing")
}
func (s Student) do() {
	fmt.Println(s.name + " is studying")
}
func (e Emplyee) do() {
	fmt.Println(e.name +" is working")
}


func main() {
    
	human := new(Human)
	human.name = "human";
	human.age = 40;
	fmt.Println("--------human-------")
	human.eat()
	human.do()
	human.growUp()

fmt.Println("--------employee-------")
	emp := new(Human)
	emp.name = "emp";
	emp.age = 30;
	employee := Emplyee{emp,12} 
	employee.eat()
	employee.do()
	employee.growUp()

fmt.Println("--------student-------")
	stu := Student{Human{"stu",14},98.5}
	stu.eat()
	stu.do()
stu.growUp()

fmt.Println("--------person interface-------")

	var p person

	h := Human{"human",30}
	fmt.Println("------- human-->person ------------")
	p = h
	p.do()
	p.eat()
	//p.growUp()
	h.growUp()
	
	fmt.Println("----------student-->person ----------")
	p = stu
	p.do()
	p.eat()
	
	fmt.Println("-----------emplyee-->person-----------")
	p = employee
	p.do()
	p.eat()

	fmt.Println("\n###############################\n")

	var ps pSpecial

	fmt.Println("human ----------------------")
	ps = human
	ps.growUp()

	// 这里会报错 因为stu没有实现*T的接口 (因为*T的实现在human，stu内的匿名human不带*)
	// fmt.Println("student----------------------")
	// ps = stu
	// ps.growUp()
	
	fmt.Println("emplyee----------------------")
	ps = employee
	ps.growUp()
}







