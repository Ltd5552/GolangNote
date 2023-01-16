package main

import "fmt"

func main() {

	s := make([]string, 3)
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("get:", s[2])   // c
	fmt.Println("len:", len(s)) // 3

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println(s) // [a b c d e f]

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println(c) // [a b c d e f]

	fmt.Println(s[2:5]) // [c d e]
	fmt.Println(s[:5])  // [a b c d e]
	fmt.Println(s[2:])  // [c d e f]

	good := []string{"g", "o", "o", "d"}
	fmt.Println(good) // [g o o d]

	//声明切片
	var s1 []int

	//初始化赋值
	s2 := []int{1, 2, 3}

	//使用make()创建，
	//make([]Type, size, cap)
	//区别长度和容量
	var s3 = make([]int, 2, 6)

	//基于数组切片得到s4
	array := [4]int{1, 2, 3, 4}
	s4 := array[1:3]

	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
	fmt.Println(s4)

	//将切片打散
	s4 = append(s4, s2...)
}
