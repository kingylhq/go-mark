package main

import (
	"encoding/json"
	"fmt"
	"sync"
	"time"
)

func main() {

	//mapDefine()

	//userMap()
	//
	//m := make(map[string]int)
	//m["a"] = 1
	//receiveMap(m)
	//fmt.Println("m =", m)
	//
	//concurrentMapLock()

	genJsonMap()
}

//1.14 map 集合
//在 Go 中，map 集合是无序的键值对集合。相比切片和数组，map 集合对索引的自定义程度更高，可以使用任意类型作为索引，也可以存储任意类型的数据。
//但是 map 集合中，存储的键值对的顺序是不确定的。当获取 map 集合中的值时，如果键不存在，则返回类型的零值。
//1.14.1 声明 map 集合
//方式 1，仅声明 map：
//var <map name> map[<key type>]<value type>
//方式 2，使用内置函数 make() 初始化：
//<map name> := make(map[<key type>]<value type>)
//// 还可以使用map，提前指定容量
//<map name> := make(map[<key type>]<value type>, <capacity>)
//指定合适的初始容量，可以减少使用 map 存储键值对时触发扩容，减少扩容操作可以提高一些使用 map 的性能。
//方式 3，在初始化时，同时插入键值对：
//// 不会插入任何键值对
//<map name> := map[<key type>]<value type> {}
//// 插入键值对
//<map name> := map[<key type>]<value type> {
//<key1>: <value1>,
//<key2>: <value2>,
//...
//}

func mapDefine() {
	var m1 map[string]string
	fmt.Println("m1 length:", len(m1))

	m2 := make(map[string]string)
	fmt.Println("m2 length:", len(m2))
	fmt.Println("m2 =", m2)

	m3 := make(map[string]string, 10)
	fmt.Println("m3 length:", len(m3))
	fmt.Println("m3 =", m3)

	m4 := map[string]string{}
	fmt.Println("m4 length:", len(m4))
	fmt.Println("m4 =", m4)

	m5 := map[string]string{
		"key1": "value1",
		"key2": "value2",
	}
	fmt.Println("m5 length:", len(m5))
	fmt.Println("m5 =", m5)
}

// 1.14.2 使用 map 集合
// 获取元素：
// <value> := <map name>[<key>]
// <value>,<exist flag> := <map name>[<key>]
// 插入或修改键值对：
// <map name>[<key>] = <value>
// 通过内置函数 len() 获取 map 中键值对数量：
// length := len(<map name>)
// 遍历 map 集合：
// for <key name>, <value name> := range <map name> {
// <expression>
// ...
// }
// for <key name> := range <map name> {
// <expression>
// ...
// }
// 使用内置函数 delete() 删除 map 集合中指定 key：
// delete(<map name>, <key>)
func userMap() {
	//n := make(map[string]int, 10)
	//n["1"] = int(1)

	mm := make(map[string]string, 10)
	mm["lsy"] = "lsy"
	mm["yy"] = "yy"
	mm["wdd"] = "wdd"
	mm["cf"] = "cf"
	mm["pp"] = "pp"
	lsy, exists1 := mm["lsy3"]
	fmt.Println(lsy, exists1)
	//fmt.Println(mm)
	delete(mm, "pp")
	//fmt.Println(mm)

	//m := make(map[string]int, 10)
	//m["1"] = int(1)
	//m["2"] = int(2)
	//m["3"] = int(3)
	//m["4"] = int(4)
	//m["5"] = int(5)
	//m["6"] = int(6)
	//
	//// 获取元素
	//value1 := m["1"]
	//fmt.Println("m[\"1\"] =", value1)
	//
	//value1, exist := m["1"]
	//fmt.Println("m[\"1\"] =", value1, ", exist =", exist)
	//
	//valueUnexist, exist := m["10"]
	//fmt.Println("m[\"10\"] =", valueUnexist, ", exist =", exist)
	//
	//// 修改值
	//fmt.Println("before modify, m[\"2\"] =", m["2"])
	//m["2"] = 20
	//fmt.Println("after modify, m[\"2\"] =", m["2"])
	//
	//// 获取map的长度
	//fmt.Println("before add, len(m) =", len(m))
	//m["10"] = 10
	//fmt.Println("after add, len(m) =", len(m))
	//
	//// 遍历map集合main
	//for key, value := range m {
	//	fmt.Println("iterate map, m[", key, "] =", value)
	//}
	//
	//// 使用内置函数删除指定的key
	//_, exist_10 := m["10"]
	//fmt.Println("before delete, exist 10: ", exist_10)
	//delete(m, "10")
	//_, exist_10 = m["10"]
	//fmt.Println("after delete, exist 10: ", exist_10)
	//
	//// 在遍历时，删除map中的key
	//for key := range m {
	//	fmt.Println("iterate map, will delete key:", key)
	//	delete(m, key)
	//}
	//fmt.Println("m = ", m)
}

