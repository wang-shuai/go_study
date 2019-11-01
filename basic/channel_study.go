package main 
import "fmt"

func fibonacci1 (n int,c chan int){
	x,y := 1,1
	for i:=0;i<n;i++{
		c <- x
		x,y=y,x+y
	}
	close(c)
}

func list(c,quit chan int) {
	for{
		select{
		case i := <- c:
			fmt.Println(i)
		case <- quit:
			fmt.Println("quit")
			return
		}
	}
}


func fibonacci2(c,quit chan int) {
	x,y := 1,1
	for{
		select{
			case c <- x:
				x,y = y,x+y
			case <- quit:
				close(c)
				fmt.Println("quit")
				return
		}
	}
}

func main() {
	c := make(chan int, 10)
	go fibonacci1(cap(c),c)
	for i := range c{
		fmt.Println(i)
	}


	cc := make(chan int)
	quit := make(chan int)

	go func(){
		for i:=1;i<10;i++{
		  cc <- i
		}
		quit <- 0
		close(cc)
		close(quit)
	}()
	list(cc,quit)
	

	ch := make(chan int)
	qt := make(chan int)

	go func(){
		for i:=1;i<10;i++{
		  fmt.Println(<-ch)
		}
		qt <- 0
		close(qt)
	}()
	fibonacci2(ch,qt)
}