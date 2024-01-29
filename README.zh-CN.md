# Go Lottery

**简体中文** | [English](./README.md)

## 概述

通过对所有参与者的唯一标识符进行哈希计算，并与一个不可预测、公开可查的事件（比如 BTC 未来某个时间点的区块信息的 Hash）的哈希值进行对比，来确定哪个参与者的哈希值与之最接近，从而确定中奖者。

优势：

- 真正的随机：中奖结果与一个不可预测的事件（比如 BTC 未来某个时间点的区块信息的 Hash）绑定。

- 公开可验证：由于随机事件（比如 BTC 未来某个时间点的区块信息的 Hash）是公开可查的，结果的公正性可以被公众监督和验证。

## 使用

```go
package main

import "github.com/wengchaoxi/go-lottery/lottery"

func main() {
	lot := lottery.NewLottery()

	timeStr := "2024/01/01-00:00:00"
	randomEvent := lottery.BTCBlockInfo(timeStr, "2006/01/02-15:04:05")
	lot.SetRandomEvent(randomEvent)

	participants := []string{"Alice", "Bob", "Charlie", "Dave"}
	lot.Run(participants)
	lot.PrintWinner()
}
```

## 由来

这个程序的灵感来自于 [peng-zhihui/BilibiliLottery](https://github.com/peng-zhihui/BilibiliLottery)
