package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"go-mark/02-go-ethereum-develop/constants"
	"go-mark/03-ethclient/utils"
	"log"
	"math/big"

	"golang.org/x/crypto/sha3"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

// 创建一个ERC20 代币转账 交易
// 对应 remix contracts ---> 02 ---> MyToken 的合约，合约重新发布后需要修改对应的智能合约地址
func main() {
	// 0.0853 ETH
	// 0xb60595ea53a7813db3eed82392cea38ff4c7f4095c057d49d08c95a7174796ca  success
	// 0x6850a972df239fd6c5ae1ea98fd9fdb45974ca8c06e50d2d5d220a3d713254fd  pending - index - success
	// 0xc224000b8e86975ab5228377f2b5b1e254dd1dc439ebbd92cb5e0e43c8ce9718  success
	//client, err := ethclient.Dial("https://eth-sepolia.g.alchemy.com/v2/")
	client, err := ethclient.Dial(constants.HTTPS_URL)
	if err != nil {
		log.Fatal(err)
	}
	// 参数是账户的私钥，
	// 命令创建
	// geth account new
	// 执行这条命令后，会提示你输入密码，然后生成一个新的账户并显示其私钥。请记住将私钥保存在一个安全的地方。
	// 浏览器MetaMask 小狐狸创建的账户，如下是MateMask账户秘钥
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
	// 定义 0.0001 ETH 的值
	//ethValue := big.NewFloat(0.0001)
	//
	//// 1 ETH = 10^18 wei
	//weiPerEth := new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil)
	//
	//// 将 0.001 ETH 转换为 wei
	//valueInWei := new(big.Float).Mul(ethValue, new(big.Float).SetInt(weiPerEth))
	//
	//// 将结果转换为整数形式
	//value := big.NewInt(0) // in wei (0 eth)
	//valueInWei.Int(value)
	value := ConvertUtils.Convert(0.0002)
	fmt.Println("value金额: ", value)
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	//toAddress := common.HexToAddress("0x4592d8f8d7b001e72cb26a73e4fa1806a51ac79d")
	//tokenAddress := common.HexToAddress("0x28b149020d2152179873ec60bed6bf7cd705775d")
	// 要发送代币的地址存储在变量中。以太坊账户地址
	toAddress := common.HexToAddress("0x38bE9EeEBad7C2327fDee9bF412Cb84A1AE16681")
	// 将代币合约地址分配给变量。智能合约地址
	tokenAddress := common.HexToAddress("0x1ccb1BdDC8C876cA3e1C5Fd8c0045D9fE57CcDFE")

	transferFnSignature := []byte("transfer(address,uint256)")
	hash := sha3.NewLegacyKeccak256()
	hash.Write(transferFnSignature)
	methodID := hash.Sum(nil)[:4]
	fmt.Println("方法ID encode: ", hexutil.Encode(methodID)) // 0xa9059cbb
	paddedAddress := common.LeftPadBytes(toAddress.Bytes(), 32)
	fmt.Println("以太坊账户地址encode: ", hexutil.Encode(paddedAddress)) // 0x0000000000000000000000004592d8f8d7b001e72cb26a73e4fa1806a51ac79d
	amount := new(big.Int)
	fmt.Println("明文金额: ", amount)
	amount.SetString("1000000000000000000000", 10) // 1000 tokens
	paddedAmount := common.LeftPadBytes(amount.Bytes(), 32)
	fmt.Println("金额: ", hexutil.Encode(paddedAmount)) // 0x00000000000000000000000000000000000000000000003635c9adc5dea00000
	var data []byte
	data = append(data, methodID...)
	data = append(data, paddedAddress...)
	data = append(data, paddedAmount...)

	gasLimit, err := client.EstimateGas(context.Background(), ethereum.CallMsg{
		To:   &toAddress,
		Data: data,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("gas上限值: ", gasLimit) // 21810
	tx := types.NewTx(&types.LegacyTx{
		Nonce:    nonce,
		To:       &tokenAddress,
		Value:    value,
		Gas:      gasLimit + 100000,
		GasPrice: gasPrice,
		Data:     data,
	})
	// Sepolia测试网的网络ID是11155111。
	chainID, err := client.NetworkID(context.Background())
	fmt.Println("网络 链ID: ", chainID)
	if err != nil {
		log.Fatal(err)
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("交易发送hash签名16进制: %s", signedTx.Hash().Hex()) // tx sent: 0xa56316b637a94c4cc0331c73ef26389d6c097506d581073f927275e7a6ece0bc
}
