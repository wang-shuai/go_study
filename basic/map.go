package main
import "fmt"

func main() {
	m1 := make(map[int]string)
	m1[1] = "first"
	m1[2] = "second"
	if v,ok := m1[3];ok {
		fmt.Println(v)
	}else{
		fmt.Println("key not found")
	}

	for k,v := range m1{
		fmt.Println(k,v)
	}
}