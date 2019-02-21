package main

import (
	"sync"
	"time"
)

type aliasMutex = sync.Mutex // alias 别名

type myMutext sync.Mutex // 这是个新的类型，不继承相关方法

func main(){
	var mtx  aliasMutex //	var mtx  myMutex 下边代码 编译不过
	mtx.Lock()
	mtx.Unlock()

	time.Now().Date()
}
