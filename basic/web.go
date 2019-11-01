package main

import (
	"fmt"
	"net/http"
	"strings"
	"log"
	"net/url"
)

func sayHelloName(w http.ResponseWriter,r *http.Request){
	r.ParseForm()

	for k,v :=range r.Form{
		fmt.Println("key:",k)
		fmt.Println("val:",strings.Join(v,""))
	}
	fmt.Fprint(w,"hello shine w")
}

func main(){
	v := url.Values{}
	v.Set("name","ava")
	v.Add("friend","jess")
	v.Add("friend","sarah")
	fmt.Println(v.Encode())
	fmt.Println(v.Get("friend"))
	fmt.Println(v["friend"])

	http.HandleFunc("/",sayHelloName)
	err:= http.ListenAndServe(":9090",nil)
	if err != nil{
		log.Fatal("listenAndServe:",err)
	}


}
