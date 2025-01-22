package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"go-mark/02-go-ethereum-develop/constants"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	store "go-mark/02-go-ethereum-develop/18-contract-deploy/contracts"
)

// Querying a Smart Contract(查询智能合约)
// 这写章节需要了解如何将智能合约的 ABI 编译成 Go 的合约文件。如果你还没看， 前先读上一个章节 。
// 在上个章节我们学习了如何在 Go 应用程序中初始化合约实例。 现在我们将使用新合约实例提供的方法来阅读智能合约。
// 如果你还记得我们在部署过程中设置的合约中有一个名为 version 的全局变量。 因为它是公开的，这意味着它们将成为我们自动创建的 getter 函数。
// 常量和 view 函数也接受 bind.CallOpts 作为第一个参数。了解可用的具体选项要看相应类的文档 一般情况下我们可以用 nil。
func main() {
	client, err := ethclient.Dial(constants.HTTPS_URL)
	if err != nil {
		log.Fatal(err)
	}

	address := common.HexToAddress("0x1742a898e7218f1591B8B117932e86F3392237b1")
	instance, err := store.NewStore(address, client)
	if err != nil {
		log.Fatal(err)
	}

	//version, err := instance.Version(nil)
	//if err != nil {
	//	log.Fatal(err)
	//}
	version, err := instance.Version(&bind.CallOpts{Pending: false, Context: context.Background()})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("智能合约Store版本: ", version) // "2.0"
}
