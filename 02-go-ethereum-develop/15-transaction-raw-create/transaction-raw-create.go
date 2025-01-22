package main

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rlp"
	"go-mark/02-go-ethereum-develop/constants"
	ConvertUtils "go-mark/03-ethclient/utils"
	"log"
)

// 构建原始交易（Raw Transaction）
func main() {

	// 签名hash: 0x80b7a4db8311c2a432142ab48906cdc418d1f001d4e5c02afd80a2874ce0e4c6  pending

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
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	//value := big.NewInt(1000000000000000000) // in wei (1 eth)
	value := ConvertUtils.Convert(0.0002)
	gasLimit := uint64(21000) // in units
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	//toAddress := common.HexToAddress("0x4592d8f8d7b001e72cb26a73e4fa1806a51ac79d")
	toAddress := common.HexToAddress(constants.MateMaskAccount2Address)
	var data []byte
	// 看过上个章节, 那么你知道如何加载你的私钥来签名交易。 我们现在假设你知道如何做到这一点，现在你想让原始交易数据能够在以后广播它。
	//首先构造事务对象并对其进行签名，例如：
	tx := types.NewTx(&types.LegacyTx{
		Nonce:    nonce,
		To:       &toAddress,
		Value:    value,
		Gas:      gasLimit,
		GasPrice: gasPrice,
		Data:     data,
	})

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	// chainID == 11155111 是测试网络 Sepolia
	fmt.Println("链ID: ", chainID)

	// signedTx.Hash().Hex() == Transaction Hash  (区块链浏览器查询的交易hash值)
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	// 只是创建了一个包含已签名交易的 types.Transactions 切片，但并不会发送交易。要发送交易，你需要使用以太坊客户端的 SendTransaction 方法。
	ts := types.Transactions{signedTx}
	rawTxBytes, err := rlp.EncodeToBytes(ts[0])
	fmt.Println("rawTxBytes: ", rawTxBytes)

	if err != nil {
		log.Fatal(err)
	}
	rawTxHex := hex.EncodeToString(rawTxBytes)
	fmt.Println("原始交易码Encoding: ", rawTxHex) // f86...772

	decodeBytes, err := hex.DecodeString(rawTxHex)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("decodeBytes: ", decodeBytes)
	// 上面是构建一个原始交易，下面的是做测试交易数据
	//tranErr := client.SendTransaction(context.Background(), signedTx)
	//if tranErr != nil {
	//	log.Fatal(tranErr)
	//}
	//fmt.Printf("签名hash: %s", signedTx.Hash().Hex())

	// 广播到测试网络
	// 接下来初始化一个新的 types.Transaction 指针并从 go-ethereum rlp 包中调用 DecodeBytes，将原始事务字节和指针传递给以太坊事务类型。
	//RLP 是以太坊用于序列化和反序列化数据的编码方法。
	sendTx := new(types.Transaction)
	rlp.DecodeBytes(rawTxBytes, &sendTx)

	// 使用我们的以太坊客户端轻松地广播交易。
	err = client.SendTransaction(context.Background(), sendTx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("交易hash: %s", sendTx.Hash().Hex())
}
