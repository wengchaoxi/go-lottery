package lottery

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// Block represents the structure of a block response
type Block struct {
	Hash string `json:"hash"`
	Time int64  `json:"time"`
}

func BTCBlockInfo(timeStr, timeLayout string) string {
	targetTime, err := time.Parse(timeLayout, timeStr)
	if err != nil {
		panic(err)
	}

	// blockchain.info API endpoint for blocks around the target time
	url := fmt.Sprintf("https://blockchain.info/blocks/%d?format=json", targetTime.Unix()*1000)
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	var blocks []Block
	if err := json.Unmarshal(body, &blocks); err != nil {
		panic(err)
	}

	// Find the block that is closest to the target time
	var closestBlock Block
	smallestDiff := int64(1<<63 - 1) // Max int64
	for _, block := range blocks {
		diff := targetTime.Unix() - block.Time
		if diff >= 0 && diff < smallestDiff {
			smallestDiff = diff
			closestBlock = block
		}
	}

	// Return the closest block's hash and timestamp
	return fmt.Sprintf("BTC_%s_%d", closestBlock.Hash, closestBlock.Time)
}
