package main

import (
	"fmt"
)

func main() {

	//1.3 定义变量
	//在 Go 中，变量主要分为两类：全局变量、局部变量。

	// 全局变量
	var s1 string = "Hello"
	var zero int
	var b1 = true

	var (
		i  int = 123
		b2 bool
		s2 = "test"
	)

	var (
		group = 2
	)
	fmt.Println(s1)
	fmt.Println(zero)
	fmt.Println(b1)
	fmt.Println(i)
	fmt.Println(b2)
	fmt.Println(s2)
	fmt.Println(group)

	fmt.Println("*****************************")
	method1()
	fmt.Println("*****************************")

	sum, sub := method5(10, 5)
	fmt.Println(sum)
	fmt.Println(sub)
	fmt.Println("*****************************")

	method6()

}

// 1.3.2 局部变量
// 在函数内或方法内定义的变量叫做局部变量
func method1() {
	// 方式1，类型推导，用得最多
	a := 1
	// 方式2，完整的变量声明写法
	var b int = 2
	// 方式3，仅声明变量，但是不赋值
	var c int
	fmt.Println(a, b, c)
}

// 方式4，直接在返回值中声明
func method2() (a int, b string) {
	// 这种方式必须声明return关键字
	// 并且同样不需要使用，并且也不用必须给这种变量赋值
	return 1, "test"
}

func method3() (a int, b string) {
	a = 1
	b = "test"
	return
}

func method4() (a int, b string) {
	return
}

func method5(x int, y int) (sum int, sub int) {
	return (x + y), (x - y)
}

//1.3.3 多变量定义
//全局变量和局部变量都支持一次声明和定义多个变量。

var a1, b1, c1 int = 10, 50, 400

var m1, m2, m3 int

func method6() {

	fmt.Println("*****************************")

	fmt.Println(a1)
	fmt.Println(b1)
	fmt.Println(c1)
	fmt.Println(m1)
	fmt.Println(m2)
	fmt.Println(m3)

	fmt.Println("*****************************")

	var n1, n2, n3 int = 10, 50, 90
	fmt.Println(n1)
	fmt.Println(n2)
	fmt.Println(n3)
}
