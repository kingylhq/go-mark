package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"go-mark/02-go-ethereum-develop/constants"
	ConvertUtils "go-mark/03-ethclient/utils"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/ethereum/go-ethereum/ethclient"
	//"github.com/ethereum/go-ethereum/crypto/sha3" // 已经弃用
	"golang.org/x/crypto/sha3"
)

// 代币转账
func main() {
	// 自己的节点
	client, err := ethclient.Dial(constants.HTTPS_URL)
	if err != nil {
		log.Fatal(err)
	}

	//privateKey, err := crypto.HexToECDSA("fad9c8855b740a0b7ed4c221dbad0f33a83a49cad6b3fe8d5817ac83d38b6a19")
	privateKey, err := crypto.HexToECDSA(constants.MateMaskPrivateKey)
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	// 这个方法确保你的新交易不会与同一地址的现有交易产生 nonce 冲突。
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	// 建议 Gas 价格: 返回当前网络中建议的每单位 gas 的价格。这个建议的 gas 价格基于网络的当前拥堵情况，帮助你在合理的时间内确认交易。
	// 节省成本: 通过使用推荐的 gas 价格，可以避免支付过高的费用，也可以减少因 gas 价格太低而导致交易延迟或失败的风险
	value := ConvertUtils.Convert(0.0001) // value := big.NewInt(0) // in wei (0 eth)
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	// 这是一个接收代币的目标地址。
	toAddress := common.HexToAddress("0x4592d8f8d7b001e72cb26a73e4fa1806a51ac79d")
	// 这是代币合约的地址。
	tokenAddress := common.HexToAddress("0x28b149020d2152179873ec60bed6bf7cd705775d")

	transferFnSignature := []byte("transfer(address,uint256)")
	hash := sha3.New256()
	hash.Write(transferFnSignature)
	methodID := hash.Sum(nil)[:4]
	fmt.Println(hexutil.Encode(methodID)) // 0xa9059cbb

	// 需要将给我们发送代币的地址左填充到 32 字节。
	paddedAddress := common.LeftPadBytes(toAddress.Bytes(), 32)
	fmt.Println(hexutil.Encode(paddedAddress)) // 0x0000000000000000000000004592d8f8d7b001e72cb26a73e4fa1806a51ac79d

	amount := new(big.Int)
	amount.SetString("1000000000000000000000", 10) // 1000 tokens
	// 代币量也需要左填充到 32 个字节。
	paddedAmount := common.LeftPadBytes(amount.Bytes(), 32)
	fmt.Println(hexutil.Encode(paddedAmount)) // 0x00000000000000000000000000000000000000000000003635c9adc5dea00000

	var data []byte
	data = append(data, methodID...)
	data = append(data, paddedAddress...)
	data = append(data, paddedAmount...)

	// 估算 Gas: 根据交易目标地址和输入数据，估算执行该交易需要的 gas 量。这可以帮助你设置合适的 gasLimit，确保交易能够顺利执行而不浪费太多 gas。
	// 模拟执行: 它会在不提交交易的情况下模拟执行，返回一个合理的 gas 估算值。
	// 防止失败: 通过提前估算 gas，可以避免因设置的 gasLimit 太低而导致交易失败。
	gasLimit, err := client.EstimateGas(context.Background(), ethereum.CallMsg{
		To:   &toAddress,
		Data: data,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(gasLimit) // 23256

	//可以提高gas limit， 创建交易对象, NewTransaction 已过期，用下面的方法 NewTx
	//tx := types.NewTransaction(nonce, tokenAddress, value, gasLimit, gasPrice, data)
	tx := types.NewTx(&types.LegacyTx{
		Nonce:    nonce,
		To:       &tokenAddress,
		Value:    value,
		Gas:      gasLimit,
		GasPrice: gasPrice,
		Data:     data,
	})

	// 它返回与当前以太坊客户端连接的网络的链 ID。这对于识别网络非常重要，尤其是在处理不同网络时（例如以太坊主网、Ropsten 测试网、Goerli 测试网等）。
	// 交易签名: 在以太坊中，链 ID 用于防止交易重放攻击。在签名交易时，你需要知道链 ID 以便正确签署交易。
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	// EIP-155 签名的作用
	//防止重放攻击: 在早期版本的以太坊中，交易签名中没有包含链 ID 的信息，这导致了一个问题：同一笔交易可以被复制到不同的网络（例如从测试网到主网），
	//这就是所谓的重放攻击。EIP-155 通过在签名中加入链 ID，解决了这个问题。每个网络的链 ID 是不同的，因此即使交易在一个网络上有效，在另一个网络上也会无效。
	//签名方式: 使用 EIP-155 签名会在交易数据中附加链 ID，从而改变交易的哈希值，使得该哈希值特定于某个链。
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s", signedTx.Hash().Hex()) // tx sent: 0xa56316b637a94c4cc0331c73ef26389d6c097506d581073f927275e7a6ece0bc
}
