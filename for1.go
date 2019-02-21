package main
import "fmt"

func main() {
	for index := 10;index >0; index-- {
		if index == 5 {
			continue
		}
		fmt.Println(index)
	}


	s := make([]int, 5)
	s = append(s, 1, 2, 3)
	fmt.Println(s)

	sp := []int{5,6,7}
	s= append(s,sp...)
	fmt.Println(s)
}
