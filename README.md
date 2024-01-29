# Go Lottery

**English** | [简体中文](./README.zh-CN.md)

## Overview

Go Lottery determines winners by hashing the unique identifiers of all participants and comparing them to the hash of an unpredictable and publicly verifiable event (such as the hash of a BTC block at a future point in time). The participant whose hash value is closest to that of the event is declared the winner.

Advantages:

- True Randomness: The lottery results are tied to an unpredictable future event (e.g., the hash of a future BTC block), ensuring true randomness.

- Public Verification: As the random event (e.g., the hash of a future BTC block) is publicly available, the fairness of the results can be monitored and verified by the public.

## Usage

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

## Attribution

This program design was originally inspired from the work at [peng-zhihui/BilibiliLottery](https://github.com/peng-zhihui/BilibiliLottery)
