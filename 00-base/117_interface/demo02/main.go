package main

import "fmt"

// 使用过程中需要注意以下几点：
// Go 中接口声明的方法并不要求需要全部公开。
// 直接用接口类型作为变量时，赋值必须是类型的指针。
type Account interface {
	getBalance() int
}

type CreditCard struct {
	balance int
	limit   int
}

func (c *CreditCard) getBalance() int {
	return c.balance
}

func main() {
	c := CreditCard{balance: 100, limit: 1000}
	var a Account = &c
	fmt.Println(a.getBalance())
}
