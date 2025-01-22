package main

import (
	"context"
	"fmt"
	"log"
	"regexp"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {

	checkAddress()

}

// 如何确认一个地址并确定其是否为智能合约地址。
// 检查地址是否有效
func checkAddress() {

	// 使用简单的正则表达式来检查以太坊地址是否有效
	re := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")

	fmt.Printf("is valid: %v\n", re.MatchString("0x323b5d4c32345ced77393b3530b1eed0f346429d")) // is valid: true
	fmt.Printf("is valid: %v\n", re.MatchString("0xZYXb5d4c32345ced77393b3530b1eed0f346429d")) // is valid: false

	client, err := ethclient.Dial("https://cloudflare-eth.com")
	if err != nil {
		log.Fatal(err)
	}

	// 0x Protocol Token (ZRX) smart contract address
	//检查地址是否为账户或智能合约
	//我们可以确定，若在该地址存储了字节码，该地址是智能合约。这是一个示例，在例子中，我们获取一个代币智能合约的字节码并检查其长度以验证它是一个智能合约：
	address := common.HexToAddress("0xe41d2489571d322189246dafa5ebde1f4699f498")
	bytecode, err := client.CodeAt(context.Background(), address, nil) // nil is latest block
	if err != nil {
		log.Fatal(err)
	}

	isContract := len(bytecode) > 0
	fmt.Printf("is contract 33: %v\n", isContract) // is contract: true

	// a random user account address
	// 当地址上没有字节码时，我们知道它不是一个智能合约，它是一个标准的以太坊账户
	//address = common.HexToAddress("0x8e215d06ea7ec1fdb4fc5fd21768f4b34ee92ef4")
	//address = common.HexToAddress(constants.MateMaskAccount1Address)
	address = common.HexToAddress("0x21968594DC97350ecb86cD233953DEab9B2FF339")
	bytecode, err = client.CodeAt(context.Background(), address, nil) // nil is latest block
	if err != nil {
		log.Fatal(err)
	}

	isContract = len(bytecode) > 0
	fmt.Printf("is contract 222: %v\n", isContract) // is contract: false
}
