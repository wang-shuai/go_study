package main

import "fmt"
import "encoding/base64"
import "strconv"
import "bufio"
import "os"

func main()  {

	fmt.Println("输入祝词计算红包大小，输入-1退出")
	for{
		input := bufio.NewScanner(os.Stdin)
		input.Scan()
		if input.Text() == "-1" { 
			break
		}
		encStr := base64.StdEncoding.EncodeToString([]byte(input.Text()))
		fmt.Println(encStr)
		for _, b := range encStr {
			e := strconv.FormatInt(int64(b), 8)
			fmt.Println(b,e)
		}
	}
	return
}