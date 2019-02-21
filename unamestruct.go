package  main
import "fmt"

type Human struct{
	name string
	age int
	phone string
}
type Emplyee struct{
	Human
	speciality string
	phone string
}

func main() {
	shine := Emplyee{Human{"shine",30,"18611662004"},"computer sience","15901032363"}
	fmt.Println("shine's personal phone is : " + shine.Human.phone)
	fmt.Println("shine's other phone is : " + shine.phone)
}