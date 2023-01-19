package main

import (
	"fmt"
	"strings"
)

func main() {
	a := "hello"
	fmt.Println(strings.Contains(a, "ll"))                // true  判断是否包含子串
	fmt.Println(strings.Count(a, "l"))                    // 2  统计子串出现次数
	fmt.Println(strings.HasPrefix(a, "he"))               // true  判断是否包含前缀
	fmt.Println(strings.HasSuffix(a, "llo"))              // true  判断是否包含后缀
	fmt.Println(strings.Index(a, "ll"))                   // 2	查找子串序列
	fmt.Println(strings.Join([]string{"he", "llo"}, "-")) // he-llo	字符串拼接
	fmt.Println(strings.Repeat(a, 2))                     // hellohello  字符串重复
	fmt.Println(strings.Replace(a, "e", "E", -1))         // hEllo	替换
	fmt.Println(strings.Split("a-b-c", "-"))              // [a b c]
	fmt.Println(strings.ToLower(a))                       // hello
	fmt.Println(strings.ToUpper(a))                       // HELLO
	fmt.Println(len(a))                                   // 5
	b := "你好"
	fmt.Println(len(b)) // 6
}
