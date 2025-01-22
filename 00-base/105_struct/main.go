package main

import (
	"fmt"
	"sync"
)

// 1.5 结构体
// 在 Go 中，开发者可以通过类型别名(alias types)和结构体的形式支持用户自定义类型。
//
// 结构体的目的就是把数据聚合在一起，能够便捷的操作这些数据。
//
// Go 中没有类的概念。
// Go 中，全局变量、全局常量、结构体、字段、方法，只有两种公开类型，公开与非公开。非公开是针对包级别的，
// 也就是说如果全局变量声明在不同的源文件中，但是这些源文件属于相同的包，那么这些中的代码可以引用这些不公开的全局变量。
// 不属于相同的包就访问不到了。并且公开的属性是首字母大写，非公开的属性首字母是小写，仅按照这个规则来定义是否公开。

//1.5.1 定义结构体
//在 Go 中最常用的方法使用 type 和 struct 关键字定义一个结构体，结构体中的字段都有不同的名字，并且字段可以是任意类型，
//比如结构体本身、指针甚至是函数：

// 方式 1：
type Person struct {
	Name  string
	Age   int
	Call  func() byte
	Map   map[string]string
	Ch    chan string
	Arr   [32]uint8
	Slice []interface{}
	Ptr   *int
	once  sync.Once
}

// 方式 2：
type Custom struct {
	field1, field2, field3 byte
}

// 结构体本身除了定义字段以外，还可以针对字段添加字段标记：
type Other struct{}
type Person2 struct {
	Name  string            `json:"name" gorm:"column:<name>"`
	Age   int               `json:"age" gorm:"column:<name>"`
	Call  func()            `json:"-" gorm:"column:<name>"`
	Map   map[string]string `json:"map" gorm:"column:<name>"`
	Ch    chan string       `json:"-" gorm:"column:<name>"`
	Arr   [32]uint8         `json:"arr" gorm:"column:<name>"`
	Slice []interface{}     `json:"slice" gorm:"column:<name>"`
	Ptr   *int              `json:"-"`
	O     Other             `json:"-"`
}

// 1.5.1.1 定义匿名字段
// 结构体中的字段不是一定要有字段名，也可以仅定义类型，这种只有类型没有字段名的字段被称为匿名字段。
// 同一个结构体中相同类型的匿名字段只能同时存在一个，但是可以同时声明多个不同类型的匿名字段。
type CustomLsy struct {
	int
	string
	Other string
}

//1.5.2 定义匿名结构体
//匿名结构体是没有定义名称的结构体。
//
//匿名结构体无法定义自己的类型方法。
//
//方式 1，仅可在函数外声明，这种方式可以看成是声明了一个匿名的结构体，实例化后赋值给了的全局变量：
//
//var <Name> = struct {
//<FiledName1> <type1>
//<FiledName2> <type2>
//...
//<type3>
//<type4>
//...
//} {}
//方式 2，是方式 1 的完整写法：
//
//var <var name> = struct {
//<FiledName1> <type1> `<tag1>:"<any string>"`
//<FiledName2> <type2> `<tag2>:"<any string>"`
//...
//<type3>
//<type4>
//...
//} {
//<FiledName1>: <value1>,
//<FiledName2>: <value2>,
//...
//<type3>: <value3>,
//<type4>: <value4>,
//}
//方式 3，在函数或方法中声明匿名结构体并实例化：

//func method() {
//	<var name> := struct {
//		<FieldName1> <type1>
//		<FieldName2> <type2>
//		...
//		<type3>
//		<type4>
//	} {
//		<FieldName1>: <value1>,
//		<FieldName2>: <value2>,
//
//		<type3>: <value3>,
//		<type4>: <value4>,
//	}
//}
//匿名结构体的主要适用场景：
//
//构建测试数据，单元测试方法中一般会直接声明一个匿名结构体的切片，通过遍历切片测试方法的各个逻辑分支。
//示例代码可以参考：go-ethereum/core/type/transaction_test.go 的 TestYParityJSONUnmarshalling 方法。
//http 处理函数中的 JSON 序列化和反序列化，但是不推荐这么使用，应该定义一个正式的结构体。优点是相比 map[string]interface{}无需检查类型、
//无需检查 key 是否存在并减少相关的代码检查。

