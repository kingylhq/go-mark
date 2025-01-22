package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	store "go-mark/02-go-ethereum-develop/18-contract-deploy/contracts"
	"go-mark/02-go-ethereum-develop/constants"
	"log"
	"math/big"
)

func main() {
	client, err := ethclient.Dial(constants.HTTPS_URL)
	if err != nil {
		log.Fatal(err)
	}

	// 验证合约是否存在
	//contractAddress := common.HexToAddress("0x509A282A95840Fb2097D9772E0407068d23f4bB6")
	//code, err := client.CodeAt(context.Background(), contractAddress, nil)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//if len(code) == 0 {
	//	log.Fatal("No contract code at address", contractAddress.Hex())
	//} else {
	//	fmt.Println("Contract code found at address", contractAddress.Hex())
	//}

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

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	// 这个是合约的地址
	address := common.HexToAddress("0x509A282A95840Fb2097D9772E0407068d23f4bB6")
	instance, err := store.NewStore(address, client)
	if err != nil {
		log.Fatal(err)
	}

	key := [32]byte{}
	value := [32]byte{}
	copy(key[:], []byte("foo"))
	copy(value[:], []byte("bar"))

	tx, err := instance.SetItem(auth, key, value)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s \n", tx.Hash().Hex()) // tx sent: 0xa5c0dbd082632ad0e70563a290c3650cc4b88ef4f7fa1a54d4fdd92ff3f82159

	result, err := instance.Items(nil, key)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(result[:])) // "bar"
}
