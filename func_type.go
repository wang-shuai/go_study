package main
import "fmt"

type testInt func(int) bool

func isOdd(v int) bool {
	if v%2 == 0{
		return false
	}
	return true
}
func isEven(v int) bool {
	if v%2 ==0 {
		return true
	}
	return false
}

func filter(slice []int,f testInt) []int {
	var result []int
	for _,value := range slice{
		if f(value){
			result = append(result,value)
		}
	}
	return result
}

func main() {
	slice :=  []int{1,2,3,4,5,6,7,8,9}
	fmt.Println("slice = ",slice)
	odd := filter(slice,isOdd)
	fmt.Println("odd = ",odd)
	even := filter(slice,isEven)
	fmt.Println("even = ",even)
}