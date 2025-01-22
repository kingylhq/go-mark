package ConvertUtils

import (
	"math/big"
)

func Convert(eth float64) *big.Int {
	// 定义 0.0001 ETH 的值
	ethValue := big.NewFloat(eth)

	// 1 ETH = 10^18 wei
	weiPerEth := new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil)

	// 将 0.001 ETH 转换为 wei
	valueInWei := new(big.Float).Mul(ethValue, new(big.Float).SetInt(weiPerEth))

	// 将结果转换为整数形式
	value := big.NewInt(0) // in wei (0 eth)
	valueInWei.Int(value)
	return value
}
