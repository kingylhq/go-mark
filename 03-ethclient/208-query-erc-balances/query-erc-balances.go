package main

//func main() {
//	//client, err := ethclient.Dial("https://cloudflare-eth.com")
//	client, err := ethclient.Dial(constants.HTTPS_URL)
//	if err != nil {
//		log.Fatal(err)
//	}
//	// Golem (GNT) Address
//	tokenAddress := common.HexToAddress("0x5355d939fe772bA37a979058AAeC786C2b8ef34d")
//	instance, err := token.NewToken(tokenAddress, client)
//	if err != nil {
//		log.Fatal(err)
//	}
//	address := common.HexToAddress(constants.MateMaskAccount1Address)
//	bal, err := instance.BalanceOf(&bind.CallOpts{}, address)
//	if err != nil {
//		log.Fatal(err)
//	}
//	name, err := instance.Name(&bind.CallOpts{})
//	if err != nil {
//		log.Fatal(err)
//	}
//	symbol, err := instance.Symbol(&bind.CallOpts{})
//	if err != nil {
//		log.Fatal(err)
//	}
//	decimals, err := instance.Decimals(&bind.CallOpts{})
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Printf("name: %s\n", name)         // "name: Golem Network"
//	fmt.Printf("symbol: %s\n", symbol)     // "symbol: GNT"
//	fmt.Printf("decimals: %v\n", decimals) // "decimals: 18"
//	fmt.Printf("wei: %s\n", bal)           // "wei: 74605500647408739782407023"
//	fbal := new(big.Float)
//	fbal.SetString(bal.String())
//	value := new(big.Float).Quo(fbal, big.NewFloat(math.Pow10(int(decimals))))
//	fmt.Printf("balance: %f", value) // "balance: 74605500.647409"
//
//}
