package lottery

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/big"
	"sync"
)

type LotteryParticipant struct {
	ID   string
	Hash string
	Diff *big.Int
}

type Lottery struct {
	RandomEvent     string
	RandomEventHash string

	Winner *LotteryParticipant
}

func NewLottery() *Lottery {
	return &Lottery{
		Winner: &LotteryParticipant{},
	}
}

func (l *Lottery) SetRandomEvent(randomEvent string) {
	l.RandomEvent = randomEvent
	l.RandomEventHash = l.Hash(randomEvent)
}

func (l *Lottery) Hash(val string) string {
	hash := sha256.Sum256([]byte(val))
	return hex.EncodeToString(hash[:])
}

func (l *Lottery) HashDiff(otherHash string) *big.Int {
	first := new(big.Int)
	second := new(big.Int)
	first.SetString(l.RandomEventHash, 16)
	second.SetString(otherHash, 16)
	diff := new(big.Int)
	diff.Sub(first, second)
	return diff.Abs(diff)
}

func (l *Lottery) Run(participants []string) string {
	var wg sync.WaitGroup
	var mu sync.Mutex
	var potentialWinners []LotteryParticipant

	tmp := new(big.Int).SetUint64(1)
	tmp = tmp.Lsh(tmp, 256)
	l.Winner.Diff = tmp.Sub(tmp, big.NewInt(1))

	batchSize := 1000
	batchCount := len(participants) / batchSize
	if len(participants)%batchSize != 0 {
		batchCount++
	}

	for i := 0; i < batchCount; i++ {
		start := i * batchSize
		end := start + batchSize
		if end > len(participants) {
			end = len(participants)
		}
		wg.Add(1)
		go func(participantsBatch []string) {
			defer wg.Done()

			localWinner := LotteryParticipant{
				Diff: tmp.Sub(tmp, big.NewInt(1)),
			}
			for _, p := range participantsBatch {
				participantHash := l.Hash(p)
				diff := l.HashDiff(participantHash)
				if diff.Cmp(localWinner.Diff) == -1 {
					localWinner.ID = p
					localWinner.Hash = participantHash
					localWinner.Diff = diff
				}
			}

			mu.Lock()
			switch localWinner.Diff.Cmp(l.Winner.Diff) {
			case -1:
				l.Winner = &localWinner
				potentialWinners = []LotteryParticipant{localWinner}
			case 0:
				potentialWinners = append(potentialWinners, localWinner)
			}
			mu.Unlock()
		}(participants[start:end])
	}
	wg.Wait()

	if len(potentialWinners) > 1 {
		l.Winner = &potentialWinners[0]
		for _, potentialWinner := range potentialWinners[1:] {
			if potentialWinner.Diff.Cmp(l.Winner.Diff) == -1 {
				l.Winner = &potentialWinner
			}
		}
	}
	return l.Winner.ID
}

func (l *Lottery) PrintWinner() {
	fmt.Printf("Random Event: `%s`\n", l.RandomEvent)
	fmt.Printf("Random Event Hash: `%s`\n", l.RandomEventHash)
	fmt.Printf("Winner ID: `%s`\n", l.Winner.ID)
	fmt.Printf("Winner ID Hash: `%s`\n", l.Winner.Hash)
}
