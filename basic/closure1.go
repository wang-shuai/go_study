package main

import (
	"fmt"
	// "reflect"
	"runtime"
	"sync"
	"time"
)

func main() {
	runtime.GOMAXPROCS(1)
	wg := sync.WaitGroup{}
	wg.Add(20)

	for i := 0; i < 10; i++ {
		go func(k int) {
			fmt.Println("A:", k, "  ", i)
			wg.Done()
		}(i)
	}

	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println("B:", i)
			wg.Done()
		}(i)
	}

	wg.Wait()

	fmt.Println(`......`)
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println("C:", i)
		}(i)
	}

	time.Sleep(200)

	// funs := test()
	funs := test1()
	for _, f := range funs {
		f()
	}

	// defer func() {
	// 	if err := recover(); err != nil {
	// 		fmt.Println("++++")
	// 		f := err.(func() string)
	// 		fmt.Println(err, f(), reflect.TypeOf(err).Kind().String())
	// 	} else {
	// 		fmt.Println("fatal")
	// 	}
	// }()

	// defer func() {
	// 	panic(func() string {
	// 		return "defer panic"
	// 	})
	// }()
	// panic("panic")

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("fatal")
		}
	}()

	defer func() {
		panic("defer panic")
	}()
	panic("panic")

}

func test() []func() {
	var funs []func()
	for i := 0; i < 2; i++ {
		funs = append(funs, func() {
			println(&i, i)
		})
	}
	return funs
}
func test1() []func() {
	var funs []func()
	for i := 0; i < 2; i++ {
		x := i
		funs = append(funs, func() {
			println(&x, x)
		})
	}
	return funs
}
