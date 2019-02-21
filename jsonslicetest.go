package main

import (
	"encoding/json"
	"fmt"
)

type Change struct {
	Mid     int      //菜单Id
	Actions []string //拥有的权限 "add"  "view"  "delete"  "update"
}
type Change_slice struct {
	ChgArr []Change //一个角色对应的菜单以及权限
}

func main() {

	//对象序列化为json字符串---------------------------------Begin
	var c1, c2 Change
	var msg Change_slice
	c1.Mid = 1
	c1.Actions = []string{"view", "add"}
	c2.Mid = 2
	c2.Actions = []string{"delete", "add", "update"}
	msg.ChgArr = []Change{c1, c2}
	fmt.Println(msg)
	b, err := json.Marshal(msg)
	if err == nil {
		fmt.Println(err)
	}
	fmt.Println("change_slice", string(b))

	var changes []Change
	changes = append(changes, c1)
	changes = append(changes, c2)
	b, err = json.Marshal(changes)
	if err == nil {
		fmt.Println(err)
	}
	fmt.Println("[]change:", string(b))
}
