package lottery_test

import (
	"fmt"
	"testing"

	"github.com/wengchaoxi/go-lottery/lottery"
)

func TestRandomEventBlockchainInfo(t *testing.T) {
	info := lottery.BTCBlockInfo("2024/01/01-00:00:00", "2006/01/02-15:04:05")
	fmt.Println(info)
}
