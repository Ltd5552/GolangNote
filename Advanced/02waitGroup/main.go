package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// waitGroup
//
//	func f() {
//		rand.Seed(time.Now().UnixNano()) //随机数种子
//		for i := 0; i < 5; i++ {
//			r1 := rand.Int()    //int64
//			r2 := rand.Intn(10) //0 <= x < 10
//			fmt.Println(r1, r2)
//		}
//	}
func f1(i int) {
	defer wg.Done() //当f1即将结束时一定会执行wg.Done()
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(500)))
	fmt.Println(i)
}

var wg sync.WaitGroup

func main() {
	//f()
	//wg.Add(10)
	for i := 0; i < 10; i++ {
		wg.Add(1) //计数
		go f1(i)
	} //如何知道这10个goroutine都执行结束了？

	wg.Wait() //等待wg的计数器减为0
}
