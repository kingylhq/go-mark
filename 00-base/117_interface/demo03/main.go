package main

import "fmt"

// 接口可以嵌套。
// 接口中声明的方法，参数可以没有名称。
// 如果函数参数使用 interface{}可以接受任何类型的实参。同样，可以接收任何类型的值也可以赋值给 interface{}类型的变量
type PayMethod interface {
	Pay(int)
}

type CreditCard struct {
	balance int
	limit   int
}

// 这个是一个方法，哟接受类型 *CreditCard
func (c *CreditCard) Pay(amount int) {
	if c.balance < amount {
		fmt.Println("余额不足")
		return
	}
	c.balance -= amount
}

func anyParam(param interface{}) {
	fmt.Println("param: ", param)
}

func main() {
	c := CreditCard{balance: 600, limit: 1000}
	c.Pay(200)
	var a PayMethod = &c
	fmt.Println("a.Pay: ", a)

	var b interface{} = &c
	fmt.Println("b: ", b)
	fmt.Println("a.Pay after: ", a)

	anyParam(c)
	anyParam(1)
	anyParam("123")
	anyParam(a)
}
