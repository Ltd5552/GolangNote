package main

import "fmt"

func main() {

	i := 1

	// 经典三段式
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// 省略初始语句或结束语句，但分号不能省略
	for ; i < 10; i++ {
		fmt.Println(i)
	}
	for i := 0; i < 10; {
		i++
		fmt.Println(i)
	}

	// 省略初始和结束语句，分号也省略，这里就像while循环了
	for i < 10 {
		fmt.Println(i)
	}

	// 全部省略，死循环
	for {
	}
}
