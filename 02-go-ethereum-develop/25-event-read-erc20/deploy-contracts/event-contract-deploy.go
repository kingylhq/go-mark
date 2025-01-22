package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"go-mark/02-go-ethereum-develop/constants"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	store "go-mark/02-go-ethereum-develop/25-event-read-erc20/contracts"
)

func main() {

	//0xB881F62387cEcFD550e6A8e8ec725c9cA1F607C9
	//0x699ef2cdca9b98301ced60743d8d3226e853eadda6e8337abfdcabe62b729572
	client, err := ethclient.Dial(constants.HTTPS_URL)
	if err != nil {
		log.Fatal(err)
	}

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

	// 预估费用
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	// 获取链ID
	chainID, chainErr := client.ChainID(context.Background())
	if chainErr != nil {
		log.Fatal(chainErr)
	}

	// bind.NewKeyedTransactor(privateKey) 已过期，使用NewKeyedTransactorWithChainID平替
	auth, authErr := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if authErr != nil {
		log.Fatal(authErr)
	}

	auth.Nonce = big.NewInt(int64(nonce))
	//auth.Value = ConvertUtils.Convert(0.0001) //
	auth.Value = big.NewInt(0) // in wei
	auth.GasLimit = uint64(30000000)
	//auth.GasPrice = gasPrice
	auth.GasPrice = new(big.Int).Mul(gasPrice, big.NewInt(2))

	// 部署合约
	address, tx, instance, err := store.DeployErc20(auth, client)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(address.Hex())   //
	fmt.Println(tx.Hash().Hex()) //

	_ = instance
}
