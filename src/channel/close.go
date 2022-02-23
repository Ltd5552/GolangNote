package main

func main() {
	ch1 := make(chan int, 2)
	ch1 <- 10
	ch1 <- 20
}
