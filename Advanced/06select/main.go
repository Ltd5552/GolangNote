package main

// 选择select
func main() {
	c := make(chan int)
	select {
	case a, ok := <-c:

	}
}
