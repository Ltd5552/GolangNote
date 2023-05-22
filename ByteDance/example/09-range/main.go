package main

import "fmt"

func main() {
	//nums := []int{2, 3, 4}
	//sum := 0
	//for i, num := range nums {
	//	sum += num
	//	if num == 2 {
	//		fmt.Println("index:", i, "num:", num) // index: 0 num: 2
	//	}
	//}
	//fmt.Println(sum) // 9
	//
	//m := map[string]string{"a": "A", "b": "B"}
	//for k, v := range m {
	//	fmt.Println(k, v) // b 8; a A
	//}
	//for k := range m {
	//	fmt.Println("key", k) // key a; key b
	//}
	//
	ForRangeArray()
	ForRangeSlice()
}

// ForRangeArray 数组是值类型，这里的e完全是复制的副本，在开始range表达式求完值就固定了，是静态的
func ForRangeArray() {
	numbers1 := [...]int{1, 2, 3, 4, 5, 6}
	maxIndex1 := len(numbers1) - 1
	for i, e := range numbers1 {
		if i == maxIndex1 {
			numbers1[0] += e
		} else {
			numbers1[i+1] += e
		}
	}
	fmt.Println(numbers1)
	//[7 3 5 7 9 11]
}

// ForRangeSlice 切片是引用类型，即使是复制的副本对其修改也会影响到原来的切片，所以是动态变化的
func ForRangeSlice() {
	numbers2 := []int{1, 2, 3, 4, 5, 6}
	maxIndex2 := len(numbers2) - 1
	for i, e := range numbers2 {
		if i == maxIndex2 {
			numbers2[0] += e
		} else {
			numbers2[i+1] += e
		}
	}
	fmt.Println(numbers2)
	//[22 3 6 10 15 21]
}
