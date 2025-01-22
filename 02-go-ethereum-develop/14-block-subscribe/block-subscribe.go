package main

import (
	"context"
	"fmt"
	"go-mark/02-go-ethereum-develop/constants"
	"log"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

// 订阅新区块
func main() {

	// 设置订阅以便在新区块被开采时获取事件。首先，我们需要一个支持 websocket RPC 的以太坊服务提供者。
	// 在示例中，我们将使用 infura 的 websocket 端点
	client, err := ethclient.Dial(constants.WSS_URL) // 自己的应用及项目id
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	if err != nil {
		log.Fatal(err)
	}

	// 创建一个新的通道，用于接收最新的区块头。
	headers := make(chan *types.Header)
	// 调用客户端的 SubscribeNewHead 方法，它接收我们刚创建的区块头通道，该方法将返回一个订阅对象。
	sub, err := client.SubscribeNewHead(context.Background(), headers)
	if err != nil {
		log.Fatal(err)
	}

	// 订阅将推送新的区块头事件到我们的通道，因此我们可以使用一个 select 语句来监听新消息。订阅对象还包括一个 error 通道，该通道将在订阅失败时发送消息。
	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case header := <-headers:
			fmt.Println("区块头hash: ", header.Hash().Hex()) // 0x98dc8117a60f98cd1a1e0fba7698be3c1d85eb8c5955e3a85f345f412ec9c082
			// 要获得该区块的完整内容，我们可以将区块头的摘要传递给客户端的 BlockByHash 函数。
			block, err := client.BlockByHash(context.Background(), header.Hash())
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println("区块的hash：", block.Hash().Hex())   // 0x98dc8117a60f98cd1a1e0fba7698be3c1d85eb8c5955e3a85f345f412ec9c082
			fmt.Println("区块号: ", block.Number().Uint64()) // 7440317
			//fmt.Println(block.Time().Uint64())     // 1529525947
			fmt.Println("区块时间: ", block.Time())                // 1736258064
			fmt.Println("区块nonce: ", block.Nonce())            // 0
			fmt.Println("区块交易长度: ", len(block.Transactions())) // 7
			fmt.Println("******************************************************")
		}
	}
}
