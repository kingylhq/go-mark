package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	store "go-mark/02-go-ethereum-develop/18-contract-deploy/contracts"
	"go-mark/02-go-ethereum-develop/constants"
)

const (
	// 合约地址
	contractAddress = "0x23ac3f64854B699b4F9470E17CEaB691f870C5eE"
)

func main() {

	loadContracts()

}

func loadContracts() {

	client, err := ethclient.Dial(constants.HTTPS_URL)
	if err != nil {
		panic(err)
	}

	storeContract, storeErr := store.NewStore(common.HexToAddress(contractAddress), client)
	if storeErr != nil {
		panic(storeErr)
	}

	fmt.Println("storeContract:", storeContract)

	_ = storeContract

}
