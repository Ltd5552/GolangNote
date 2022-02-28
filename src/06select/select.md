##select
- select的使用类似于switch语句，它有一系列case分支和一个默认的分支；
- select不同于switch的顺序执行判断，是随机的
- 每个case会对应一个通道的通信（接收或发送）过程。
- select会一直等待，直到某个case的通信操作完成时，就会执行case分支对应的语句。
```go
select {
case <-chan1:
// 如果chan1成功读到数据，则进行该case处理语句
case chan2 <- 1:
// 如果成功向chan2写入数据，则进行该case处理语句
default:
// 如果上面都没有成功，则进入default处理流程
}
```
- select可以同时监听一个或多个channel，直到其中一个channel ready
- 如果多个channel同时ready，则随机选择一个执行



- 可以用来判断channel是否存满
```go
func write(ch chan string) {
   for {
      select {
      // 写数据
      case ch <- "hello":
         fmt.Println("write hello")
      default:
         fmt.Println("channel full")
      }
      time.Sleep(time.Millisecond * 500)
   }
}
```