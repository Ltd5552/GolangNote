package main

import "sync"

//once
//Go语言中的sync包中提供了一个针对只执行一次场景的解决方案
//——sync.Once，sync.Once只有一个Do方法，其签名如下：
//func (o *Once) Do(f func())
//Do方法需要传入一个没有参数、没有返回值的变量（函数），
//可能需要搭配闭包来使用（匿名函数

type singleton struct{}

var instance *singleton
var once sync.Once

func GetInstance() *singleton {
	once.Do(func() {
		instance = &singleton{}
	})
	return instance
}

func main() {
	GetInstance()
}
