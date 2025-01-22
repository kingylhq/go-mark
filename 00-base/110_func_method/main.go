package main

import "fmt"

//函数、闭包与方法
//1.10.1 函数
//函数只有三个主要部分，分别是名称、参数列表、返回类型列表。
//其中名称是必须的，参数列表和返回类型列表是可选的，也就是说函数可以没有参数，也没有返回值。

// 定义函数
func funcLsy() {
	fmt.Println("Hello, world! 定义函数")
}

//1.10.2 闭包
//闭包，也被称为匿名函数，顾名思义，即没有函数名，通常在函数内或者方法内定义，或者作为参数、返回值进行传递。
//匿名函数的优势是可以直接使用当前函数内在匿名函数声明之前声明的变量。
//定义方式：
//// 声明函数变量
//var <closure name> func(<parameter list>) (<return types>)
//// 声明闭包
//var  <closure name> func(<parameter list>) (<return types>) = func(<parameter list>) (<return types>) {
//	<expressions>
//	...
//}
//// 声明并立刻执行
//func(<parameter list>) (<return types>) {
//	<expressions>
//	...
//}(<value list>)
//
//// 作为参数，并调用
//func <function name>(...,<name> func(<parameter list>) (<return types>), ...) {
//	...
//	<var1>,... := <name>(<value list>)
//	...
//}
//
//// 作为返回值
//func <function name>(...) (func(<parameter list>) (<return types>)) {
//	...
//	<var1>,... := <name>(<value list>)
//	...
//}

type A struct {
	i int
}

func (a *A) add(v int) int {
	a.i += v
	return a.i
}

// 声明函数变量
var function1 func(int) int

// 声明闭包
var squart2 func(int) int = func(p int) int {
	p *= p
	return p
}

func main() {
	a := A{1}
	// 把方法赋值给函数变量
	function1 = a.add

	// 声明一个闭包并直接执行
	// 此闭包返回值是另外一个闭包（带参闭包）
	returnFunc := func() func(int, string) (int, string) {
		fmt.Println("this is a anonymous function")
		return func(i int, s string) (int, string) {
			return i, s
		}
	}()

	// 执行returnFunc闭包并传递参数
	ret1, ret2 := returnFunc(1, "test")
	fmt.Println("call closure function, return1 = ", ret1, "; return2 = ", ret2)

	fmt.Println("a.i = ", a.i)
	fmt.Println("after call function1, a.i = ", function1(3))
	fmt.Println("a.i = ", a.i)
}

func add(a, b int) int {

	// 函数中遇到defer，会将defer后面的代码压入栈中，等函数返回时，再执行
	defer fmt.Println("a=", a)
	defer fmt.Println("b=", b)
	// 栈的特点是先进后出，函数执行完毕后，取出栈帧语句，按照先进后出开始执行,因此这里先打印b，在打印a

	// defer 压入栈帧的数据还是之前形参的值，压栈的时候将值拷贝一份
	a += 10
	b += 20
	return a + b
}
