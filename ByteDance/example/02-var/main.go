package main

import (
	"fmt"
	"math"
)

const (
	i = 1
	j = 2
)

// 省略值表示与上面的相同
const (
	a = 7
	b // 7
	c // 7
)

func main() {

	var a = "initial"

	var b, c int = 1, 2

	var d = true

	var e float64

	f := float32(e)

	g := a + "foo"
	fmt.Println(a, b, c, d, e, f) // initial 1 2 true 0 0
	fmt.Println(g)                // initialapple

	const s string = "constant"
	const h = 500000000
	const i = 3e20 / h
	fmt.Println(s, h, i, math.Sin(h), math.Sin(i))
}
