#  走进Go语言基础

## 1简介

1. 高性能、高并发
2. 语法简单、学习曲线平缓
3. 丰富的标准库
4. 完善的工具链
5. 静态链接
6. 快速编译
7. 跨平台
8. 垃圾回收

```go
package main

import(
	"net/http"
)

func main(){
	http.Handle("/", httpFileServer(http.Dir(".")))
	http.ListenAndServer(":8080", nil)
}

// 一段可以承载静态文件访问的支持高并发、高性能的服务端代码
```



## 2入门-基础语法

- Hello World

  运行`go run example`

  编译`go build example`

### 2.1变量和常量
- 变量

  - 强类型语言，每个变量都有自己的变量类型

  - 字符串是内置类型，可直接通过加号拼接，也能直接用等号比较两个字符串（和python一样）

  - 变量声明的两种方式

    - 通过`var`，一般会自动推导变量类型，也可以显式写出类型

      ```go
      var a = "Ltd"
      var b, c int = 5, 2  // 可以多个赋值，语法糖
      var d float64  // 可以不赋值，只声明，默认空值
      ```

    - 通过`:=`，在函数内部声明并初始化，自动根据值推导类型

      ```go
      a := "Ltd"
      ```

  
  - 特别注意：
  
    - 函数外的每个语句都必须以关键字开始（const, func, import, type, var）
  
      ![image-20230115105826242](https://images.ltd7.ltd/img/studygo/tip1.png)
  
    - `:=`不能在函数外使用，注意同一作用域下同名变量不能重复使用，会报错
  
    - `_`常用于占位，表示忽略该返回值
  
      

- 常量

  - 和var声明一样，但常量在定义时必须赋值

    ```go
    const (
    	i = 1
    	j = 2
    )
    
    // 省略值表示与上面的相同
    const (
        a = 7
        b      // 7
        c      // 7
    )
    ```

  - go语言中的常量没有确定的类型，会根据使用的上下文来自动确定类型



### 2.2流程控制

- if else
  - 注意格式
    - 没有小括号
    - 注意大括号位置，换行
  

```go
if a == b {
	fmt.Println("=")
}
```

特殊写法：在if表达式之前再加一个执行语句

```go
// 这里getA()返回一个值赋值给a，再与b比较
if a := getA(); a == b {
	fmt.Println("=")
}
```

与正常的区别在于这个返回值的作用范围限制在了if-else中

![image-20230115134838354](https://images.ltd7.ltd/img/studygo/ifelse.png)



- 循环

  只有for循环，在经典的C三段式for循环下进行了增强，每一段都可以省略，支持使用`break`和`continue`跳出和继续循环

  ```go
  // 经典三段式
  for i := 0; i < 10; i++ {
      fmt.Println(i)
  }
  
  // 省略初始语句或结束语句，但分号不能省略
  for ; i < 10; i++{
      fmt.Println(i)
  } 
  for i := 0 ;i < 10;{
      i++
      fmt.Println(i)
  }
  
  // 省略初始和结束语句，分号也省略，这里就像while循环了
  for i < 10{
      fmt.Println(i)
  } 
  
  // 全部省略，死循环
  for {
  }
  ```



- switch
  - 与C或C++主要区别在于，go语言不需要在case中加break即可跳出，而C或C++会继续往下跑完所有case
  - 可以使用任意的变量类型，甚至可以取代任意的if-else语句，使得代码逻辑更加清晰
  - switch的一个分支可以有多个值，但只有一个default分支
  - `fallthrough`关键字可以执行满足条件case的下一个case（这个case跳过判断语句，直接执行），且可以叠加

```go
func switch1() {
    a := 1
    switch a {
    case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	default:
		fmt.Println("other")
	}
}

func switch2() {
	s := 1
	switch {
	case s == 1:
		fmt.Println("case1")
		fallthrough
	case s == 2:
		fmt.Println("case2")
	case s == 3:
		fmt.Println("case3")
	default:
		fmt.Println("...")
	}
}

// 此时结果会是:
// case1
// case2
```



### 2.3数组

数组是一个具有编号且长度固定的元素序列，但在实际应用场景下更多的是用切片（slice）

- 数组的长度必须是常量且不变，长度属于数组类型一部分，不同长度的int数组类型不同（不像python有list、set这种）
- 数组越界是panic
- 数组是值类型
- 支持 “==”、“!=” 操作符

```go
var array0 [3]int                 // 1.初始化为int类型的零值
var array1 = [3]int{1, 2, 3}      //   指定初始值
var array2 = [...]int{1, 2, 3, 4} // 2.让编译器根据值的个数自己推断，多维数组只有第一层可以这样用
array3 := [...]int{3: 2, 6: 5}    // 3.通过索引来初始化数组，长度则是最大索引，其余值为对应零值

fmt.Println(array0) //[0 0 0]

if fmt.Sprintf("%T", array1) != fmt.Sprintf("%T", array2) {
    fmt.Println("类型不同")
    fmt.Println("array1:", fmt.Sprintf("%T", array1))
    fmt.Println("array2:", fmt.Sprintf("%T", array2))
}
//类型不同
//array1: [3]int
//array2: [4]int

fmt.Println(array3) // [0,0,0,2,0,0,5]
```

注意声明的几种方式



### 2.4slice

由于数组的局限性（长度固定且与类型绑定），所以产生了切片类型

切片是一个有相同类型元素的可以任意更改长度的序列，是**基于数组的封装**，也有更多丰富的操作

可以用make来创建一个切片，可以像数组一样取值，可以使用append来追加元素，一般用于**快速地操作一块数据集合**

- 切片是引用类型，但自身是结构体，是值传递

- 注意使用append的时候必须把返回值赋给原数组
  - slice底层原理是存储了**长度、容量加一个指向底层数组的指针**，如果容量不够的话会扩容并返回新的slice
- 可以像python一样的切片，但不支持负索引

#### 创建方式

```go
//声明切片
var s1 []int

//初始化赋值
s2 := []int{1, 2, 3}

//使用make()创建
//make([]Type, size, cap)
//区别长度和容量
var s3 = make([]int, 2, 6)

//基于数组切片得到s4
array := [4]int{1, 2, 3, 4}
s4 := array[1:3]
```

长度是这个切片包含的元素个数，容量是这个切片完全扩容后的大小（长度+底层数组剩下的空间）

`len()`求长度   `cap()`求容量    判断空切片使用 `len(a) == 0`

- 一个长度和容量都为0的切片不一定是nil，一个nil值的切片只能说明没有底层数组
  - `s1 := []int{}`  s1 != nil
  - `s2 := make(int[], 0)`  s2 != nil
- 扩容策略是按照1，2，4，8成倍自动扩容

#### 赋值拷贝添加

直接赋值的话是共享底层数组`s1 := s2`

使用内置`copy()`才是分配新的地址空间保存 `copy(newslice, oldslice)`

```
s1 := []int{1, 2, 3}
```



#### 删除元素

目标元素索引是Index

`a = append(a[:index], a[index+1:]...)`



### 2.5map

map是一种完全无序、基于k-v的数据结构，引用类型，使用make初始化后才能使用

```go
m := make(map[string]int)
```

- 当索引不存在的键的时候值为对应类型的空值

  ```go
  m := make(map[string]int)
  fmt.Println(m["unknown"]) // 0
  ```

- 使用`range`遍历键值对

  ```go
  //遍历，顺序随机
  for k, v := range M {
  fmt.Println(k, v)
  }
  ```

- 使用`delete()`删除键值对

  ```go
  //删除键值对
  delete(m, "one")
  ```



### 2.6函数

#### 特点

- 支持不定参数
- 支持多返回值
- 支持定义返回值名称
- 支持匿名和闭包
- 函数可以赋值给变量，是一种类型

- 无重载
- 无默认参数

#### 参数

**引用传递与值传递**

golang中map、slice、chan、指针、interface默认以引用的方式传递——赋值地址

**不定参数**

使用`...`

```go
func function(args ...int) {    //0个或多个参数
}
```

而args是一个切片类型，可通过切片的操作进行参数访问

当使用slice作为变参的时候需要使用`...`将slice展开

```go
s := []int{1, 2, 3}
function(s...)
```
**任意类型的不定参数**

使用interface{}传递任意类型数据

```go
func function(args ...interface{}) {
}
```



#### 返回值

- 没有参数的return返回变量当前值

  ```go
  func add(a int, b int) (c int) {
  	//c在返回值那里声明的，“裸”返回
  	c = a + b
  	return
  }
  ```

- 返回值只能一个个接收，不能使用容器，但可以通过`_`忽略

返回顺序

```go
func add3(x, y int) (z int) {
	defer func() {
		//z = 4
		z += 100
		fmt.Println(z) // z = 104
	}()

	z = x + y
	fmt.Println(z) //3
	//return的内容执行完了赋值给z了再调用defer
	return z + 1
}

z := add3(1, 2)
fmt.Println(z + 1) //105
```



#### 匿名函数

没有名字的函数，可以直接像普通变量一样传递使用

#### 延时调用

**defer**，在return之前才执行，常用来做资源清理

- 关闭文件句柄

- 释放锁

- 释放数据库连接

- 同步内存和磁盘（zap）

  `defer logger.Sync()`

多个defer会形成栈，先定义的后执行

#### 异常处理

`defer()`和`recover()`

处理异常后，逻辑恢复到defer后的那个点

recover只能在defer函数那一层中直接使用，其他方式都是返回空值



### 2.7指针

相比于C/CPP，Go的指针操作很简单，主要用途就是对于传入参数进行修改，不能进行偏移和运算

`&`取地址， `*`取值

go中的值类型都有对应的指针类型



### 2.8结构体

结构体是带类型的字段的集合，支持指针实现对结构体的修改和开销节约

Go没有类的概念，而是通过结构体和接口实现比面向对象具有更高的扩展性和灵活性

#### 自定义类型和类型别名

使用`type`关键字

```go
// MyInt 类型别名
type MyInt = int

// NewInt 自定义类型
type NewInt int
```

#### 定义结构体

结构体是值类型，而字段和结构体名首字母的大小写表明私有或公共

```go
type user struct {
	name     string
	password string
}
```



#### 实例化

```go
a := user{name: "Ltd", password: "1024"}
b := user{"Ltd", "1024"}
c := user{name: "Ltd"}
c.password = "1024"
var d user
d.name = "Ltd"
d.password = "1024"
```



#### 结构体方法

注意结构体方法绑定的时候是否用指针类型的，一种就是修改副本，一种是修改原始的

[官方解释](https://go.dev/doc/faq#methods_on_values_or_pointers)

```go
func (p *Point) use() {
}
func (p Point) use() {
}
```



#### 面试题

```go
m := make(map[string]*student)
stus := []student{
   {name: "L", age: 11},
   {name: "T", age: 12},
   {name: "D", age: 13},
}

for _, stu := range stus {
   fmt.Println(&stu)
   // 这里指向的是stu的内存空间，而range是值拷贝，当遍历完了都指向了最后一个stu的地址
   m[stu.name] = &stu
}
for k, v := range m {
   fmt.Println(k, "=>", v.name)
}
```

特别注意对slice和map数据类型都包含了指向底层数据的指针，对他们赋值的时候需要特别注意，后续修改的话结构体字段值也会改变

### 2.9字符串操作

```go
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
```

#### 格式化

标准包fmt，Printf、Sprintf方法

可以用 %v 来打印任意类型的变量，%T打印数据类型



### 2.10Json处理

结构体和json格式转换，只要结构体每个字段第一个字母大写就可以使用`json.Marshal`来序列化，而反序列化结果每个字段也是大写字母开头`json.Unmarshal`
