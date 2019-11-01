package main

import (
	"fmt"
	"reflect"
)

func main() {

	sn1 := struct {
		age  int
		name string
	}{age: 11, name: "qq"}

	sn2 := struct {
		age  int
		name string
	}{age: 11, name: "qq"}

	//进行结构体比较时候，只有相同类型的结构体才可以比较，结构体是否相同不但与属性类型个数有关，还与属性顺序相关。
	if sn1 == sn2 {
		fmt.Println("sn1 == sn2")
	}

	sn3 := struct {
		name string
		age  int
	}{age: 11, name: "qq"}

	// mismatched types
	// if sn1 == sn3 {
	// 	fmt.Println("sn1 == sn3")
	// }

	sm1 := struct {
		age int
		m   map[string]string
	}{age: 11, m: map[string]string{"a": "1"}}

	sm2 := struct {
		age int
		m   map[string]string
	}{age: 11, m: map[string]string{"a": "1"}}

	// 还有一点需要注意的是结构体是相同的，但是结构体属性中有不可以比较的类型，如map,slice。 如果该结构属性都是可以比较的，那么就可以使用“==”进行比较操作。
	// 可以使用reflect.DeepEqual进行比较

	// if sm1 == sm2 {
	// 	fmt.Println("sm1 == sm2")
	// }
	if reflect.DeepEqual(sm1, sm2) {
		fmt.Println("sm1 == sm2")
	}
}
