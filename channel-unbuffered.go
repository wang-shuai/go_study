package main
import "fmt"

func sum(a []int,c chan int){
	total := 0
	for _,i := range a{
		fmt.Println(i)
		total += i
	}

	c <- total
}

func main(){
	a := []int{1,2,3,4,5,6,7}
	c := make(chan int)
	
	go sum(a[:len(a)/2],c)
	go sum(a[len(a)/2:],c)
	x,y := <-c, <-c

	fmt.Println(x,y,x+y)
}