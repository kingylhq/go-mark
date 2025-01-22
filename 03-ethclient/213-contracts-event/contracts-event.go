package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"go-mark/02-go-ethereum-develop/constants"
	"log"
	"math/big"
	"strings"
)

//2.13 合约事件
//智能合约具有在执行期间“发出”事件的能力，事件在以太坊中也称为“日志”。
//在以太坊智能合约中，事件被广泛使用，方便在发生相对重要的动作时记录、通知，特别是在代币合约（即 ERC-20）中，以指示代币合约发生转账操作。
//2.13.0 准备
//需要提前编译部署合约，store.sol 合约代码：

//pragma solidity ^0.8.26;
//contract Store {
//	event ItemSet(bytes32 indexed key, bytes32 value);
//	string public version;
//	mapping (bytes32 => bytes32) public items;
//	constructor(string memory _version) {
//	version = _version;
//	}
//	function setItem(bytes32 key, bytes32 value) external {
//		items[key] = value;
//		emit ItemSet(key, value);
//	}
//}

func main() {

	//QueryEvent()

	SubscribeEvent()

}

// 2.13.1 查询事件
func QueryEvent() {

	var StoreABI = `[{"inputs":[{"internalType":"string","name":"_version","type":"string"}],"stateMutability":"nonpayable","type":"constructor"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"bytes32","name":"key","type":"bytes32"},{"indexed":false,"internalType":"bytes32","name":"value","type":"bytes32"}],"name":"ItemSet","type":"event"},{"inputs":[{"internalType":"bytes32","name":"","type":"bytes32"}],"name":"items","outputs":[{"internalType":"bytes32","name":"","type":"bytes32"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"bytes32","name":"key","type":"bytes32"},{"internalType":"bytes32","name":"value","type":"bytes32"}],"name":"setItem","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[],"name":"version","outputs":[{"internalType":"string","name":"","type":"string"}],"stateMutability":"view","type":"function"}]`

	client, err := ethclient.Dial(constants.HTTPS_URL)
	if err != nil {
		log.Fatal(err)
	}

	// 智能合约可以可选地释放“事件”，其作为交易收据的一部分存储日志。读取这些事件相当简单。首先我们需要构造一个过滤查询。
	// 我们从 go-ethereum 包中导入 FilterQuery 结构体并用过滤选项初始化它。我们告诉它我们想过滤的区块范围并指定从中读取此日志的合约地址。
	// 在示例中，我们将从在智能合约章节创建的智能合约中读取特定区块所有日志。
	contractAddress := common.HexToAddress(constants.ContractsAddress)
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(6920583),
		// ToBlock:   big.NewInt(2394201),
		Addresses: []common.Address{
			contractAddress,
		},
		// Topics: [][]common.Hash{
		//  {},
		//  {},
		// },
	}

	// 下一步是调用 ethclient 的 FilterLogs，它接收我们的查询并将返回所有的匹配事件日志。
	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}

	contractAbi, err := abi.JSON(strings.NewReader(StoreABI))
	if err != nil {
		log.Fatal(err)
	}

	for _, vLog := range logs {
		fmt.Println(vLog.BlockHash.Hex())
		fmt.Println(vLog.BlockNumber)
		fmt.Println(vLog.TxHash.Hex())
		event := struct {
			Key   [32]byte
			Value [32]byte
		}{}
		err := contractAbi.UnpackIntoInterface(&event, "ItemSet", vLog.Data)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(common.Bytes2Hex(event.Key[:]))
		fmt.Println(common.Bytes2Hex(event.Value[:]))

		//主题(Topics)
		//事件最多可以有 4 个 topic。
		//topics[0], 第一个 topic 是固定的，必定存在，是按照 <EventName>([EventFieldType...]) 的模式生成事件签名，并取哈希后得到。
		//另外三个是可索引 topic，可索引 topic 是被 indexed 关键字修饰的事件字段的值。
		//每有一个事件字段被 indexed 关键字修饰，就会多一个 topic 值。
		//并且字段被索引之后，这个值就不会再被记录到事件的 Data 字段中。
		//在 solidity 中声明事件时，在类型与参数名称之间添加 indexed 关键字，来标记可索引 topic：
		var topics []string
		for i := range vLog.Topics {
			topics = append(topics, vLog.Topics[i].Hex())
		}

		// 第一个主题总是事件的签名。我们的示例合约不包含可索引的事件，但如果它确实包含，这是如何读取事件主题。
		fmt.Println("topics[0]=", topics[0])
		if len(topics) > 1 {
			fmt.Println("indexed topics:", topics[1:])
		}
	}

	// 首个主题只是被哈希过的事件签名。
	eventSignature := []byte("ItemSet(bytes32,bytes32)")
	hash := crypto.Keccak256Hash(eventSignature)
	fmt.Println("signature topics=", hash.Hex())
}