// 1.5.3 定义嵌套结构体
// 在 Go 中，不存在类似 Java 中的继承关系，只有结构体之间的嵌套关系，但是可以达到类似继承的效果。
//
// type <Name1> struct {
// ...
// }
//
// type <Name2> struct {
// <Name1>
// ...
// }
// 根据上面代码中的声明关系，Name1 结构体中的声明的所有字段和方法，如果在 Name2 结构体中不存在相同的字段名和方法名，
// 那么 Name2 结构体的示例是可以直接调用的。而相同的字段名和方法名的这些属性，也可以通过 Name2.Name1 的方式获取到 Name1 中公开的字段和方法。
type A struct {
	a string
}

func (a A) string() string {
	return a.a
}

func (a A) stringA() string {
	return a.a
}

func (a A) setA(v string) {
	a.a = v
}

func (a *A) stringPA() string {
	return a.a
}

func (a *A) setPA(v string) {
	a.a = v
}

type B struct {
	A
	b string
}

func (b B) string() string {
	return b.b
}

func (b B) stringB() string {
	return b.b
}

type C struct {
	B
	a string
	b string
	c string
	d []byte
}

func (c C) string() string {
	return c.c
}

func (c C) modityD() {
	c.d[2] = 3
}

func callStructMethod() {
	var a A
	a = A{
		a: "a",
	}
	a.string()
}

func NewC() C {
	return C{
		B: B{
			A: A{
				a: "ba",
			},
			b: "b",
		},
		a: "ca",
		b: "cb",
		c: "c",
		d: []byte{1, 2, 3},
	}
}

func main() {
	c := NewC()
	cp := &c
	fmt.Println(c.string())
	fmt.Println(c.stringA())
	fmt.Println(c.stringB())

	fmt.Println(cp.string())
	fmt.Println(cp.stringA())
	fmt.Println(cp.stringB())

	//c.setA("1a")
	//fmt.Println("------------------c.setA")
	//fmt.Println(c.A.a)
	//fmt.Println(cp.A.a)

	//cp.setA("2a")
	//fmt.Println("------------------cp.setA")
	//fmt.Println(c.A.a)
	//fmt.Println(cp.A.a)

	//c.setPA("3a")
	//fmt.Println("------------------c.setPA")
	//fmt.Println(c.A.a)
	//fmt.Println(cp.A.a)

	//cp.setPA("4a")
	//fmt.Println("------------------cp.setPA")
	//fmt.Println(c.A.a)
	//fmt.Println(cp.A.a)

	//cp.modityD()
	//fmt.Println("------------------cp.modityD")
	//fmt.Println(cp.d)
}

//type A struct {
//	a     string
//	bytes [2]byte
//}
//
//func (a A) string() string {
//	return a.a
//}
//
//func (a A) stringA() string {
//	return a.a
//}
//
//func (a A) setA(v string) {
//	a.a = v
//}
//
//func (a *A) stringPA() string {
//	return a.a
//}
//
//func (a *A) setPA(v string) {
//	a.a = v
//}
//
//func value(a A, value string) {
//	a.a = value
//}
//
//func point(a *A, value string) {
//	a.a = value
//}
//
//func main() {
//	a := A{
//		a: "a",
//	}
//
//	value(a, "any")
//
//	point(&a, "any")
//
//	pa := &a
//
//	// a *A
//	// a.setPA("pa")
//
//	// a A
//	fmt.Println(a.string())
//	// a A
//	fmt.Println(a.stringA())
//	// a *A
//	fmt.Println(a.stringPA())
//
//	// a A
//	fmt.Println(pa.string())
//	// a A
//	fmt.Println(pa.stringA())
//	// a *A
//	fmt.Println(pa.stringPA())
//}
