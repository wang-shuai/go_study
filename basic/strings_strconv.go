package main

import (
	. "fmt"
	"strings"
	"strconv"
)

func main() {
	Println(strings.Contains("seafood","food"));
	
	s := []string{"foo","bar"}
	Println(strings.Join(s,","))

	Println(strings.Index("chicken","ken"))

	Println("ba" + strings.Repeat("na",2))

	Println(strings.Replace("oink oink oink","k","ky",2))
	Println(strings.Replace("oink oink oink","k","ky",-1))

	Printf("%v \n",strings.Split("a,b,c",","))

	Println(strings.Trim("!!!abc!!!","!"))
	Printf("fields are: %v \n" , strings.Fields("  foo  bar  "))

	ss := "abcd"
	tt := '王'
	str := make([]byte,0,100)
	str = strconv.AppendInt(str,4567,10)
	Println(string(str))
	str = strconv.AppendBool(str,false)
	Println(string(str))
	str = strconv.AppendQuote(str,ss)
	Println(string(str))
	str = strconv.AppendQuoteRune(str,tt)
	Println(string(str))

	a := strconv.FormatBool(false)
	b := strconv.FormatFloat(123.12,'g',12,64)
	c := strconv.FormatInt(1234,2)
	d := strconv.FormatUint(12345,10)
	e := strconv.Itoa(1023)
	Println(a,b,c,d,e)


	q, err := strconv.ParseBool("false")
	checkError(err)
	w, err := strconv.ParseFloat("123.21",64)
	checkError(err)
	r, err := strconv.ParseInt("1234",10,64)
	checkError(err)
	t, err := strconv.ParseUint("12345",10,64)
	checkError(err)
	y, err := strconv.Atoi("1023")
	checkError(err)
	Println(q,w,r,t,y)

}

func checkError(e error) {
	if e!=nil{
		Println(e)
	}
}









