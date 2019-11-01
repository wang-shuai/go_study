package main

import (
	"fmt"
	"encoding/json"
)

// 要想序列化，属性必须要大写，跟json匹配使用 标签 json:"propertyname" 形式
type Entry struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Pay  Pay    `json:"json"`
}
type Pay struct {
	Msg string `json:"message"`
}

//type mid Pay
//
//func (lp *Pay) UnmarshalJSON(b []byte) error {
//	var s string
//	if err := json.Unmarshal(b, &s); err != nil {
//		return err
//	}
//	var f mid
//	if err := json.Unmarshal([]byte(s), &f); err != nil {
//		return err
//	}
//
//	*lp = Pay(f)
//
//	return nil
//}

func main() {

	pay := Pay{ "3"}
	en := Entry{"1","2",pay}
	fmt.Printf("%v",en)

	doc,_ := json.Marshal(en)
	fmt.Print(doc)

	//doc1 := []byte(`{
	//	"id":123,
	//	"name":"ttt",
	//	"pay":"{\"message\":\"test message\"}"
	//}`)

	var entry Entry
	if err := json.Unmarshal(doc, &entry); err != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("%v", entry)
}
