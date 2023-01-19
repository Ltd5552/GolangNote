package main

import "fmt"

type user struct {
	name     string
	password string
}

func main() {
	a := user{name: "Ltd", password: "1024"}
	b := user{"Ltd", "1024"}
	c := user{name: "Ltd"}
	c.password = "1024"
	var d user
	d.name = "Ltd"
	d.password = "1024"

	fmt.Println(a, b, c, d)                 // {Ltd 1024} {Ltd 1024} {Ltd 1024} {Ltd 1024}
	fmt.Println(checkPassword(a, "haha"))   // false
	fmt.Println(checkPassword2(&a, "haha")) // false

	m := make(map[string]*student)
	stus := []student{
		{name: "L", age: 11},
		{name: "T", age: 12},
		{name: "D", age: 13},
	}

	for _, stu := range stus {
		fmt.Println(&stu)
		// 这里指向的是stu的内存空间，而range是值拷贝，当遍历完了都指向了最后一个stu的地址
		m[stu.name] = &stu
	}
	for k, v := range m {
		fmt.Println(k, "=>", v.name)
	}
}

func checkPassword(u user, password string) bool {
	return u.password == password
}

func checkPassword2(u *user, password string) bool {
	return u.password == password
}

// MyInt 类型别名
type MyInt = int

// NewInt 自定义类型
type NewInt int

type student struct {
	name string
	age  int
}
