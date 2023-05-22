package main

import (
	"errors"
	"fmt"
)

type user struct {
	name     string
	password string
}

func findUser(users []user, name string) (v *user, err error) {
	for _, u := range users {
		if u.name == name {
			return &u, nil
		}
	}
	return nil, errors.New("not found")
}

func main() {
	//u, err := findUser([]user{{"wang", "1024"}}, "wang")
	//if err != nil {
	//	//调用err.Error()方法
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println(u.name) // wang
	//
	//if u, err := findUser([]user{{"wang", "1024"}}, "li"); err != nil {
	//	fmt.Println(err) // not found
	//	return
	//} else {
	//	fmt.Println(u.name)
	//}
	Equal()
}

func Equal() {
	err1 := errors.New("err1")
	err2 := errors.New("err1")
	fmt.Printf("%T", err1)
	fmt.Printf("%T", err2)
	if err1 == err2 {
		fmt.Println("==")
	} else {
		fmt.Println("!=")
	}
}
