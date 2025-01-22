package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"go-mark/02-go-ethereum-develop/constants"
	"log"
)

// 2.9 订阅区块
func main() {

	// 订阅区块需要 websocket RPC URL。
	client, err := ethclient.Dial(constants.WSS_URL)
	if err != nil {
		log.Fatal(err)
	}

	// 接下来，我们将创建一个新的通道，用于接收最新的区块头。
	headers := make(chan *types.Header)

	//现在我们调用客户端的 SubscribeNewHead 方法，它接收我们刚创建的区块头通道，该方法将返回一个订阅对象。
	sub, err := client.SubscribeNewHead(context.Background(), headers)
	if err != nil {

		log.Fatal(err)
	}

	//订阅将推送新的区块头事件到我们的通道，因此我们可以使用一个 select 语句来监听新消息。订阅对象还包括一个 error 通道，该通道将在订阅失败时发送消息。

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case header := <-headers:
			fmt.Println("区块头hash值: ", header.Hash().Hex())
			fmt.Println("区块头号", header.Number.Uint64())
			fmt.Println("区块头父hash", header.ParentHash.Hex())
			//fmt.Println(header.Time)
			//fmt.Println(header.Difficulty.Uint64())
			//fmt.Println(header.GasLimit)
			//fmt.Println(header.GasUsed)
			//fmt.Println(header.Extra)
			//fmt.Println(header.MixDigest.Hex())
			//fmt.Println(header.Nonce)
			//fmt.Println(header.BaseFee.Uint64())
			//fmt.Println(header.UncleHash.Hex())
			//fmt.Println(header.Root.Hex())
			//fmt.Println(header.TxHash.Hex())
			//fmt.Println(header.Coinbase.Hex())
			//fmt.Println(header.ReceiptHash.Hex())
			//fmt.Println(header.Bloom.Big().Text(16))
			//fmt.Println(header.Size())
			//fmt.Println(header.Hash().Big().Text(16))
			//fmt.Println(header.ParentHash.Big().Text(16))
			//fmt.Println(header.UncleHash.Big().Text(16))
			//fmt.Println(header.TxHash.Big().Text(16))
			//fmt.Println(header.Root.Big().Text(16))
			//fmt.Println(header.ReceiptHash.Big().Text(16))
			block, err := client.BlockByHash(context.Background(), header.Hash())
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("区块hash: ", block.Hash().Hex())
			fmt.Println("区块号: ", block.Number().Uint64())
			fmt.Println("区块当前时间戳: ", block.Time())
			fmt.Println("区块Nonce: ", block.Nonce())
			fmt.Println("区块交易数: ", len(block.Transactions()))
			fmt.Println("*********************************")

		}
	}

}
