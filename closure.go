package main

import (. "fmt")

func test(x int) (func(),func())  {
    return func() {
        Println(&x,x)
        x+=10
    }, func() {
        Println(&x,x)
    }
}

func main()  {
    a,b:=test(100)
    
    // 闭包 引用相同的变量值，相当于引用传递。所以 该例子中 a和b的执行顺序影响输出结果
    b() // 100 
    a() // 100

    // a() // 100 
    // b() // 110
}