// 2.13.2 订阅事件
// 订阅事件日志，和订阅区块一样，需要 websocket RPC URL。
func SubscribeEvent() {
	var StoreABI = `[{"inputs":[{"internalType":"string","name":"_version","type":"string"}],"stateMutability":"nonpayable","type":"constructor"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"bytes32","name":"key","type":"bytes32"},{"indexed":false,"internalType":"bytes32","name":"value","type":"bytes32"}],"name":"ItemSet","type":"event"},{"inputs":[{"internalType":"bytes32","name":"","type":"bytes32"}],"name":"items","outputs":[{"internalType":"bytes32","name":"","type":"bytes32"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"bytes32","name":"key","type":"bytes32"},{"internalType":"bytes32","name":"value","type":"bytes32"}],"name":"setItem","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[],"name":"version","outputs":[{"internalType":"string","name":"","type":"string"}],"stateMutability":"view","type":"function"}]`

	// 订阅事件日志，和订阅区块一样，需要 websocket RPC URL。
	//client, err := ethclient.Dial("wss://rinkeby.infura.io/ws")
	client, err := ethclient.Dial(constants.WSS_URL)
	if err != nil {
		log.Fatal(err)
	}

	// 接收事件的方式是通过 channel。 从 go-ethereum core/types 包创建一个类型为 Log 的 channel。
	contractAddress := common.HexToAddress(constants.ContractsAddress)
	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
	}
	// 接收事件的方式是通过 channel。 从 go-ethereum core/types 包创建一个类型为 Log 的 channel。
	logs := make(chan types.Log)

	// 通过从客户端调用 SubscribeFilterLogs 方法来订阅，它接收查询选项和输出通道。 这将返回包含 unsubscribe 和 error 方法的订阅结构。
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal(err)
	}
	contractAbi, err := abi.JSON(strings.NewReader(StoreABI))
	if err != nil {
		log.Fatal(err)
	}

	// 最后，我们要做的就是使用 select 语句设置一个连续循环来读入新的日志事件或订阅错误。
	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case vLog := <-logs:
			fmt.Println(vLog.BlockHash.Hex())
			fmt.Println(vLog.BlockNumber)
			fmt.Println(vLog.TxHash.Hex())
			event := struct {
				Key   [32]byte
				Value [32]byte
			}{}
			err := contractAbi.UnpackIntoInterface(&event, "ItemSet", vLog.Data)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(common.Bytes2Hex(event.Key[:]))
			fmt.Println(common.Bytes2Hex(event.Value[:]))
			var topics []string
			for i := range vLog.Topics {
				topics = append(topics, vLog.Topics[i].Hex())
			}
			fmt.Println("topics[0]=", topics[0])
			if len(topics) > 1 {
				fmt.Println("index topic:", topics[1:])
			}
		}
	}
	//补充：
	//除了从查询事件和订阅事件能够获得合约事件，还可以从交易收据（TransactionReceipt）的 Logs 字段获取合约事件数据。
}
