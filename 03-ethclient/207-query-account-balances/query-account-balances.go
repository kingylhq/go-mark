package main

import (
	"context"
	"fmt"
	"go-mark/02-go-ethereum-develop/constants"
	"log"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// 2.7 查询账户余额
func main() {
	client, err := ethclient.Dial(constants.HTTPS_URL)
	if err != nil {
		log.Fatal(err)
	}

	//调用 ethclient 的 BalanceAt 方法，给它传递账户地址和可选的区块号。将区块号设置为 nil 将返回最新的余额。
	//account := common.HexToAddress("0x25836239F7b632635F815689389C537133248edb")
	// 以太坊钱包账号地址
	account := common.HexToAddress(constants.MateMaskAccount1Address)
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(balance)

	// 传区块高度能读取指定区块时的账户余额，区块高度必须是 big.Int 类型。
	blockNumber := big.NewInt(5532993)
	balanceAt, err := client.BalanceAt(context.Background(), account, blockNumber)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(balanceAt) // 25729324269165216042
	// 以太坊中的数字是使用尽可能小的单位来处理的，因为它们是定点精度，在 ETH 中它是_wei_。要读取 ETH 值，您必须做计算 wei/10^18。
	// 因为我们正在处理大数，我们得导入原生的 Go math 和 math/big 包。这是您做的转换。
	fbalance := new(big.Float)
	fbalance.SetString(balanceAt.String())
	ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))
	fmt.Println(ethValue) // 25.729324269165216041

	// 待处理的余额
	//有时您想知道待处理的账户余额是多少，例如，在提交或等待交易确认后。客户端提供了类似 BalanceAt 的方法，名为 PendingBalanceAt，它接收账户地址作为参数。
	pendingBalance, err := client.PendingBalanceAt(context.Background(), account)
	fmt.Println(pendingBalance) // 25729324269165216042
}
