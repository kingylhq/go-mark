package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"go-mark/02-go-ethereum-develop/constants"
	"log"
)

func main() {

	ClientInit()

}

func ClientInit() {
	//client, err := ethclient.Dial("https://cloudflare-eth.com")
	client, err := ethclient.Dial(constants.HTTPS_URL)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("we have a connection")
	_ = client // we'll use this in the upcoming sections
}
