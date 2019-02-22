package main
import "fmt"
import "os"

var user = os.Getenv("USER")

func init(){
	if user == ""{
		panic("no value for $USER")
	}else {
		fmt.Println(user)	
	}
}

func makePanic() {
	panic("panic happen")
}

func throwsPanic(f func())(b bool){
	defer func(){
		if x:=recover();x!=nil{
			b=true
			fmt.Println(x)
		}
	}()
	f()
	return b
}

func main() {
	fmt.Println("makePanic ",throwsPanic(makePanic))
}