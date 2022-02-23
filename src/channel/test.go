package main

import (
	"fmt"
	"sync"
)

//Aim:
//1.启动一个goroutine，生成100个数发送到ch1中
//2.启动一个goroutine，从ch1中取值，计算其平方并放到ch2中
//3.在main中从ch2取值打印出来

var w sync.WaitGroup

func f1(ch1 chan int) {
	defer w.Done()
	for i := 0; i < 100; i++ {
		ch1 <- i
	}
	close(ch1)
}
func f2(ch1, ch2 chan int) {
	defer w.Done()
	for x := range ch1 {
		ch2 <- x * x
	}
	close(ch2)
}

func main() {
	a := make(chan int, 100)
	b := make(chan int, 100)
	w.Add(2)
	go f1(a)
	go f2(a, b)
	w.Wait()
	for ret := range b {
		fmt.Println(ret)
	}

}
