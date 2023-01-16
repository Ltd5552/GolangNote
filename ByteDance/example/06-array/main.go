package main

import (
	"fmt"
	"reflect"
)

func main() {
	var array0 [3]int                 // 1.初始化为int类型的零值
	var array1 = [3]int{1, 2, 3}      //   指定初始值
	var array2 = [...]int{1, 2, 3, 4} // 2.让编译器根据值的个数自己推断
	array3 := [...]int{3: 2, 6: 5}    // 3.通过索引来初始化数组，长度则是最大索引，其余值为对应零值

	fmt.Println(array0) //[0 0 0]

	//Start1 := time.Now()
	if fmt.Sprintf("%T", array1) != fmt.Sprintf("%T", array2) {
		fmt.Println("类型不同")
		fmt.Println("array1:", fmt.Sprintf("%T", array1))
		fmt.Println("array2:", fmt.Sprintf("%T", array2))
	}
	//类型不同
	//array1: [3]int
	//array2: [4]int

	//Duration1 := float64(time.Since(Start1) / time.Millisecond)
	//fmt.Println(Duration1)

	//Start2 := time.Now()
	if reflect.TypeOf(array1) != reflect.TypeOf(array2) {
		fmt.Println("类型不同")
		fmt.Println("array1:", reflect.TypeOf(array1))
		fmt.Println("array2:", reflect.TypeOf(array2))
	}
	//Duration2 := float64(time.Since(Start2) / time.Millisecond)
	//fmt.Println(Duration2)
	fmt.Println(array3) // [0,0,0,2,0,0,5]
}
