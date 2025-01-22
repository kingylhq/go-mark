package main

import (
	"fmt"
	"unsafe"
)

func main() {

	sliceFunc()

	fmt.Println("*****************************")

	updateArrSlicePtr()

	fmt.Println("*****************************")

	copySlice()
}

// 切片
// 切片(Slice)并不是数组或者数组指针，而是数组的一个引用，
// 切片本身是一个标准库中实现的一个特殊的结构体，这个结构体中有三个属性，分别代表数组指针、长度、容量。
// 具体可以查看 golang 源码仓库中 src/runtime/slice.go 文件中：
type slice struct {
	array unsafe.Pointer
	len   int
	cap   int
}

// 1.13.1 声明与初始化切片
// 切片的申明方式与声明数组的方式非常相似，与数组相比，切片不用声明长度:
// var <slice name> []<type>
func sliceFunc() {

	// 方式1，声明并初始化一个空的切片
	var s1 []int = []int{}
	fmt.Println(s1)
	fmt.Println("*****************************")

	// 方式2，类型推导，并初始化一个空的切片
	var s2 = []int{}
	fmt.Println(s2)
	fmt.Println("*****************************")

	// 方式3，与方式2等价
	s3 := []int{}
	fmt.Println(s3)
	fmt.Println("*****************************")

	// 方式4，与方式1、2、3 等价，可以在大括号中定义切片初始元素
	s4 := []int{1, 2, 3, 4}
	fmt.Println(s4)
	fmt.Println("*****************************")

	// 方式5，用make()函数创建切片，创建[]int类型的切片，指定切片初始长度为0
	s5 := make([]int, 0)
	fmt.Println(s5)
	fmt.Println("*****************************")

	// 方式6，用make()函数创建切片，创建[]int类型的切片，指定切片初始长度为2，指定容量参数4
	s6 := make([]int, 2, 4)
	fmt.Println(s6)
	fmt.Println("*****************************")

	// 方式7，引用一个数组，初始化切片
	//a := [5]int{6, 5, 4, 3, 2}
	arr := [10]string{"lsy", "wdd", "yy", "cf", "nn", "cs", "ll", "gg", "ww", "mi"}
	fmt.Println(arr)
	fmt.Println("*****************************")

	// 从数组下标2开始，直到数组的最后一个元素 [yy cf nn cs ll gg ww mi] [2, 10]
	s7 := arr[2:]
	fmt.Println(s7)
	fmt.Println("*****************************")

	// 从数组下标1开始，直到数组下标3的元素，创建一个新的切片 [wdd yy], [1,3)
	s8 := arr[1:3]
	fmt.Println(s8)
	fmt.Println("*****************************")

	// 从0到下标2的元素，创建一个新的切片 [lsy wdd] [0,2)
	s9 := arr[:2]
	fmt.Println(s9)
	fmt.Println("*****************************")
}

// 当切片是基于同一个数组指针创建出来时，修改数组中的值时，同样会影响到这些切片。
func updateArrSlicePtr() {
	a := [5]int{6, 5, 4, 3, 2}
	// 从数组下标2开始，直到数组的最后一个元素
	s7 := a[2:]
	// 从数组下标1开始，直到数组下标3的元素，创建一个新的切片
	s8 := a[1:3]
	// 从0到下标2的元素，创建一个新的切片
	s9 := a[:2]
	fmt.Println(s7)
	fmt.Println(s8)
	fmt.Println(s9)
	a[0] = 9
	a[1] = 8
	a[2] = 7
	fmt.Println(s7)
	fmt.Println(s8)
	fmt.Println(s9)
}

// 1.13.2 使用切片
// 1.13.2.1 访问切片
// 访问切片中的元素，与访问数组一样
func useSlice() {
	s1 := []int{5, 4, 3, 2, 1}
	// 下标访问切片
	e1 := s1[0]
	e2 := s1[1]
	e3 := s1[2]
	fmt.Println(s1)
	fmt.Println("*****************************")
	fmt.Println(e1)
	fmt.Println("*****************************")
	fmt.Println(e2)
	fmt.Println("*****************************")
	fmt.Println(e3)
	fmt.Println("*****************************")

	// 向指定位置赋值
	s1[0] = 10
	s1[1] = 9
	s1[2] = 8
	fmt.Println(s1)

	// range迭代访问切片
	for i, v := range s1 {
		fmt.Println("before modify, s1[%d] = %d", i, v)
	}
}

// 切片还可以使用 len() 和 cap() 函数访问切片的长度和容量。
// 长度表示切片可以访问到底层数组的数据范围。
// 容量表示切片引用的底层数组的长度。
// 当切片是 nil 时，len() 和 cap() 函数获取的到值都是 0。
// 切片的长度小于等于切片的容量。
func lenCap() {

	var nilSlice []int
	fmt.Println("nilSlice length:", len(nilSlice))
	fmt.Println("nilSlice capacity:", cap(nilSlice))

	s2 := []int{9, 8, 7, 6, 5}
	fmt.Println("s2 length: ", len(s2))
	fmt.Println("s2 capacity: ", cap(s2))
}

// 1.13.2.2 切片添加元素
// 切片是变长的，可以向切片追加新的元素，可以使用内置的 append() 向切片追加元素。
// 内置函数 append() 只有切片类型可以使用，第一个参数必须是切片类型，后面追加的元素参数是变长类型，一次可以追加多个元素到切片。
// 并且每次 append() 都会返回一个新的切片引用。
func optElement() {
	s3 := []int{}
	fmt.Println("s3 = ", s3)
	// append函数追加元素
	s3 = append(s3)
	s3 = append(s3, 1)
	s3 = append(s3, 2, 3)
	fmt.Println("s3 = ", s3)

	//除了使用 append() 函数向切片追加元素以外，还可以使用 append() 向指定位置添加元素，以及移除指定位置的元素。
	//向指定位置添加元素的代码示例：
	s4 := []int{1, 2, 4, 5}
	s4 = append(s4[:2], append([]int{3}, s4[2:]...)...)
	fmt.Println("s4 = ", s4)

	//移除指定位置元素代码示例：
	s5 := []int{1, 2, 3, 5, 4}
	s5 = append(s5[:3], s5[4:]...)
	fmt.Println("s5 = ", s5)
}

// 1.13.2.3 复制切片
// 可以使用内置函数 copy() 把某个切片中的所有元素复制到另一个切片，复制的长度是它们中最短的切片长度。
func copySlice() {

	//src1 := []int{1, 2, 3}
	//dst1 := make([]int, 4, 5)
	//
	//fmt.Println("before copy, src1 = ", src1)
	//fmt.Println("before copy, dst1 = ", dst1)
	//copy(dst1, src1)
	//fmt.Println("before copy, src1 = ", src1)
	//fmt.Println("before copy, dst1 = ", dst1)

	src2 := []int{1, 2, 3, 4, 5}
	dst2 := make([]int, 3, 3)

	fmt.Println("before copy, src2 = ", src2)
	fmt.Println("before copy, dst2 = ", dst2)

	copy(dst2, src2)

	fmt.Println("before copy, src2 = ", src2)
	fmt.Println("before copy, dst2 = ", dst2)
}
