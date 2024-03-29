package main

import "fmt"

func main() {
	ch1 := make(chan int, 2)
	ch1 <- 10
	ch1 <- 20
	//close(ch1)

	//ch1 <- 2
	//关闭后可以读出，但无法继续写入
	//range 就读出去了
	for x := range ch1 {
		fmt.Println(x)
	}
	//a := <-ch1 //10出去
	//b := <-ch1 //20出去
	//fmt.Println(a, b)
	x, ok := <-ch1 //此时仍然可以执行，不会报错，第一个返回值是对应类型的零值，第二个返回值是false
	//当使用range方式的时候，它内置了判断ok的函数，当遇到false的时候会结束
	fmt.Println(x, ok)
}
