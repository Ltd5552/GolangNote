package main

import "fmt"

func add(a, b int) (c int) {
	//c在返回值那里声明的，可以看作局部变量，进行“裸”返回
	c = a + b
	return

	//若又声明了同名参数需显式返回
	//var c = a + b
	//return c
}

func add2(a, b int) int {
	return a + b
}

func add3(x, y int) (z int) {
	defer func() {
		//z = 4
		z += 100
		fmt.Println(z) // z = 104
	}()

	z = x + y
	fmt.Println(z) //3
	//return的内容执行完了赋值给z了再调用defer
	return z + 1
}

func exists(m map[string]string, k string) (v string, ok bool) {
	v, ok = m[k]
	return v, ok
}

func function(args ...int) {

}

func getError() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err.(string)) // 将 interface{} 转型为具体类型。
		}
	}()

	panic("panic error!")
}

func main() {
	res := add(1, 2)
	fmt.Println(res) // 3

	v, ok := exists(map[string]string{"a": "A"}, "a")
	fmt.Println(v, ok) // A True

	s := []int{1, 2, 3}
	function(s...)

	z := add3(1, 2)
	fmt.Println(z + 1) //105

	getError()
}
