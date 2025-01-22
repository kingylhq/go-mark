package main

import (
	"fmt"
	"unsafe"
)

func main() {

	method1()

	fmt.Println("*****************************")

	method2()

	fmt.Println("***************************** method3")

	method3()

	fmt.Println("***************************** method4")

	method4()

}

// 1.4 指针
// Go 中，指针是一个变量，它存储了另一个变量的内存地址。
//
// 通过指针，可以访问存储在指定内存地址中的数据。
//
// 1.4.1 指针声明与初始化
// 在 Go 中，声明一个指针类型变量需使用星号 * 标识：
//
// var <name> *<type>
// 初始化指针必须通过另外一个变量， 如果没有赋值：
//
// p = &<var name>
// 也可以使用一个结构体实例或者变量直接声明并且赋值给一个指针：
//
// p := &<struct type>{}
// p := &<var name>
// 同时还可以获取指针的指针：
//
// var p **<type>
func method1() {
	var p1 *int
	var p2 *string

	i := 1
	s := "Hello"
	// 基础类型数据，必须使用变量名获取指针，无法直接通过字面量获取指针
	// 因为字面量会在编译期被声明为成常量，不能获取到内存中的指针信息
	p1 = &i
	p2 = &s

	p3 := &p2
	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(p3)

	fmt.Println(*p1)
	fmt.Println(*p2)
	fmt.Println(*p3)
}

// 1.4.2 使用指针访问值
// 使用星号引用指针来访问指针指向的值：
//
// <name1> <type1> = *p
// <name2> := *p
// 除了访问值以外，同样可以通过指针修改原始变量的值。
func method2() {
	var p1 *int
	i := 1
	p1 = &i
	fmt.Println(*p1 == i)
	*p1 = 2
	fmt.Println(i)
}

func method3() {
	a := 2
	var p *int
	fmt.Println(&a)
	p = &a
	fmt.Println(p, &a)

	var pp **int
	pp = &p
	fmt.Println(pp, p)
	**pp = 3
	fmt.Println(pp, *pp, p)
	fmt.Println(**pp, *p)
	fmt.Println(a, &a)
}

// 1.4.4 指针、unsafe.Pointer 和 uintptr
// 在 Go 中，指针不能直接参与计算，否则会在编译的时候就包错：
//
// var a int
// var p *int
// p = &a
// p = p + 1
// fmt.Println(p)
// 输出错误信息：
//
// invalid operation: p + 1 (mismatched types *int and untyped int)
// 但是 Go 中提供了其他方式，来操作指针，即引入了 unsafe.Point 类型和 uintptr 类型，来帮助我们操作指针。
//
// uintptr 类型是把内存地址转换成了一个整数，然后对这个整数进行计算后，在把 uintptr 转换成指针，达到指针偏移的效果。
//
// unsafe.Pointer 是普通指针与 uintptr 之间的桥梁，通过 unsafe.Pointer 实现三者的相互转换。
//
// *T <---> unsafe.Pointer <---> uintptr
// 把指针转换成 unsafe.Pointer:
//
// var p *<type>
// var a <type>
// p = &a
//
// up1 := unsafe.Pointer(p)
// up2 := unsafe.Pointer(&a)
// 把 unsafe.Pointer 转成 uintptr
//
// uintpr = uintptr(up1)
// 注意，这个操作非常危险，并且结果不可控，在一般情况下是不需要进行这种操作。
func method4() {

	//a := "Hello, world!"
	a := 10
	upA := uintptr(unsafe.Pointer(&a))
	upA += 1

	fmt.Println(&upA)
	c := (*uint8)(unsafe.Pointer(upA))
	fmt.Println(*c)

}
