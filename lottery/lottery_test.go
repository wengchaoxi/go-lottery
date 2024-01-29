package lottery_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/wengchaoxi/go-lottery/lottery"
)

func TestLotteryRun(t *testing.T) {
	lot := lottery.NewLottery()
	lot.SetRandomEvent("This is random event")

	participants := []string{"p_0"}
	for i := 1; i < 10000000; i++ {
		participants = append(participants, fmt.Sprintf("p_%d", i))
	}

	start := time.Now()
	lot.Run(participants)
	elapsed := time.Since(start)
	fmt.Println("Cost Time:", elapsed)

	winner := lot.Run(participants)
	fmt.Printf("Winner ID: `%s`\n", winner)
}
