package main

import (
	"testing"
)

func TestEmptyBalance(t *testing.T) {
	wallet := &JobcoinAPI{"http://jobcoin.gemini.com/sanitary/api/"}

	if wallet.CheckBalance("Alice") != 37.5 {
		t.Fail()
	}
}
