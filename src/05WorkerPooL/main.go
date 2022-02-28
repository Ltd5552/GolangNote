package main

import (
	"fmt"
	"math/rand"
)

//Aim
//使用goroutine和channel实现一个计算int64随机数各位数和的程序。
//开启一个goroutine循环生成int64类型的随机数，发送到jobChan
//开启24个goroutine从jobChan中取出随机数计算各位数的和，将结果发送到resultChan
//主goroutine从resultChan取出结果并打印到终端输出

//type Job struct {
//	// id
//	Id int
//	// 需要计算的随机数
//	RandNum int
//}
//
//type Result struct {
//	// 这里必须传对象实例
//	job *Job
//	// 求和
//	sum int
//}

func main() {
	//首先创建两个通道
	jobChan := make(chan int, 128)    //用于接收随机数
	resultChan := make(chan int, 128) //用于接收结果
	go func(re chan int) {
		for i := range re {
			fmt.Println(i)
		}
	}(resultChan)
	creatPool(64, jobChan, resultChan) //开启64个协程
	for true {
		x := rand.Intn(24)
		jobChan <- x
	}
}

//创建线程池
func creatPool(num int, jobChan <-chan int, result chan<- int) {
	for i := 0; i < num; i++ {
		go func(jobChan <-chan int, resultChan chan<- int) {
			for job := range jobChan {
				x := job * 2
				resultChan <- x //传入resultChan
			}
		}(jobChan, result) //单向通道
	}

}