// 1.14.3 map 作为参数
// map 集合也是引用类型，和切片一样，将 map 集合作为参数传给函数或者赋值给另一个变量，它们都指向同一个底层数据结构，
// 对 map 集合的修改，都会影响到原始实参。
func receiveMap(param map[string]int) {
	fmt.Println("before modify, in receiveMap func: param[\"a\"] = ", param["a"])
	param["a"] = 2
	param["b"] = 3
}

// 1.14.4 并发时使用 map 集合
// 下面示例模拟 map 集合在程序中，同时被多个 goroutine 读写：
func concurrentMap() {
	m := make(map[string]int)

	go func() {
		for {
			m["a"]++
		}
	}()

	go func() {
		for {
			m["a"]++
			fmt.Println(m["a"])
		}
	}()

	select {
	case <-time.After(time.Second * 5):
		fmt.Println("timeout, stopping")
	}
}

// 运行 concurrentMap()函数，会直接报错提示：
// fatal error: concurrent map writes
// 如果读写操作同时存在也会报错提示：
// fatal error: concurrent map read and map write
// 当 map 集合会被并发访问时，需要在使用 map 集合时，添加互斥锁：
func concurrentMapLock() {
	m := make(map[string]int)
	var wg sync.WaitGroup
	var lock sync.Mutex
	wg.Add(2)

	go func() {
		defer wg.Done()
		for { // 最好限制循环次数以避免无限循环
			lock.Lock()
			m["a"]++
			lock.Unlock()
		}
	}()

	go func() {
		defer wg.Done()
		for { // 最好限制循环次数以避免无限循环
			lock.Lock()
			m["a"]++
			fmt.Println(m["a"])
			lock.Unlock()
		}
	}()

	done := make(chan struct{})
	go func() {
		wg.Wait()
		close(done)
	}()

	select {
	case <-time.After(time.Second * 5):
		fmt.Println("timeout, stopping")
	case <-done:
		fmt.Println("all goroutines finished")
	}
}

//也可以使用 Go 标准库中的实现 sync.Map，但是 sync.Map 适用于读多写少的场景，并且内存开销会比普通的 map 集合更大。
//所以碰到这种情况，更推荐使用普通的互斥锁来保证 map 集合的并发读写的线程安全性。

func genJsonMap() {
	res := make(map[string]interface{})
	res["code"] = 200
	res["msg"] = "success"
	res["data"] = map[string]interface{}{
		"username": "Tom",
		"age":      "30",
		"hobby":    []string{"读书", "爬山"},
	}
	fmt.Println("map data :", res)

	//序列化
	jsons, errs := json.Marshal(res)
	if errs != nil {
		fmt.Println("json marshal error:", errs)
	}
	fmt.Println("--- map to json ---")
	fmt.Println("json data :", string(jsons))

	//反序列化
	//res2 := make(map[string]interface{})
	//errs = json.Unmarshal([]byte(jsons), &res2)
	//if errs != nil {
	//	fmt.Println("json marshal error:", errs)
	//}
	//fmt.Println("")
	//fmt.Println("--- json to map ---")
	//fmt.Println("map data :", res2)
}
