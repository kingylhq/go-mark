package main

import (
	"crypto/ecdsa"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"golang.org/x/crypto/sha3"
)

func main() {

	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}

	privateKeyBytes := crypto.FromECDSA(privateKey)
	fmt.Println("私钥key16进制: ", hexutil.Encode(privateKeyBytes)[2:]) // 0xfad9c8855b740a0b7ed4c221dbad0f33a83a49cad6b3fe8d5817ac83d38b6a19

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	fmt.Println("公钥key16进制去掉前四个字节: ", hexutil.Encode(publicKeyBytes)[4:]) // 0x049a7df67f79246283fdc93af76d4f8cdd62c4886e8cd870944e817dd0b97934fdd7719d0810951e03418205868a5c1b40b192451367f28e0088dd75e15de40c05

	// 使用 crypto.PubkeyToAddress 函数从给定的 ECDSA 公钥生成一个以太坊地址。.Hex() 方法将该地址转换为十六进制字符串。
	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	fmt.Println(address) // 0x96216849c49358B10257cb55b28eA603c874b05E

	// 创建了一个新的 Keccak-256 哈希对象，用于后续的哈希计算。sha3.NewLegacyKeccak256() 是来自以太坊库的一个函数，
	// 它返回一个实现了 hash.Hash 接口的对象
	hash := sha3.NewLegacyKeccak256()
	// 将 publicKeyBytes 数组从第二个字节开始的所有字节写入到哈希对象中。publicKeyBytes[1:] 表示从索引 1 开始（不包括第一个字节）的子切片
	hash.Write(publicKeyBytes[1:])
	//首先调用 hash.Sum(nil) 来计算哈希值，并返回一个字节切片。然后，通过 [12:] 截取这个字节切片从第 13 个字节开始的
	//部分（注意，Go 语言中的索引是从 0 开始的）。最后，使用 hexutil.Encode 将截取后的字节切片编码为十六进制字符串
	fmt.Println(hexutil.Encode(hash.Sum(nil)[12:])) // 0x96216849c49358b10257cb55b28ea603c874b05e
}
