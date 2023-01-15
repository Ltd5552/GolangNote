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
