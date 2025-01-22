package main

import "fmt"

func main() {

	//整型：int，int8，int16，int32，int64，uint，uint8，uint16，uint32，uint64，uintptr。
	//
	//以 u 开头的整型被称为无符号整数类型，即都是非负数。而后面的数字代表这个值在内存中占有多少二进制位。
	//
	//比如 uint8 将占有 8 位，其最大值是 255（即
	//2
	//8
	//−
	//1
	//）。
	//
	//比如 int8，同样占有 8 位，但是最高位是符号位，所以最大值是 127（即
	//2
	//7
	//−
	//1
	//），最小值是-128（即$-2^7$）。
	//
	//以 uint8 举例，为整型赋值时，可以直接使用十六进制、八进制、二进制以及十进制声明。
	// 十六进制
	var a uint8 = 0xF
	var b uint8 = 0xf

	// 八进制
	var c uint8 = 017
	var d uint8 = 0o17
	var e uint8 = 0o17

	// 二进制
	var f uint8 = 0b1111
	var g uint8 = 0b1111

	// 十进制
	var h uint8 = 15
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)
	fmt.Println(h)

	//浮点数：float32，float64。
	//float32 是单精度浮点数，精确到小数点后 7 位。
	//float64 是双精度浮点数，可以精确到小数点后 15 位。
	var float1 float32 = 10.12345678987654
	float2 := 10.0
	var float3 float64 = 10.12345678987654
	fmt.Println(float1)
	fmt.Println(float2)
	fmt.Println(float3)

	//对于浮点类型需要被自动推到的变量，其类型都会被自动设置为 float64，而不管它的字面量是否是单精度。
	//所以用上面的例子再加一行：
	//float1 = float2

	//必须强制类型转换成 float32 才可以编译通过：
	float1 = float32(float2)

	//在实际开发中，也更推荐使用 float64 类型，因为官方标准库 math 包中，所有有关数学运算的函数的入参都是 float64 类型。
	//复数：complex64，complext128
	//整型与浮点数日常中常见的数字都是实数，复数是实数的延伸，可以通过两个部分构成，一个实部，一个虚部，常见的声明形式如下：
	//var z complex64 = a + bi
	//a 和 b 均为实数，i 为虚数单位，当 b = 0 时，z 就是常见的实数。
	//当 a = 0 且 b ≠ 0 时，将 z 为纯虚数。
	//var c1 complex64
	//c1 = 1.10 + 0.1i
	c2 := 1.10 + 0.1i
	//c3 := complex(1.10, 0.1) // c2与c3是等价的
	//复数类型和浮点型有类似的机制，默认的自动推到的复数类型是 complex128。
	//并且 Go 还提供了内置函数 real 和 imag，分别获取复数的实部和虚部：

	x := real(c2)
	y := imag(c2)
	fmt.Println("*****************************")
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println("*****************************")

	lsy := 99.20 + 0.5i
	l1 := real(lsy)
	l2 := imag(lsy)
	fmt.Println(l1)
	fmt.Println(l2)
	fmt.Println("*****************************")

	//1.2.1.3 byte 类型
	//byte 是 uint8 的内置别名，可以把 byte 和 uint8 视为同一种类型。
	//在 Go 中，字符串可以直接被转换成 []byte（byte 切片）。
	var s string = "Hello, world!"
	var bytes []byte = []byte(s)
	fmt.Println("convert \"Hello, world!\" to bytes: ", bytes)
	//同时[]byte 也可以直接转换成 string。

	var bytes2 []byte = []byte{72, 101, 108, 108, 111, 44, 32, 119, 111, 114, 108, 100, 33}
	var s2 string = string(bytes2)
	fmt.Println(s2)
	fmt.Println("*****************************")

	// a = 97, b = 98, c = 99
	var mark string = "abclsymark"
	var bl []byte = []byte(mark)
	fmt.Println(mark)
	fmt.Println(bl)
	fmt.Println("*****************************")

	//1.2.1.4 rune 类型
	//rune 是 int32 的内置别名，可以把 rune 和 int32 视为同一种类型。但 rune 是特殊的整数类型。
	//在 Go 中，一个 rune 值表示一个 Unicode 码点。一般情况下，一个 Unicode 码点可以看做一个 Unicode 字符。
	// 有些比较特殊的 Unicode 字符有多个 Unicode 码点组成。
	//一个 rune 类型的值由一个个被单引号包住的字符组成，比如：
	var r1 rune = 'a'
	var r2 rune = '世'
	fmt.Println(r1)
	fmt.Println(r2)
	//字符串可以直接转换成 []rune（rune 切片）。

	var s5 string = "abc，你好，世界！"
	var runes []rune = []rune(s5)
	fmt.Println(s5)
	fmt.Println(runes)
	fmt.Println("*****************************")

	//1.2.2 字符串 - string
	//在 Go 中，字符串是 UTF-8 编码的，并且所有的 Go 源码都必须是 UTF-8 编码。
	//字符串的字面量有两种形式。
	//一种是解释型字面表示（interpreted string literal，双引号风格）。

	var s6 string = "Hello\nworld!\n"
	//另一种是直白字面量表示（raw string literal， 反引号风格）。
	var s7 string = `Hello world!`
	//上面举例的两种字符串是等价的：
	var s8 string = "Hello\nworld!\n"
	var s9 string = `Hello world!`
	fmt.Println(s6 == s7)
	fmt.Println(s8 == s9)

}
