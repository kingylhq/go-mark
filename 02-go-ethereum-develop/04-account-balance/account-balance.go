package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"go-mark/02-go-ethereum-develop/constants"
	"log"
	"math"
	"math/big"
)

func main() {

	//client, err := ethclient.Dial("https://cloudflare-eth.com")
	//client, err := ethclient.Dial(constants.HTTPS_URL)
	client, err := ethclient.Dial("http://localhost:7545")
	if err != nil {
		log.Fatal(err)
	}

	// 0x71c7656ec7ab88b098defb751b7401b5f6d8976f 是一个特殊的以太坊地址，通常被称为 以太坊 Faucet 地址。
	//这是一个常用的测试网络水龙头地址，提供测试用的 ETH 给开发者，用于在测试网络（如 Ropsten、Rinkeby 等）上进行智能合约开发和测试。
	//account := common.HexToAddress("0x71c7656ec7ab88b098defb751b7401b5f6d8976f")

	account := common.HexToAddress(constants.MateMaskAccount1Address)
	// 获取账户余额，参数1：当前区块上下文，参数2：要查询的账户地址，参数3：要查询的区块号（可选）
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("账户余额: ", balance) // 25893180161173005034

	blockNumber := big.NewInt(5532993)
	balanceAt, err := client.BalanceAt(context.Background(), account, blockNumber)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("指定区块的账户余额: ", balanceAt) // 25729324269165216042
	//
	fbalance := new(big.Float)
	fbalance.SetString(balanceAt.String())
	ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))
	fmt.Println("eth金额: ", ethValue) // 25.729324269165216041
	//
	pendingBalance, err := client.PendingBalanceAt(context.Background(), account)
	fmt.Println("指定账户的待处理余额: ", pendingBalance) // 41139534651425818064

	// *big.Int 是一个结构体，不能直接使用 == 进行比较，需要使用 Cmp 方法进行比较。
	fmt.Println("是否相等: ", balance.Cmp(pendingBalance) == 0)
}
