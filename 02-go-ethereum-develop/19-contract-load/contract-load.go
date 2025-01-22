package main

import (
	"fmt"
	"go-mark/02-go-ethereum-develop/constants"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	store "go-mark/02-go-ethereum-develop/18-contract-deploy/contracts"
)

// 加载智能合约
func main() {

	client, err := ethclient.Dial(constants.HTTPS_URL)
	if err != nil {
		log.Fatal(err)
	}

	address := common.HexToAddress("0x1742a898e7218f1591B8B117932e86F3392237b1")
	// 一旦使用 abigen 工具将智能合约的 ABI 编译为 Go 包，下一步就是调用“New”方法，其格式为“New”，所以在我们的例子中如果你
	// 回想一下它将是_NewStore_。 此初始化方法接收智能合约的地址，并返回可以开始与之交互的合约实例
	instance, err := store.NewStore(address, client)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("contract is loaded")
	_ = instance
}
