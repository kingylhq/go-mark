package main

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"go-mark/02-go-ethereum-develop/constants"
	store "go-mark/03-ethclient/210-deploy-contracts/contracts"
	ConvertUtils "go-mark/03-ethclient/utils"
	"log"
	"math/big"
	"time"
)

const (

	// store合约的字节码(Store.bin 文件内存)
	contractBytecode = "608060405234801561000f575f5ffd5b5060405161087838038061087883398181016040528101906100319190610193565b805f908161003f91906103ea565b50506104b9565b5f604051905090565b5f5ffd5b5f5ffd5b5f5ffd5b5f5ffd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b6100a58261005f565b810181811067ffffffffffffffff821117156100c4576100c361006f565b5b80604052505050565b5f6100d6610046565b90506100e2828261009c565b919050565b5f67ffffffffffffffff8211156101015761010061006f565b5b61010a8261005f565b9050602081019050919050565b8281835e5f83830152505050565b5f610137610132846100e7565b6100cd565b9050828152602081018484840111156101535761015261005b565b5b61015e848285610117565b509392505050565b5f82601f83011261017a57610179610057565b5b815161018a848260208601610125565b91505092915050565b5f602082840312156101a8576101a761004f565b5b5f82015167ffffffffffffffff8111156101c5576101c4610053565b5b6101d184828501610166565b91505092915050565b5f81519050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f600282049050600182168061022857607f821691505b60208210810361023b5761023a6101e4565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f6008830261029d7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82610262565b6102a78683610262565b95508019841693508086168417925050509392505050565b5f819050919050565b5f819050919050565b5f6102eb6102e66102e1846102bf565b6102c8565b6102bf565b9050919050565b5f819050919050565b610304836102d1565b610318610310826102f2565b84845461026e565b825550505050565b5f5f905090565b61032f610320565b61033a8184846102fb565b505050565b5b8181101561035d576103525f82610327565b600181019050610340565b5050565b601f8211156103a25761037381610241565b61037c84610253565b8101602085101561038b578190505b61039f61039785610253565b83018261033f565b50505b505050565b5f82821c905092915050565b5f6103c25f19846008026103a7565b1980831691505092915050565b5f6103da83836103b3565b9150826002028217905092915050565b6103f3826101da565b67ffffffffffffffff81111561040c5761040b61006f565b5b6104168254610211565b610421828285610361565b5f60209050601f831160018114610452575f8415610440578287015190505b61044a85826103cf565b8655506104b1565b601f19841661046086610241565b5f5b8281101561048757848901518255600182019150602085019450602081019050610462565b868310156104a457848901516104a0601f8916826103b3565b8355505b6001600288020188555050505b505050505050565b6103b2806104c65f395ff3fe608060405234801561000f575f5ffd5b506004361061003f575f3560e01c806348f343f31461004357806354fd4d5014610073578063f56256c714610091575b5f5ffd5b61005d600480360381019061005891906101d7565b6100ad565b60405161006a9190610211565b60405180910390f35b61007b6100c2565b604051610088919061029a565b60405180910390f35b6100ab60048036038101906100a691906102ba565b61014d565b005b6001602052805f5260405f205f915090505481565b5f80546100ce90610325565b80601f01602080910402602001604051908101604052809291908181526020018280546100fa90610325565b80156101455780601f1061011c57610100808354040283529160200191610145565b820191905f5260205f20905b81548152906001019060200180831161012857829003601f168201915b505050505081565b8060015f8481526020019081526020015f20819055507fe79e73da417710ae99aa2088575580a60415d359acfad9cdd3382d59c80281d48282604051610194929190610355565b60405180910390a15050565b5f5ffd5b5f819050919050565b6101b6816101a4565b81146101c0575f5ffd5b50565b5f813590506101d1816101ad565b92915050565b5f602082840312156101ec576101eb6101a0565b5b5f6101f9848285016101c3565b91505092915050565b61020b816101a4565b82525050565b5f6020820190506102245f830184610202565b92915050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f601f19601f8301169050919050565b5f61026c8261022a565b6102768185610234565b9350610286818560208601610244565b61028f81610252565b840191505092915050565b5f6020820190508181035f8301526102b28184610262565b905092915050565b5f5f604083850312156102d0576102cf6101a0565b5b5f6102dd858286016101c3565b92505060206102ee858286016101c3565b9150509250929050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f600282049050600182168061033c57607f821691505b60208210810361034f5761034e6102f8565b5b50919050565b5f6040820190506103685f830185610202565b6103756020830184610202565b939250505056fea26469706673582212207e5948071fb4b4ea28a688d9bef925d42a8b34d7364bf3c782f847b987ffd34164736f6c634300081c0033"
)

func main() {

	// 0.4714
	DeployStoreWay1()

	//DeployStoreWay2()

}

