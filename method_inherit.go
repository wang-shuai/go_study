package main
import "fmt"

type People struct{}

func (p *People) ShowA() {
	fmt.Println("showA")
	p.ShowB()
}
func (p *People) ShowB() {
	fmt.Println("showB")
}

type Student struct{
	*People
}
func(s Student)ShowB(){
	fmt.Println("Student showB")
}


type Teacher struct {
	People
}

func (t *Teacher) ShowB() {
	fmt.Println("teacher showB")
}

func main() {
	t := Teacher{}
	t.ShowA()
	t.ShowB()
	// t.People.ShowA()

	tt := new(Teacher)
	tt.ShowA()
	tt.ShowB()


	s := Student{}
	s.ShowA()
	s.ShowB()

	ss := new (Student)
	ss.ShowA()
	ss.ShowB()
}