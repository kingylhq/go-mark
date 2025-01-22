package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"go-mark/02-go-ethereum-develop/constants"
	ConvertUtils "go-mark/03-ethclient/utils"
	"log"
)

// 转账以太币 ETH
// 在本课程中，您将学习如何将 ETH 从一个帐户转移到另一个帐户。如果您已熟悉以太坊，那么您就知道如何交易包括您打算转账的以太币数量量，燃气限额，
// 燃气价格，一个随机数(nonce)，接收地址以及可选择性的添加的数据。 在广告发送到网络之前，必须使用发送方的私钥对该交易进行签名。
func main() {

	// tx sent: 0x57796e0026714c4c77b43fbddfea737c5eb8b4ad9b93cd9948ec7d790fdb3d9a
	// tx sent: 0xc4cb7128473e6ca0ab638b570e4b31d36986ac102a1617710c3a523a875d46e5

	client, err := ethclient.Dial(constants.HTTPS_URL)
	if err != nil {
		log.Fatal(err)
	}

	// 加载您的私钥。
	//privateKey, err := crypto.HexToECDSA("fad9c8855b740a0b7ed4c221dbad0f33a83a49cad6b3fe8d5817ac83d38b6a19")
	privateKey, err := crypto.HexToECDSA(constants.MateMaskPrivateKey)
	if err != nil {
		log.Fatal(err)
	}

	//之后我们需要获得帐户的随机数(nonce)。 每笔交易都需要一个 nonce。 根据定义，nonce 是仅使用一次的数字。 如果是发送交易的新帐户，
	//则该随机数将为“0”。 来自帐户的每个新事务都必须具有前一个 nonce 增加 1 的 nonce。很难对所有 nonce 进行手动跟踪，
	//于是 ethereum 客户端提供一个帮助方法 PendingNonceAt，它将返回你应该使用的下一个 nonce。
	//该函数需要我们发送的帐户的公共地址 - 这个我们可以从私钥派生。
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	// 返回你应该使用的下一个 nonce
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	// 设置我们将要转移的 ETH 数量。 但是我们必须将 ETH 以太转换为 wei，因为这是以太坊区块链所使用的。 以太网支持最多 18 个小数位，
	// 因此 1 个 ETH 为 1 加 18 个零。 这里有一个小工具可以帮助您在 ETH 和 wei 之间进行转换: https://etherconverter.netlify.com
	//value := big.NewInt(1000000000000000000) // in wei (1 eth)
	value := ConvertUtils.Convert(0.0001) // in wei (1 eth)
	//ETH 转账的燃气应设上限为“21000”单位。
	gasLimit := uint64(21000) // in units
	//然而，燃气价格总是根据市场需求和用户愿意支付的价格而波动的，因此对燃气价格进行硬编码有时并不理想。 go-ethereum 客户端提供
	// SuggestGasPrice 函数，用于根据'x'个先前块来获得平均燃气价格。
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	// 将 ETH 发送给谁
	//toAddress := common.HexToAddress("0x4592d8f8d7b001e72cb26a73e4fa1806a51ac79d")
	toAddress := common.HexToAddress(constants.MateMaskAccount2Address)
	var data []byte
	// 调用 NewTransaction 来生成我们的未签名以太坊事务，这个函数需要接收 nonce，地址，值，燃气上限值，燃气价格和可选发的数据。
	// 发送 ETH 的数据字段为“nil”。 在与智能合约进行交互时，我们将使用数据字段，仅仅转账以太币是不需要数据字段的。
	//tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)
	tx := types.NewTx(&types.LegacyTx{
		Nonce:    nonce,
		To:       &toAddress,
		Value:    value,
		Gas:      gasLimit,
		GasPrice: gasPrice,
		Data:     data,
	})

	// SignTx 方法需要 EIP155 签名者，这个也需要我们先从客户端拿到链 ID。
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	// 使用发件人的私钥对事务进行签名。 为此，我们调用 SignTx 方法，该方法接受一个未签名的事务和我们之前构造的私钥。
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	// 通过在客户端上调用“SendTransaction”来将已签名的事务广播到整个网络。
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s", signedTx.Hash().Hex())
}
