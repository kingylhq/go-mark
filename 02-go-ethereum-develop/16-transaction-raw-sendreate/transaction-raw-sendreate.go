package main

import (
	"context"
	"encoding/hex"
	"fmt"
	"go-mark/02-go-ethereum-develop/constants"
	"log"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rlp"
)

func main() {

	client, err := ethclient.Dial(constants.HTTPS_URL)
	if err != nil {
		log.Fatal(err)
	}
	// 在上个章节中我们学会了如何创建原始事务。 现在，我们将学习如何将其广播到以太坊网络，以便最终被处理和被矿工打包到区块。
	//首先将原始事务十六进制解码为字节格式。
	rawTx := "f86d1a84650c77ba8252089438be9eeebad7c2327fdee9bf412cb84a1ae1668186b5e620f48000808401546d72a0d459548995494ddff76f32263ceffda12a5f2191aa33cc9abe43513794afcdaba00f2b83d1f62548b0d57f9b9f5186f7c1434ad001363d4c68ef821a4abad"
	rawTxBytes, err := hex.DecodeString(rawTx)
	if err != nil {
		log.Fatal(err)
	}

	// 接下来初始化一个新的 types.Transaction 指针并从 go-ethereum rlp 包中调用 DecodeBytes，将原始事务字节和指针传递给以太坊事务类型。
	//RLP 是以太坊用于序列化和反序列化数据的编码方法。
	tx := new(types.Transaction)
	rlp.DecodeBytes(rawTxBytes, &tx)

	// 使用我们的以太坊客户端轻松地广播交易。
	err = client.SendTransaction(context.Background(), tx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("交易hash: %s", tx.Hash().Hex()) // tx sent: 0xc429e5f128387d224ba8bed6885e86525e14bfdc2eb24b5e9c3351a1176fd81f
}
