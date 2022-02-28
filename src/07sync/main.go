package main

import (
	"fmt"
	"sync"
)

//mutex互斥锁
var x = 0
var wg sync.WaitGroup
var lock sync.Mutex

//读写互斥锁：读的次数远远大于写的次数时使用
func add() {
	lock.Lock()
	for i := 0; i < 500000; i++ {
		//lock.Lock()
		x = x + 1
		//lock.Unlock()
	}
	lock.Unlock()
	wg.Done()
}
func main() {
	wg.Add(3)
	go add()
	go add()
	go add()
	//使用了公共空间导致同时写入的时候部分丢失了，少于150000
	//协程越多，重复写入可能性越大，差距越大
	wg.Wait()
	fmt.Println(x)
}
