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
