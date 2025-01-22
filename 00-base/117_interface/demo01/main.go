package main

import "fmt"

// 1.17 interface 接口
// 在 Go 中接口是一种抽象类型，是一组方法的集合，里面只声明方法，而没有任何数据成员。
// 而在 Go 中实现一个接口也不需要显式的声明，只需要其他类型实现了接口中所有的方法，就是实现了这个接口。
// 定义一个接口：
// type <interface_name> interface {
// <method_name>(<method_params>) [<return_type>...]
// ...
// }
// PaymentMethod 接口定义了支付方法的基本操作
type PayMethod interface {
	// 接口，包含两个方法，其中一个是支付方法，另一个是获取余额方法(Account接口嵌入到了PayMethod中，
	// 则PayMethod它就包含了Account接口中的所有方法)
	Account
	Pay(amount int) bool
}

type Account interface { // 接口包含一个方法
	GetBalance() int
}

// CreditCard 结构体实现 PaymentMethod 接口
type CreditCard struct {
	balance int // 已经消费额度
	limit   int // 最大允许消费限制
}

func (c *CreditCard) Pay(amount int) bool {
	if c.balance+amount <= c.limit {
		c.balance += amount
		fmt.Printf("信用卡支付成功: %d\n", amount)
		return true
	}
	fmt.Println("信用卡支付失败: 超出额度")
	return false
}

func (c *CreditCard) GetBalance() int {
	return c.balance
}

// DebitCard 结构体实现 PaymentMethod 接口
type DebitCard struct {
	balance int
}

func (d *DebitCard) Pay(amount int) bool {
	if d.balance >= amount {
		d.balance -= amount
		fmt.Printf("借记卡支付成功: %d\n", amount)
		return true
	}
	fmt.Println("借记卡支付失败: 余额不足")
	return false
}

func (d *DebitCard) GetBalance() int {
	return d.balance
}

// 使用 PaymentMethod 接口的函数
func purchaseItem(p PayMethod, price int) {
	if p.Pay(price) {
		fmt.Printf("购买成功，剩余余额: %d\n", p.GetBalance())
	} else {
		fmt.Println("购买失败")
	}
}

func main() {
	creditCard := &CreditCard{balance: 0, limit: 1000}
	debitCard := &DebitCard{balance: 500}

	fmt.Println("使用信用卡购买:")
	purchaseItem(creditCard, 800)

	fmt.Println("\n使用借记卡购买:")
	purchaseItem(debitCard, 300)

	fmt.Println("\n再次使用借记卡购买:")
	purchaseItem(debitCard, 300)
}
