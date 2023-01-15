package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

//当使用自定义数据类型的时候还是需要用lock

var x int64
var wg sync.WaitGroup

//var lock sync.Mutex

func add() {
	defer wg.Done()

	//lock.Lock()
	//x = x + 1
	//lock.Unlock()

	atomic.AddInt64(&x, 1) //代替上面三步，且效率提高

}

func main() {
	wg.Add(100000)
	for i := 0; i < 100000; i++ {
		go add()
	}
	wg.Wait()
	fmt.Println(x)

	CAS()
}

func CAS() {
	x = 100
	ok := atomic.CompareAndSwapInt64(&x, 150, 200)
	//当old的100与x相同的时候返回true并将x改为new值
	//不同的时候返回false，x值不变
	fmt.Println(ok, x)
}
