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

	store "go-mark/02-go-ethereum-develop/18-contract-deploy/contracts"
)

func main() {

	// 0x5BB4e8eb1837bE837cAf48800fe7B13A08DF74A0
	// 0x70e7fdba655ee562ab067c8162c73554d49c05dd0622aca717f4e9a9b294fd96

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
	input := "2.0"
	address, tx, instance, err := store.DeployStore(auth, client, input)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(address.Hex())   // 0xfaCb6F19cfb7Bac5Cd3a559161f9068de2CF0B1E
	fmt.Println(tx.Hash().Hex()) // 0x491fc341a281aa5d9815238c86e891edf9f46264d5d21c695f47ad1bf8b37992

	_ = instance
}
