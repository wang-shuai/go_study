package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var token = make(chan int, 50)

func walkDir(dir string, fileSizes chan<- int64, wait *sync.WaitGroup) {
	defer wait.Done() //递归增加waitgroup的数量，这里每次退出一层递归，递减一次
	for _, entry := range dirEntries(dir) {
		if entry.IsDir() {
			wait.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			go walkDir(subdir, fileSizes, wait)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

func dirEntries(dir string) []os.FileInfo {
	token <- 1 //超过缓存数量 这里会阻塞 限制协程数量
	defer func() {
		<-token
	}()

	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "walker: %v \n", err)
	}
	return entries
}

func main() {
	var nfiles, nbytes int64 //文件数量、文件大小
	var wait sync.WaitGroup

	root := ""
	verbose := false
	t1 := time.Now()
	filesizes := make(chan int64)
	tick := make(<-chan time.Time)

	flag.StringVar(&root, "p", ".", "input dir")
	flag.BoolVar(&verbose, "v", false, "add verbose if you want")
	flag.Parse()

	if verbose {
		tick = time.Tick(500 * time.Millisecond)
	}

	wait.Add(1)
	go walkDir(root, filesizes, &wait)

	go func() {
		wait.Wait()
		close(filesizes)
	}()

loop:
	for {
		select {
		case <-tick:
			fmt.Fprintf(os.Stdout, "%v files, %f gb \n", nfiles, float64(nbytes)/1e9)
		case size, ok := <-filesizes:
			if !ok {
				break loop
			}
			nfiles++
			nbytes += size
		}
	}

	fmt.Fprintf(os.Stdout, "%v files, %f gb \n", nfiles, float64(nbytes)/1e9)
	fmt.Println(time.Since(t1))
}
