package main

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {

	queryBlocks()

}

// 在以太坊上查询和发送交易。注意这里的交易 transaction 是指广义的对以太坊状态的更改，它既可以指具体的以太币转账，代币的转账，
// 或者其他对智能合约的创建或者调用。而不仅仅是传统意义的买卖交易。
func queryBlocks() {

	client, err := ethclient.Dial("https://cloudflare-eth.com")
	if err != nil {
		log.Fatal(err)
	}

	// 区块头，调用客户端的 HeaderByNumber 来返回有关一个区块的头信息。若您传入 nil，它将返回最新的区块头
	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(header.Number.String()) // 5671744

	// 完整区块
	//调用客户端的 BlockByNumber 方法来获得完整区块。您可以读取该区块的所有内容和元数据，例如，区块号，区块时间戳，区块摘要，区块难度以及交易列表等等。
	blockNumber := big.NewInt(5671744)
	block, err := client.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(block.Number().Uint64()) // 5671744
	//fmt.Println(block.Time().Uint64())       // 1527211625
	fmt.Println(block.Time())                // 1527211625
	fmt.Println(block.Difficulty().Uint64()) // 3217000136609065
	fmt.Println(block.Hash().Hex())          // 0x9e8751ebb5069389b855bba72d94902cc385042661498a415979b7b6ee9ba4b9
	fmt.Println(len(block.Transactions()))   // 144

	// 调用 Transaction 只返回一个区块的交易数目。
	count, err := client.TransactionCount(context.Background(), block.Hash())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(count) // 144
}