// 2.10.1 使用 abigen 工具
// 部署合约，方式一
func DeployStoreWay1() {

	//部署合约地址:  0x23ac3f64854B699b4F9470E17CEaB691f870C5eE
	//交易事务tx: 0x66f3a60a2ea913a43ac588eddc259e077cdcdcac0978f64a94e26ab641126c89
	client, err := ethclient.Dial(constants.HTTPS_URL)
	if err != nil {
		log.Fatal(err)
	}

	//privateKey, err := crypto.GenerateKey()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//privateKeyBytex := crypto.FromECDSA(privateKey)
	//privateKeyHex := hex.EncodeToString(privateKeyBytex)
	//publicKey := privateKey.Public()

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
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		log.Fatal(err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	//auth.Value = big.NewInt(0)
	auth.Value = ConvertUtils.Convert(0.0001)
	auth.GasPrice = gasPrice
	auth.GasLimit = uint64(300000) // in units

	input := "1.02"
	address, tx, instance, err := store.DeployStore(auth, client, input)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("部署合约地址: ", address.Hex())
	fmt.Printf("交易事务tx: %s\n", tx.Hash().Hex())
	_ = instance
}

// 2.10.2 仅使用 ethclient 工具
// 以太坊中，部署合约其实也是发起了一笔交易，并不是一定需要 abigen 工具生成 go 代码。
// 不是只能使用生成的 go 的合约代码才能部署合约。
// 可以仅使用 ethclient，使用更底层的方法，直接通过发送交易的方式来部署合约。
// 使用 2.10.0 步骤中，生成的 store_sol_Store.bin 文件中的字符串作为交易数据，部署 store 合约。
func DeployStoreWay2() {

	//交易发送成功 Transaction sent: 0xfcd74f12f0c1162df2873ace21c6a8a8c68bd22ca27c8a370ab71ff4b1f8b61c   部署失败，余额不够
	//交易被挖矿了 Transaction mined: %!s(uint64=0), Contract deployed at: 0x80F09e0F648D75FA4736987DAd6DAfc99092F6B6

	//insufficient funds for gas * price + value: balance 79289143594361328, tx cost 127606525380000000, overshot 48317381785638672
	// 太耗费eth

	//交易发送成功 Transaction sent: 0x305aa3d940ea40b706512f0dab077d3ec916e9f19233874107694c73e4b6a862
	//
	client, err := ethclient.Dial(constants.HTTPS_URL)
	if err != nil {
		log.Fatal(err)
	}

	// 私钥进行 椭圆曲线数字加密算法 16进制私钥
	privateKey, keyErr := crypto.HexToECDSA(constants.MateMaskPrivateKey)
	if keyErr != nil {
		log.Fatal(keyErr)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, nonceErr := client.PendingNonceAt(context.Background(), fromAddress)
	if nonceErr != nil {
		log.Fatal(nonceErr)
	}

	gasPrice, gasPriceErr := client.SuggestGasPrice(context.Background())
	if gasPriceErr != nil {
		log.Fatal(gasPriceErr)
	}

	// 解码合约字节码
	data, dataErr := hex.DecodeString(contractBytecode)
	if dataErr != nil {
		log.Fatal(dataErr)
	}

	amount := ConvertUtils.Convert(0.0001)
	// 创建交易
	tx := types.NewContractCreation(nonce, amount, 10000000, gasPrice, data)

	// 获取链ID
	chainID, chainIDErr := client.ChainID(context.Background())
	if chainIDErr != nil {
		log.Fatal(chainIDErr)
	}

	// 签名交易
	signedTx, signErr := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if signErr != nil {
		log.Fatal(signErr)
	}
	// 发送交易
	sendErr := client.SendTransaction(context.Background(), signedTx)
	if sendErr != nil {
		log.Fatal(sendErr)
	}
	fmt.Printf("交易发送成功 Transaction sent: %s\n", signedTx.Hash().Hex())

	// 等待被挖矿
	receipt, receiptErr := WaitMinedForReceipt(client, signedTx.Hash())
	if receiptErr != nil {
		log.Fatal(receiptErr)
	}
	fmt.Printf("交易被挖矿了 Contract deployed state: %d\n", receipt.Status)
	fmt.Printf("交易被挖矿了 Contract deployed at: %s\n", receipt.ContractAddress.Hex())

}

// 等待被挖矿
func WaitMinedForReceipt(client *ethclient.Client, txHash common.Hash) (*types.Receipt, error) {
	for {
		receipt, err := client.TransactionReceipt(context.Background(), txHash)
		if err == nil {
			return receipt, nil
		}
		if err != ethereum.NotFound {
			return nil, err
		}
		// 等待一段时间后再次查询
		time.Sleep(1 * time.Second)
	}
}
