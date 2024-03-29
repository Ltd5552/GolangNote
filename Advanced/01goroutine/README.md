## goroutine与线程
```go
func main() {
	for i := 0; i < 100; i++ {
		go func(i int) {
			fmt.Println(i) //此时用的是函数外面的i，一次只
		}(i) //闭包函数的话
	}
	fmt.Println("main")
	time.Sleep(time.Second)
}
```
### 可增长的栈
OS线程（操作系统线程）一般都有固定的栈内存（通常为2MB）,一个goroutine的栈在其生命周期开始时只有很小的栈（典型情况下2KB），goroutine的栈不是固定的，他可以按需增大和缩小，goroutine的栈大小限制可以达到1GB，虽然极少会用到这么大。所以在Go语言中一次创建十万左右的goroutine也是可以的。
### goroutine调度机制
GPM是Go语言运行时（runtime）层面的实现，是go语言自己实现的一套调度系统。区别于操作系统调度OS线程。
- G goroutine 每个go关键字都会创建一个协程，里面有与所在P的绑定等信息
- M machine   是GO运行时对操作系统内核线程的虚拟，M与内核线程是一一映射的关系，G都要放到M上才能运行
- P processor 管理着一组G队列，并进行调度（比如把占用CPU时间较长的G暂停、运行后续的G等等）当自己的队列消费完了就去全局队列里取，如果全局队列里也消费完了会去其他P的队列里抢任务。

单从线程调度讲，Go语言相比起其他语言的优势在于OS线程是由OS内核来调度的，goroutine则是由Go运行时（runtime）自己的调度器调度的，这个调度器使用一个称为m:n调度的技术（复用/调度m个goroutine到n个OS线程）。
