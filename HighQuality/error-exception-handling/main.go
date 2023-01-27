package main

import "fmt"

func doSomething() error {
	return nil
}

// Func1 warp errors, after golang 1.13
func Func1() error {
	if err := doSomething(); err != nil {
		return fmt.Errorf("test error warp: %w", err)
	}

	return nil
}
func main() {

}
