package main

import (
	"fmt"
	"time"
)

func main() {

	a := 1
	switch a {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	case 4, 5:
		fmt.Println("four or five")
	default:
		fmt.Println("other")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

	switch2()
}

func switch2() {
	s := 1
	switch {
	case s == 1:
		fmt.Println("case1")
		fallthrough
	case s == 2:
		fmt.Println("case2")
	case s == 3:
		fmt.Println("case3")
	default:
		fmt.Println("...")
	}
}

func switch3() {
	//value1 := [...]int8{0, 1, 2, 3, 4, 5, 6}
	//switch 1 + 3 {
	//case value1[0], value1[1]:
	//	fmt.Println("0 or 1")
	//case value1[2], value1[3]:
	//	fmt.Println("2 or 3")
	//case value1[4], value1[5], value1[6]:
	//	fmt.Println("4 or 5 or 6")
	//}

	value2 := [...]int8{0, 1, 2, 3, 4, 5, 6}
	switch value2[4] {
	//这里虽然是无类型常量，但能够自动转为int8类型，所以编译可以通过
	case 0, 1:
		fmt.Println("0 or 1")
	case 2, 3:
		fmt.Println("2 or 3")
	case 4, 5, 6:
		fmt.Println("4 or 5 or 6")
	}
}

//func switch4() {
//	var x interface{}
//
//	switch i := x.(type) {
//	case nil:
//		printString("x is nil") // type of i is type of x (interface{})
//	case int:
//		printInt(i) // type of i is int
//	case float64:
//		printFloat64(i) // type of i is float64
//	case func(int) float64:
//		printFunction(i) // type of i is func(int) float64
//	case bool, string:
//		printString("type is bool or string") // type of i is type of x (interface{})
//	default:
//		printString("don't know the type") // type of i is type of x (interface{})
//	}
//}
