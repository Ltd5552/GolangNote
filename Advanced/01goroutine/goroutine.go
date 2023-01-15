package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 100; i++ {
		go func(i int) {
			fmt.Println(i) //此时用的是函数外面的i，一次只
		}(i) //闭包函数的话
	}
	fmt.Println("main")
	time.Sleep(time.Second)
}

//goroutine什么时候结束?
//对应的函数执行结束了，groutine就结束了
//而main函数执行完了，所有goroutine都结束了
