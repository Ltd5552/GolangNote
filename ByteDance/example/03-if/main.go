package main

import "fmt"

func main() {

	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}

	ifElse() //特殊用法

}
func getA() int {
	return 1
}

func ifElse() {
	b := 1
	if a := getA(); a == b {
		fmt.Println("=")
	}
	//不能在if-else外使用
	//fmt.Println(a)
}
