package main

import (
	"fmt"
	"os"
	"path/filepath"
	"io/ioutil"
	"log"
	"regexp"
)

func main(){
	os.Chdir("E:\\歌曲")
	path,_ := os.Getwd()
	walkDir(path)
}

func walkDir(path string){
	files,err:= ioutil.ReadDir(path)
	if err!=nil{
		log.Fatal(err)
	}

	for _,file := range files{
		if(file.IsDir()){
			walkDir(filepath.Join(path, file.Name()))
		}else{
			// fmt.Println(file.Name())
			rename(path,file.Name())
		}

	}
}

func rename(path,filename string){
	reg,_ := regexp.Compile("\\d+_")
	if reg.MatchString(filename){
		result := reg.ReplaceAllString(filename,"")
		fmt.Println(filename ,"--->",result)
		os.Rename(filepath.Join(path, filename),filepath.Join(path, result))
	}else{
		fmt.Println(filename)
	}
}