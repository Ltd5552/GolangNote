package main

import (
	"fmt"
	"runtime"
	"sync"
)

func a() {
	for i := 0; i < 10; i++ {
		fmt.Printf("A:%d\n", i)
	}
}
func b() {
	for i := 0; i < 10; i++ {
		fmt.Printf("B:%d\n", i)
	}
}

var group sync.WaitGroup

func main() {
	runtime.GOMAXPROCS(2) //默认cpu的逻辑核心数，默认跑满，当只用一个核时相当于一个一个的来
	group.Add(2)
	go a()
	go b()
	group.Wait()
}
