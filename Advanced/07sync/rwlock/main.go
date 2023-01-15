package main

import (
	"fmt"
	"sync"
	"time"
)

// rwLock  读写互斥锁
// 当读的操作数量远大于写操作的时候，读写互斥锁效率高
// 读的时候别人不需要等待，只有写的时候要等
var (
	x = 0
	//lock   sync.Mutex
	wg     sync.WaitGroup
	rwLock sync.RWMutex
)

func read() {
	defer wg.Done()
	rwLock.RLock()
	fmt.Println(x)
	time.Sleep(time.Millisecond)
	rwLock.RUnlock()
}

func write() {
	defer wg.Done()
	rwLock.Lock()
	x = x + 1
	time.Sleep(time.Millisecond * 5)
	rwLock.Unlock()
}
func main() {

	start := time.Now()
	for i := 0; i < 100; i++ {
		go write()
		wg.Add(1)
	}
	wg.Wait()
	for i := 0; i < 1000; i++ {
		go read()
		wg.Add(1)
	}
	wg.Wait()
	fmt.Println(time.Now().Sub(start))
}
