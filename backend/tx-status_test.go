package main

import (
	"testing"
)

func TestTxStatus(t *testing.T) {
	s := &TxStatus{
		Statuses: map[string][]string{},
	}

	s.PushUpdate("wallet1", "a")
	s.PushUpdate("wallet1", "b")

	a := s.GetUpdates("wallet1")
	if a[0] != "a" || a[1] != "b" {
		t.Fail()
	}
}
