package main

import (
	"fmt"
	"testing"
)

func TestEmptyBalance(t *testing.T) {
	wallet := &JobcoinAPI{"http://jobcoin.gemini.com/sanitary/api/"}

	fmt.Println(wallet.CheckBalance("Alice"))
}
