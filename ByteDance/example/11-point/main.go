package main

//func add2(n int) {
//	n += 2
//}
//
//func add2ptr(n *int) {
//	*n += 2
//}
//func main() {
//	n := 5
//	add2(n)
//	fmt.Println(n) // 5
//	add2ptr(&n)
//	fmt.Println(n) // 7
//}
//
//type Point struct {
//}
//
//func (p *Point) use() {
//
//}

type Dog struct {
	name string
}

func New(name string) Dog {
	return Dog{name}
}

func (dog *Dog) SetName(name string) {
	dog.name = name
}

func (dog Dog) Name() string {
	return dog.name
}

func main() {
	// 示例1。
	//New("little pig").SetName("monster") // 不能调用不可寻址的值的指针方法。
	dog := New("little pig")
	dog.SetName("monster") // 使用临时变量赋值变为可寻址的值后才能调用指针方法。

	// 示例2。
	map[string]int{"ltd": 0}["ltd"]++
	m := map[string]int{"ltd": 0}
	m["ltd"]++
	for i := range m {
		println(i) //i是字典索引结果，打印结果是ltd
	}

}
