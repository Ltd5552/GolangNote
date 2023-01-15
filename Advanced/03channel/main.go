package main

import (
	"fmt"
	"sync"
)

var b chan int //声明变量

//通道类型必须初始化
//引用类型

// make给三个固定的类型(slice/map/channel)进行初始化
var wg sync.WaitGroup

func noBuffChannel() {
	b = make(chan int) //不带缓存区通道初始化
	wg.Add(1)
	go func() {
		defer wg.Done()
		x := <-b
		fmt.Println("后台另一个goroutine从通道b中取到了", x) //一接收到10就开始打印
	}()
	b <- 10                     //这一步发生死锁 fatal error: all goroutines are asleep - deadlock!
	fmt.Println("发送到channel中了") //这里与另一句打印顺序不一定
	wg.Wait()
}
func BuffChannel() {
	b = make(chan int, 16) //带缓存区通道初始化
	b <- 10
	fmt.Println("发送到channel中了")
}
func main() {
	BuffChannel()
	noBuffChannel()
}
