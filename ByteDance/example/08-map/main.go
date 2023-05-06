package main

//func main() {
//	m := make(map[string]int)
//	m["one"] = 1
//	m["two"] = 2
//	fmt.Println(m)            // map[one:1 two:2]
//	fmt.Println(len(m))       // 2
//	fmt.Println(m["one"])     // 1
//	fmt.Println(m["unknown"]) // 0
//
//	//判断值是否存在
//	r, ok := m["unknown"]
//	if ok {
//		fmt.Println(r)
//	}
//
//	//在声明时就填充元素
//	m2 := map[string]int{"one": 1, "two": 2}
//	var m3 = map[string]int{"one": 1, "two": 2}
//	fmt.Println(m2, m3)
//
//	//指定容量为8
//	M := make(map[string]int, 8)
//
//	//遍历，顺序随机
//	for k, v := range M {
//		fmt.Println(k, v)
//	}
//
//	//删除键值对
//	delete(m, "one")
//}

func main() {
	PanicMap()
}

func PanicMap() {
	//1、unhashable type
	//a := make(map[interface{}]int)
	//b := []int{1}
	////可以绕过语法检查
	//a[b] = 1
	//	panic: runtime error: hash of unhashable type []int

	var a [2]int
	a[1] = 1
	//var a func(b int) int
	//a(1)
}
